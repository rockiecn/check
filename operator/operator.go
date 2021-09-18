package operator

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rockiecn/check/cash"
	"github.com/rockiecn/check/check"
	comn "github.com/rockiecn/check/common"
	"github.com/rockiecn/check/recorder"
)

type Operator struct {
	OpSK   string
	OpAddr common.Address

	ContractAddr common.Address

	// each provider's nonce
	Nonces map[common.Address]uint64

	Recorder *recorder.Recorder
}

type IOperator interface {
	GenCheck(value *big.Int, token common.Address, from common.Address, to common.Address) (*check.Check, error)
	DeployContract() (string, common.Address, error)
}

func New(sk string, token string) (IOperator, error) {
	op := &Operator{
		OpSK:     sk,
		OpAddr:   comn.KeyToAddr(sk),
		Nonces:   make(map[common.Address]uint64),
		Recorder: recorder.New(),
	}

	_, contract, err := op.DeployContract()
	if err != nil {
		return nil, err
	}
	op.ContractAddr = contract

	return op, nil
}

// generate a check
func (op *Operator) GenCheck(value *big.Int, token common.Address, from common.Address, to common.Address) (*check.Check, error) {

	// construct check
	chk := &check.Check{
		Value:        value,
		TokenAddr:    token,
		FromAddr:     from,
		ToAddr:       to,
		Nonce:        op.Nonces[to],
		OpAddr:       op.OpAddr,
		ContractAddr: op.ContractAddr,
	}

	// sign by operator
	chk.Sign(op.OpSK)

	// nonce increase
	op.Nonces[to] = op.Nonces[to] + 1

	// store check
	op.Recorder.Record(chk)

	return chk, nil
}

// todo: called by NewOperator
func (op *Operator) DeployContract() (string, common.Address, error) {

	var contractAddr common.Address

	client, err := comn.GetClient(comn.HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return "", contractAddr, err
	}
	defer client.Close()

	// string to ecdsa
	priKeyECDSA, err := crypto.HexToECDSA(op.OpSK)
	if err != nil {
		fmt.Println("HexToECDSA err: ", err)
		return "", contractAddr, err
	}

	// get pubkey
	pubKey := priKeyECDSA.Public()
	// ecdsa
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
		return "", contractAddr, err
	}
	// get address
	opComAddr := crypto.PubkeyToAddress(*pubKeyECDSA)
	// get nonce
	nonce, err := client.PendingNonceAt(context.Background(), opComAddr)
	if err != nil {
		log.Println(err)
		return "", contractAddr, err
	}

	// get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return "", contractAddr, err
	}

	// transfer to big.Int for contract
	bigNonce := new(big.Int).SetUint64(nonce)
	auth, err := comn.MakeAuth(op.OpSK, nil, bigNonce, gasPrice, 9000000)
	if err != nil {
		return "", common.Address{}, err
	}

	contractAddr, tx, _, err := cash.DeployCash(auth, client)
	if err != nil {
		fmt.Println("deployCashErr:", err)
		return "", contractAddr, err
	}

	return tx.Hash().String(), contractAddr, nil
}
