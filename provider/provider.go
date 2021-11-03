package provider

import (
	"errors"
	"fmt"
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

	ContractAddr common.Address
	Host         string
	Pool         map[uint64]*check.Paycheck
}

type IProvider interface {
	Verify(pchk *check.Paycheck, dataValue *big.Int) (bool, error)
	GetNextPayable() (*check.Paycheck, error)
	Withdraw(pc *check.Paycheck) (tx *types.Transaction, err error)
}

func NewProvider(sk string) (IProvider, error) {
	addr, err := utils.KeyToAddr(sk)
	if err != nil {
		return nil, err
	}
	pro := &Provider{
		ProviderSK:   sk,
		ProviderAddr: addr,
		Host:         "http://localhost:8545",
	}

	return pro, nil
}

// verify paycheck before store paycheck into pool
func (pro *Provider) Verify(pchk *check.Paycheck, blockValue *big.Int) (bool, error) {

	// value should no less than payvalue
	if pchk.Check.Value.Cmp(pchk.PayValue) < 0 {
		return false, nil
	}

	// check nonce shuould larger than contract nonce
	contractNonce, err := utils.GetNonce(pro.ContractAddr, pro.ProviderAddr)
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
		if pchk.PayValue.Cmp(blockValue) == 0 {
			return true, nil
		} else {
			return false, errors.New("payvalue verify failed for new paycheck")
		}
	} else {
		pay := new(big.Int).Sub(pchk.PayValue, old.PayValue)
		if pay.Cmp(blockValue) == 0 {
			return true, nil
		} else {
			return false, errors.New("payvalue verify failed")
		}
	}
}

// get the next payable paycheck in pool
func (pro *Provider) GetNextPayable() (*check.Paycheck, error) {
	ctrNonce, err := utils.GetNonce(pro.ContractAddr, pro.ProviderAddr)
	if err != nil {
		return nil, err
	}

	paychecks := pro.Pool

	max := ^uint64(0)
	minNonce := max
	for k := range paychecks {
		if k < ctrNonce {
			continue
		}

		if k < minNonce {
			minNonce = k
		}
	}

	return pro.Pool[minNonce], nil
}

// set contract address for provider
func (pro *Provider) SetContract(ctAddr common.Address) {
	pro.ContractAddr = ctAddr
}

// CallApplyCheque - send tx to contract to call apply cheque method.
func (pro *Provider) SendTx(pc *check.Paycheck) (tx *types.Transaction, err error) {

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
