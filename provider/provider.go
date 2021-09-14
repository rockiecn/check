package provider

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rockiecn/check/cash"
	"github.com/rockiecn/check/check"
	"github.com/rockiecn/check/utils"
)

type Provider struct {
	ProviderSK   string
	ProviderAddr string

	Host string

	//
	History map[string]*check.PayCheck // keyHash -> key, paycheck, key: "operator:xxx, provider:xxx, nonce:xxx"
}

type IProvider interface {
	NewProvider(sk string) (*Provider, error)
	VerifyPayCheck(paycheck check.PayCheck) (bool, error)
	RecordPayCheck(check *check.PayCheck) error
	// call contract
	WithDraw(paycheque *check.PayCheck) error
}

func NewProvider(sk string) (*Provider, error) {
	pro := new(Provider)
	pro.ProviderSK = sk
	pro.ProviderAddr = utils.KeyToAddr(sk)

	pro.Host = "http://localhost:8545"

	return pro, nil
}

func (pro *Provider) VerifyPayCheck(paycheck *check.PayCheck) (bool, error) {
	hash := utils.PayCheckHash(paycheck)

	// signature to public key
	pubKeyECDSA, err := crypto.SigToPub(hash, paycheck.PayCheckSig)
	if err != nil {
		log.Println("SigToPub err:", err)
		return false, err
	}

	// pub key to common.address
	singerAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	ok := singerAddr == common.HexToAddress(paycheck.Check.From)

	return ok, nil
}

// CallApplyCheque - send tx to contract to call apply cheque method.
func (pro *Provider) CallContract(paycheck *check.PayCheck) error {
	pcForContract, err := pro.PayCheckAdaptor(paycheck)
	if err != nil {
		fmt.Println("paycheck adaptor failed:", err)
	}

	cli, err := utils.GetClient(pro.Host)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return err
	}
	defer cli.Close()

	auth, err := utils.MakeAuth(pro.ProviderSK, nil, nil, big.NewInt(1000), 9000000)
	if err != nil {
		return err
	}

	// get contract instance from address
	cashInstance, err := cash.NewCash(pcForContract.Check.ContractAddr, cli)
	if err != nil {
		fmt.Println("NewCash err: ", err)
		return err
	}

	fmt.Println("tx info:")
	fmt.Printf("value:%s\n", pcForContract.Check.Value.String())
	fmt.Printf("TokenAddress:%s\n", pcForContract.Check.TokenAddr)
	fmt.Printf("NodeNonce:%s\n", pcForContract.Check.Nonce.String())
	fmt.Printf("From:%s\n", pcForContract.Check.FromAddr.String())
	fmt.Printf("To:%s\n", pcForContract.Check.ToAddr.String())
	fmt.Printf("OperatorAddress:%s\n", pcForContract.Check.OpAddr.String())
	fmt.Printf("ContractAddr:%s\n", pcForContract.Check.ContractAddr)
	fmt.Printf("checkSig:%x\n", pcForContract.Check.CheckSig)
	fmt.Printf("PayValue:%s\n", pcForContract.PayValue.String())
	fmt.Printf("paycheckSig:%x\n", pcForContract.PaycheckSig)

	_, err = cashInstance.ApplyCheck(auth, *pcForContract)
	if err != nil {
		fmt.Println("tx failed :", err)
		return err
	}

	fmt.Println("-> Now mine a block to complete tx.")

	return nil
}

// convert check.PayCheck to cash.Paycheck
func (pro *Provider) PayCheckAdaptor(paycheck *check.PayCheck) (*cash.Paycheck, error) {
	// cheque
	pcForContract := new(cash.Paycheck)

	pcForContract.Check.Value = paycheck.Check.Value
	pcForContract.Check.TokenAddr = common.HexToAddress(paycheck.Check.TokenAddr)
	pcForContract.Check.Nonce = paycheck.Check.Nonce
	pcForContract.Check.FromAddr = common.HexToAddress(paycheck.Check.From)
	pcForContract.Check.ToAddr = common.HexToAddress(paycheck.Check.To)
	pcForContract.Check.OpAddr = common.HexToAddress(paycheck.Check.OperatorAddr)
	pcForContract.Check.ContractAddr = common.HexToAddress(paycheck.Check.ContractAddr)
	pcForContract.Check.CheckSig = paycheck.Check.CheckSig
	pcForContract.PayValue = paycheck.PayValue
	pcForContract.PaycheckSig = paycheck.PayCheckSig

	return pcForContract, nil
}
