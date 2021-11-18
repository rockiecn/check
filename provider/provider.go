package provider

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rockiecn/check/internal/cash"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/utils"
)

type Provider struct {
	ProviderSK   string
	ProviderAddr common.Address

	Host string
	Pool map[uint64]*check.Paycheck
}

type IProvider interface {
	Verify(pchk *check.Paycheck, dataValue *big.Int) (bool, error)
	StorePaycheck(pchk *check.Paycheck) error
	GetNextPayable() (*check.Paycheck, error)
	Withdraw(pc *check.Paycheck) (tx *types.Transaction, err error)
	QueryBalance() (*big.Int, error)
}

func New(sk string) (IProvider, error) {
	addr, err := utils.SkToAddr(sk)
	if err != nil {
		return nil, err
	}
	pro := &Provider{
		ProviderSK:   sk,
		ProviderAddr: addr,
		Host:         "http://localhost:8545",
		Pool:         make(map[uint64]*check.Paycheck),
	}

	return pro, nil
}

// verify paycheck before store paycheck into pool
func (pro *Provider) Verify(pchk *check.Paycheck, dataValue *big.Int) (bool, error) {

	// value should no less than payvalue
	if pchk.Check.Value.Cmp(pchk.PayValue) < 0 {
		return false, errors.New("value less than payvalue")
	}

	// check nonce shuould larger than contract nonce
	contractNonce, err := utils.GetCtNonce(pchk.Check.CtrAddr, pro.ProviderAddr)
	if err != nil {
		return false, err
	}
	if pchk.Check.Nonce < contractNonce {
		return false, errors.New("nonce should not less than contract nonce")
	}

	// to address must be provider
	if pchk.Check.ToAddr != pro.ProviderAddr {
		return false, errors.New("to address must be provider")
	}

	// get paycheck in pool
	old := pro.Pool[pchk.Check.Nonce]
	// verify payvalue
	if old == nil {
		if pchk.PayValue.Cmp(dataValue) == 0 {
			return true, nil
		} else {
			return false, errors.New("payAmount not equal dataValue 1")
		}
	} else {
		payAmount := new(big.Int).Sub(pchk.PayValue, old.PayValue)
		if payAmount.Cmp(dataValue) == 0 {
			return true, nil
		} else {
			return false, errors.New("payAmount not equal dataValue 2")
		}
	}
}

// store a paycheck into pool
func (pro *Provider) StorePaycheck(pchk *check.Paycheck) error {
	if pchk == nil {
		return errors.New("paycheck nil")
	}

	pro.Pool[pchk.Check.Nonce] = pchk

	return nil
}

// get the next payable paycheck in pool
func (pro *Provider) GetNextPayable() (*check.Paycheck, error) {

	var (
		theOne   = (*check.Paycheck)(nil)
		max      = ^uint64(0)
		minNonce = max
	)

	for k, v := range pro.Pool {
		// get current nonce in contract
		ctrNonce, err := utils.GetCtNonce(v.Check.CtrAddr, pro.ProviderAddr)
		if err != nil {
			return nil, err
		}

		// nonce too old, check next
		if k < ctrNonce {
			continue
		}

		if k < minNonce {
			minNonce = k
			theOne = v
		}
	}

	return theOne, nil
}

// CallApplyCheque - send tx to contract to call apply cheque method.
func (pro *Provider) Withdraw(pc *check.Paycheck) (tx *types.Transaction, err error) {

	ethClient, err := utils.GetClient(pro.Host)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	auth, err := utils.MakeAuth(pro.ProviderSK, nil, nil, big.NewInt(1000), 9000000)
	if err != nil {
		return nil, errors.New("make auth failed")
	}

	// get contract instance from address
	cashInstance, err := cash.NewCash(pc.Check.CtrAddr, ethClient)
	if err != nil {
		return nil, errors.New("newcash failed")
	}

	// type convertion, from pc to cashpc for contract
	cashpc := cash.Paycheck{
		Check: cash.Check{
			Value:     pc.Check.Value,
			TokenAddr: pc.Check.TokenAddr,
			Nonce:     pc.Check.Nonce,
			FromAddr:  pc.Check.FromAddr,
			ToAddr:    pc.Check.ToAddr,
			OpAddr:    pc.Check.OpAddr,
			CtrAddr:   pc.Check.CtrAddr,
			CheckSig:  pc.Check.CheckSig,
		},
		PayValue:    pc.PayValue,
		PaycheckSig: pc.PaycheckSig,
	}
	tx, err = cashInstance.Withdraw(auth, cashpc)
	if err != nil {
		return nil, errors.New("tx failed")
	}

	//fmt.Println("Mine a block to complete.")

	return tx, nil
}

// CallApplyCheque - send tx to contract to call apply cheque method.
func (pro *Provider) WithdrawBatch(bc *check.BatchCheck) (tx *types.Transaction, err error) {

	// connect
	ethClient, err := utils.GetClient(pro.Host)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	// auth
	auth, err := utils.MakeAuth(pro.ProviderSK, nil, nil, big.NewInt(1000), 9000000)
	if err != nil {
		return nil, errors.New("make auth failed")
	}

	// get contract instance from address
	cashInstance, err := cash.NewCash(bc.CtrAddr, ethClient)
	if err != nil {
		return nil, errors.New("newcash failed")
	}

	// type convertion, from pc to cashbc for contract
	cashbc := cash.BatchCheck{
		OpAddr:     bc.OpAddr,
		ToAddr:     bc.ToAddr,
		CtrAddr:    bc.CtrAddr,
		TokenAddr:  bc.TokenAddr,
		BatchValue: bc.BatchValue,
		MinNonce:   bc.MinNonce,
		MaxNonce:   bc.MaxNonce,
		BatchSig:   bc.BatchSig,
	}

	tx, err = cashInstance.WithdrawBatch(auth, cashbc)
	if err != nil {
		return nil, errors.New("tx failed")
	}

	//fmt.Println("Mine a block to complete.")

	return tx, nil
}

// query provider balance
func (pro *Provider) QueryBalance() (*big.Int, error) {
	ethClient, err := utils.GetClient(utils.HOST)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	balance, err := ethClient.BalanceAt(context.Background(), pro.ProviderAddr, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
