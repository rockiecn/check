package provider

import (
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

	// recorder for paycheck
	PaycheckRecorder *recorder.PRecorder
}

type IProvider interface {
	NewProvider(sk string) (*Provider, error)
	WithDraw(paycheque *check.Paycheck) error
}

func NewProvider(sk string) (*Provider, error) {
	pro := new(Provider)
	pro.ProviderSK = sk
	pro.ProviderAddr = comn.KeyToAddr(sk)

	pro.Host = "http://localhost:8545"

	pro.PaycheckRecorder = recorder.NewPRecorder()

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
	_, err = cashInstance.ApplyCheck(auth, cashpc)
	if err != nil {
		fmt.Println("tx failed :", err)
		return err
	}

	fmt.Println("-> Now mine a block to complete tx.")

	return nil
}
