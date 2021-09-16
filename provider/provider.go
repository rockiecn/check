package provider

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/cash"
	"github.com/rockiecn/check/check"
	comn "github.com/rockiecn/check/common"
	"github.com/rockiecn/check/recorder"
)

type Provider struct {
	ProviderSK   string
	ProviderAddr common.Address

	Host string

	Recorder *recorder.Recorder
}

type IProvider interface {
	WithDraw(pc *check.Paycheck) error
}

func New(sk string) (IProvider, error) {
	pro := new(Provider)
	pro.ProviderSK = sk
	pro.ProviderAddr = comn.KeyToAddr(sk)

	pro.Recorder = recorder.New()

	pro.Host = "http://localhost:8545"

	return pro, nil
}

// CallApplyCheque - send tx to contract to call apply cheque method.
func (pro *Provider) WithDraw(pc *check.Paycheck) error {

	cli, err := comn.GetClient(pro.Host)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return err
	}
	defer cli.Close()

	auth, err := comn.MakeAuth(pro.ProviderSK, nil, nil, big.NewInt(1000), 9000000)
	if err != nil {
		return err
	}

	// get contract instance from address
	cashInstance, err := cash.NewCash(pc.Check.ContractAddr, cli)
	if err != nil {
		fmt.Println("NewCash err: ", err)
		return err
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
	_, err = cashInstance.Withdraw(auth, cashpc)
	if err != nil {
		fmt.Println("tx failed :", err)
		return err
	}

	fmt.Println("-> Now mine a block to complete tx.")

	return nil
}

// make sure a paycheck is legal
func (pro *Provider) Legalize(pc *check.Paycheck) (bool, error) {
	// paycheck signed by check.from
	if ok, _ := pc.Verify(); !ok {
		return false, errors.New("paycheck not signed by check.from")
	}
	// check signed by check.operator
	if ok, _ := pc.Check.Verify(); !ok {
		return false, errors.New("check not signed by check.operator")
	}

	// to address
	if pc.Check.ToAddr != pro.ProviderAddr {
		return false, errors.New("check.to must be provider's address")
	}

	// value
	if pc.PayValue.Cmp(pc.Check.Value) > 0 {
		return false, errors.New("illegal payvalue, should not larger than value")
	}

	// nonce
	nonceContract, _ := pro.GetNonce(pc)
	if pc.Check.Nonce < nonceContract {
		return false, errors.New("check is obsoleted, cannot withdraw")
	}

	return true, nil
}

func (pro *Provider) GetNonce(pc *check.Paycheck) (uint64, error) {

	cli, err := comn.GetClient(pro.Host)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return 0, err
	}
	defer cli.Close()

	// get contract instance from address
	cashInstance, err := cash.NewCash(pc.Check.ContractAddr, cli)
	if err != nil {
		fmt.Println("NewCash err: ", err)
		return 0, err
	}

	// get nonce
	nonce, err := cashInstance.GetNonce(nil, pc.Check.ToAddr)
	if err != nil {
		fmt.Println("tx failed :", err)
		return 0, err
	}

	return nonce, nil
}
