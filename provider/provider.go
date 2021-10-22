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

// 先查看当前nonce是否超过了切片的长度
// 如果超过了，说明这是一张新用于支付的支票，需要将切片扩展到当前nonce的长度，并且将它存放到nonce所在位置，然后返回它的payvalue作为支付金额。
// 如果nonce在切片当前长度范围内，则先看此nonce位置知否已经存在paycheck。
// 如果已存在，则计算当前支票和nonce所在位置的paycheck的payvalue差值并返回。
// 如果不存在，则将它存放到当前nonce位置，并返回其payvalue值。
// calculate the actual money the paycheck pays
func (pro *Provider) CalcPay(pc *check.Paycheck) (*big.Int, error) {
	s := pro.Pool.Data
	// put check into right position
	if pc.Check.Nonce+1 > uint64(len(s)) {
		// padding nils
		for n := uint64(len(s)); n < pc.Check.Nonce; n++ {
			s = append(s, nil)
		}
		// right position after nils, and append it
		s = append(s, pc)
		pro.Pool.Data = s
		return pc.PayValue, nil
	}

	if s[pc.Check.Nonce] != nil {
		pay := pc.PayValue.Sub(pc.PayValue, s[pc.Check.Nonce].PayValue)
		return pay, nil
	}

	s[pc.Check.Nonce] = pc
	return pc.PayValue, nil
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

// 先查看当前nonce是否越界
// 如果nonce越界，则先使用nil填充池，直到nonce前的位置，然后把paycheck添加到pool中
// 如果nonce没有越界，并且paycheck已经存在于池中，则先比较两个payvalue值
// 如果payvalue值大于之前的，则使用新paycheck替换旧的
// 否则，忽略此paycheck
// 如果nonce没越界，并且check不存在于池中，则将paycheck直接放到nonce指定的位置
// called when a paycheck is received by provider
func (p *PaycheckPool) Store(pc *check.Paycheck) error {
	// get slice
	s := p.Data

	// if nonce is out of boundary, extend pool and put paycheck into right position
	if pc.Check.Nonce+1 > uint64(len(s)) {
		// padding nils
		for n := uint64(len(s)); n < pc.Check.Nonce; n++ {
			s = append(s, nil)
		}
		// right position after nils, and append check
		s = append(s, pc)
		p.Data = s
		return nil
	}

	// if nonce is inside current pool, but paycheck already exist, substitude with new one
	if s[pc.Check.Nonce] != nil {
		if pc.PayValue.Cmp(s[pc.Check.Nonce].PayValue) > 0 {
			s[pc.Check.Nonce] = pc
			return nil
		}
	}
	// paycheck not exist, append it
	s = append(s, pc)
	p.Data = s
	return nil
}

// 先查看合约中节点对应的nonce，然后在本地paycheck池中找出第一个比它大的支票
// 如果此支票不是current支票就正常返回它。
// 如果没有合适的可提现paycheck，就返回空
func (pro *Provider) GetNextPayable() (*check.Paycheck, error) {
	// get contract nonce
	contractNonce, err := comn.GetNonce(pro.ProviderAddr, pro.ProviderAddr)
	if err != nil {
		return nil, err
	}

	var lastNonce uint64
	s := pro.Pool.Data
	if len(s) != 0 {
		lastNonce = uint64(len(s)) - 1
	} else {
		lastNonce = 0
	}

	// from contract nonce, search for the first existing paycheck
	for n := contractNonce + 1; n < lastNonce; n++ {
		// found
		if s[n] != nil {
			return s[n], nil
		}
	}

	// not found
	return nil, errors.New("payable not found")
}
