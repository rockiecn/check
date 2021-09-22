package operator

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	DeployContract(value *big.Int) (*types.Transaction, common.Address, error)
}

// new operator, a contract is deployed.
// tx's receipt should be checked to make sure contract deploying is completed.
func New(sk string, token string) (IOperator, *types.Transaction, error) {
	op := &Operator{
		OpSK:     sk,
		OpAddr:   comn.KeyToAddr(sk),
		Nonces:   make(map[common.Address]uint64),
		Recorder: recorder.New(),
	}

	// give 20 eth to new contract
	tx, addr, err := op.DeployContract(comn.String2BigInt("20000000000000000000"))
	if err != nil {
		return nil, nil, err
	}
	op.ContractAddr = addr

	return op, tx, nil
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

	// signed by operator
	err := chk.Sign(op.OpSK)
	if err != nil {
		return nil, err
	}

	// store check
	err = op.Recorder.Record(chk)
	if err != nil {
		return nil, err
	}

	// update nonce
	op.Nonces[to] = op.Nonces[to] + 1

	return chk, nil
}

// value: money to new contract
func (op *Operator) DeployContract(value *big.Int) (tx *types.Transaction, contractAddr common.Address, err error) {

	ethClient, err := comn.GetClient(comn.HOST)
	if err != nil {
		return nil, common.Address{}, err
	}
	defer ethClient.Close()

	// string to ecdsa
	priKeyECDSA, err := crypto.HexToECDSA(op.OpSK)
	if err != nil {
		return nil, common.Address{}, err
	}

	// get pubkey
	pubKey := priKeyECDSA.Public()
	// ecdsa
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, common.Address{}, errors.New("error casting public key to ECDSA")
	}
	// get address
	opComAddr := crypto.PubkeyToAddress(*pubKeyECDSA)
	// get nonce
	nonce, err := ethClient.PendingNonceAt(context.Background(), opComAddr)
	if err != nil {
		return nil, common.Address{}, err
	}

	// get gas price
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, common.Address{}, err
	}

	// transfer to big.Int for contract
	bigNonce := new(big.Int).SetUint64(nonce)
	auth, err := comn.MakeAuth(op.OpSK, value, bigNonce, gasPrice, 9000000)
	if err != nil {
		return nil, common.Address{}, err
	}

	contractAddr, tx, _, err = cash.DeployCash(auth, ethClient)
	if err != nil {
		return nil, common.Address{}, err
	}
	/*
		go func() {
			// deploy contract, wait for mining.
			for {
				txReceipt, _ := ethClient.TransactionReceipt(context.Background(), tx.Hash())
				// receipt ok
				if txReceipt != nil {
					break
				}
				fmt.Println("deploy mining..")
				time.Sleep(time.Duration(5) * time.Second)
			}
		}()
	*/
	return tx, contractAddr, nil
}
