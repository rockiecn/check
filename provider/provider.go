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
	PreStore(pchk *check.Paycheck, dataValue *big.Int) (bool, error)
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

// 验证一张paycheck的合法性。
// 首先验证两个签名是否正确。
// 然后是value值是否大于payvalue值。
// 然后是to地址跟provider地址是否相同。
// nonce值是否大于合约中to地址的当前nonce（决定了它是否能够提现）。
// nonce值是否大于txNonce的值（决定了它是否能够提现）。
// 从paycheck池中取出尾项的nonce（maxNonce）
// 如果支票的nonce大于maxNonce，说明此paycheck是使用新支票支付的，直接使用其payvalue值跟size的价值相比较，看是否相等
// 如果支票的nonce等于maxNonce，说明此paycheck是使用current支票进行支付的，计算payvalue的差值，跟size的价值相比较，看是否相等
// verify before store paycheck into pool
func (pro *Provider) PreStore(pchk *check.Paycheck, size *big.Int) (bool, error) {

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

	s := pro.Pool.Data
	var maxNonce uint64
	if len(s) == 0 {
		maxNonce = 0
	}
	maxNonce = uint64(len(s)) - 1

	// new check paying
	if pchk.Check.Nonce > maxNonce {
		if pchk.PayValue.Cmp(internal.BlockValue(size, 1)) == 0 {
			return true, nil
		}
	}
	// current check paying
	if pchk.Check.Nonce == maxNonce {
		pay := pchk.PayValue.Sub(pchk.PayValue, s[maxNonce].PayValue)
		if pay.Cmp(internal.BlockValue(size, 1)) == 0 {
			return true, nil
		}
	}

	return false, nil
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

// 先查看合约中节点对应的nonce，然后在本地paycheck池中从nonce+1开始找，一直找到一张存在的paycheck
// 如果到最后都没有找到一张paycheck，或者找到的paycheck是current支票，则返回nil
func (pro *Provider) GetNextPayable() (*check.Paycheck, error) {
	// get contract nonce
	contractNonce, err := internal.GetNonce(pro.ProviderAddr, pro.ProviderAddr)
	if err != nil {
		return nil, err
	}

	var lastNonce uint64
	s := pro.Pool.Data

	// paycheck pool is empty
	if len(s) == 0 {
		return nil, errors.New("paycheck pool is emtpy")
	}

	// last nonce is len(s)-1
	lastNonce = uint64(len(s)) - 1

	// from contractNonce+1, search for the first existing paycheck before lastNonce
	for n := contractNonce + 1; n < lastNonce; n++ {
		// found
		if s[n] != nil {
			return s[n], nil
		}
	}

	// not found
	return nil, errors.New("no payable paycheck found")
}
