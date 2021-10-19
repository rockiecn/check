package provider

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rockiecn/check/cash"
	"github.com/rockiecn/check/check"
	comn "github.com/rockiecn/check/common"
)

type Provider struct {
	ProviderSK   string
	ProviderAddr common.Address

	Host string

	TxNonce uint64

	Pool PaycheckPool
}

type IProvider interface {
	// send withdraw transaction to contract
	SendTx(pc *check.Paycheck) (tx *types.Transaction, err error)
	CalcPay(pchk *check.Paycheck) (*big.Int, error)
	PreStore(pchk *check.Paycheck, dataValue *big.Int) (bool, error)
}

func New(sk string) (IProvider, error) {
	pro := &Provider{
		ProviderSK:   sk,
		ProviderAddr: comn.KeyToAddr(sk),
		//Recorder:     NewRec(),
		Host: "http://localhost:8545",
	}

	return pro, nil
}

// CallApplyCheque - send tx to contract to call apply cheque method.
func (pro *Provider) SendTx(pc *check.Paycheck) (tx *types.Transaction, err error) {

	ethClient, err := comn.GetClient(pro.Host)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	auth, err := comn.MakeAuth(pro.ProviderSK, nil, nil, big.NewInt(1000), 9000000)
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

// calculate the actual money the paycheck pays
func (pro *Provider) CalcPay(pchk *check.Paycheck) (*big.Int, error) {
	//如果paycheck为空，则直接返回它的payvalue。
	//否则，计算当前payvalue和paycheck数组末尾项的payvalue的差值并返回。
	return nil, nil
}

func (pro *Provider) PreStore(pchk *check.Paycheck, dataValue *big.Int) (bool, error) {

	// value should no less than payvalue
	if pchk.Check.Value.Cmp(pchk.PayValue) < 0 {
		return false, errors.New("value less than payvalue")
	}

	// check nonce shuould larger than contract nonce
	contractNonce, err := comn.QueryNonce(pro.ProviderAddr, pchk.Check.ContractAddr, pro.ProviderAddr)
	if err != nil {
		return false, errors.New("query contract nonce failed")
	}
	if pchk.Check.Nonce <= contractNonce {
		return false, errors.New("check nonce too small, cannot withdraw")
	}

	// to address must be provider
	if pchk.Check.ToAddr != pro.ProviderAddr {
		return false, errors.New("check's to address not provider")
	}

	// check nonce should larger than TxNonce(last withdrawed nonce)
	if pchk.Check.Nonce <= pro.TxNonce {
		return false, errors.New("check nonce not larger than TxNonce")
	}

	//
	pay, err := pro.CalcPay(pchk)
	if err != nil {
		return false, errors.New("call CalcPay failed")
	}
	if pay != dataValue {
		return false, errors.New("pay amount not equal dataValue")
	}

	// store paycheck into pool
	pro.Pool.Store(pchk)

	return true, nil
}

type PaycheckPool struct {
	Data []*check.Paycheck //按照nonce和payvalue有序
}

// 存储一张paycheck到池中
func (p *PaycheckPool) Store(pc *check.Paycheck) error {
	return nil
}

// 找出用于下一次提现的paycheck，如果找到了则返回它，如果没找到则返回空
func (pro *Provider) GetNextPayable() (*check.Paycheck, error) {
	contractNonce, err := comn.GetNonce(pro.ProviderAddr, pro.ProviderAddr)
	if err != nil {
		return nil, err
	}

	paychecks := pro.Pool.Data

	for k, v := range paychecks {
		if v.Check.Nonce > contractNonce {
			return paychecks[k], nil
		}
	}

	// no available nonce exist in pool
	return nil, errors.New("no paycheck in pool can withdraw")
}
