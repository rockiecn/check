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

	Recorder *Recorder
}

type IProvider interface {
	WithDraw(pc *check.Paycheck) (*types.Transaction, error)
	Store(pc *check.Paycheck) (bool, error)
}

func New(sk string) (IProvider, error) {
	pro := &Provider{
		ProviderSK:   sk,
		ProviderAddr: comn.KeyToAddr(sk),
		Recorder:     NewRec(),
		Host:         "http://localhost:8545",
	}

	return pro, nil
}

// CallApplyCheque - send tx to contract to call apply cheque method.
func (pro *Provider) WithDraw(pc *check.Paycheck) (tx *types.Transaction, err error) {

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

// tests before paycheck been stored
func (pro *Provider) Store(pc *check.Paycheck) (bool, error) {

	// check signed by check.operator
	if ok, _ := pc.Check.Verify(); !ok {
		return false, errors.New("check not signed by check.operator")
	}

	// paycheck signed by check.from
	if ok, _ := pc.Verify(); !ok {
		return false, errors.New("paycheck not signed by check.from")
	}

	// payvalue must >= 0
	if pc.PayValue.Cmp(big.NewInt(0)) < 0 {
		return false, errors.New("illegal payvalue, should not be negtive")
	}

	// payvalue <= value
	if pc.PayValue.Cmp(pc.Check.Value) > 0 {
		return false, errors.New("illegal payvalue, should not larger than value")
	}

	// to address
	if pc.Check.ToAddr != pro.ProviderAddr {
		return false, errors.New("check.to must be provider's address")
	}

	// nonce >= contract.nonce
	nonceContract, _ := comn.GetNonce(pc.Check.ContractAddr, pc.Check.ToAddr)
	if pc.Check.Nonce < nonceContract {
		return false, errors.New("check is obsoleted, cannot withdraw")
	}

	// valid?
	if ok, _ := pro.Recorder.IsValid(pc); !ok {
		return false, errors.New("paycheck not valid")
	}

	// ok to store
	pro.Recorder.Record(pc)

	return true, nil
}

type Key struct {
	Operator common.Address
	Provider common.Address
	Nonce    uint64
}

type Recorder struct {
	Paychecks map[Key]*check.Paycheck
}

// generate a recorder for operator
func NewRec() *Recorder {

	r := &Recorder{
		Paychecks: make(map[Key]*check.Paycheck),
	}

	return r
}

// put a paycheck into Checks
func (r *Recorder) Record(pchk *check.Paycheck) error {

	key := Key{
		Operator: pchk.Check.OpAddr,
		Provider: pchk.Check.ToAddr,
		Nonce:    pchk.Check.Nonce,
	}

	r.Paychecks[key] = pchk
	return nil
}

// if a check is valid to store
func (r *Recorder) IsValid(pchk *check.Paycheck) (bool, error) {

	k := Key{
		Operator: pchk.Check.OpAddr,
		Provider: pchk.Check.ToAddr,
		Nonce:    pchk.Check.Nonce,
	}
	v := r.Paychecks[k]

	if v == nil {
		return true, nil // not exist, ok to store
	} else {
		return false, nil // already exist
	}
}
