package provider

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rockiecn/check/cash"
	"github.com/rockiecn/check/check"
	"github.com/rockiecn/check/internal"
)

type Paychecks map[uint64]*check.Paycheck
type Provider struct {
	ProviderSK   string
	ProviderAddr common.Address

	Host string

	TxNonce uint64

	Pool map[common.Address]Paychecks
}

type IProvider interface {
	Verify(pchk *check.Paycheck, dataValue *big.Int) (bool, error)
	SendTx(pc *check.Paycheck) (tx *types.Transaction, err error)
}

func New(sk string) (IProvider, error) {
	pro := &Provider{
		ProviderSK:   sk,
		ProviderAddr: internal.KeyToAddr(sk),
		//Recorder:     NewRec(),
		Host: "http://localhost:8545",
	}

	return pro, nil
}

// CallApplyCheque - send tx to contract to call apply cheque method.
func (pro *Provider) SendTx(pc *check.Paycheck) (tx *types.Transaction, err error) {

	ethClient, err := internal.GetClient(pro.Host)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	auth, err := internal.MakeAuth(pro.ProviderSK, nil, nil, big.NewInt(1000), 9000000)
	if err != nil {
		return nil, errors.New("make auth failed")
	}

	// get contract instance from address
	cashInstance, err := cash.NewCash(pc.Check.ContractAddr, ethClient)
	if err != nil {
		return nil, errors.New("newcash failed")
	}

	// type convertion, from pc to cashpc for contract
	cashpc := cash.Paycheck{
		Check: cash.Check{
			Value:        pc.Check.Value,
			TokenAddr:    pc.Check.TokenAddr,
			Nonce:        pc.Check.Nonce,
			FromAddr:     pc.Check.FromAddr,
			ToAddr:       pc.Check.ToAddr,
			OpAddr:       pc.Check.OpAddr,
			ContractAddr: pc.Check.ContractAddr,
			CheckSig:     pc.Check.CheckSig,
		},
		PayValue:    pc.PayValue,
		PaycheckSig: pc.PaycheckSig,
	}
	tx, err = cashInstance.Withdraw(auth, cashpc)
	if err != nil {
		return nil, errors.New("tx failed")
	}

	fmt.Println("-> Now mine a block to complete tx.")

	return tx, nil
}

// verify before store paycheck into pool
func (pro *Provider) Verify(pchk *check.Paycheck, blockValue *big.Int) (bool, error) {

	// value should no less than payvalue
	if pchk.Check.Value.Cmp(pchk.PayValue) < 0 {
		return false, nil
	}

	// check nonce shuould larger than contract nonce
	contractNonce, err := internal.QueryNonce(pro.ProviderAddr, pchk.Check.ContractAddr, pro.ProviderAddr)
	if err != nil {
		return false, nil
	}
	if pchk.Check.Nonce <= contractNonce {
		return false, nil
	}

	// to address must be provider
	if pchk.Check.ToAddr != pro.ProviderAddr {
		return false, nil
	}

	// check nonce should larger than TxNonce(last withdrawed nonce)
	if pchk.Check.Nonce <= pro.TxNonce {
		return false, nil
	}

	pc := pro.Pool[pchk.Check.ToAddr][pchk.Check.Nonce]
	// verify payvalue
	if pc == nil {
		if pchk.PayValue.Cmp(blockValue) == 0 {
			return true, nil
		}
	} else {
		pay := pchk.PayValue.Sub(pchk.PayValue, pc.PayValue)
		if pay.Cmp(blockValue) == 0 {
			return true, nil
		}
	}

	return false, nil
}
