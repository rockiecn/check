package operator

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rockiecn/check/cash"
	"github.com/rockiecn/check/check"
	com "github.com/rockiecn/check/common"
	"github.com/rockiecn/check/recorder"
)

type Operator struct {
	OpSK   string
	OpAddr common.Address

	ContractAddr common.Address
	Host         string

	// each provider's nonce
	Nonces map[common.Address]*big.Int

	// recorder for check
	CheckRecorder *recorder.CRecorder
}

type IOperator interface {
	GenerateCheck(from string, to string, value *big.Int, nonce *big.Int) (*check.Check, error)
	RecordCheck(check *check.Check) error
	DeployContract() (string, common.Address, error)
}

func NewOperator(sk string, token string) (*Operator, error) {
	op := new(Operator)

	op.Host = "http://localhost:8545"

	op.OpSK = sk
	op.OpAddr = com.KeyToAddr(sk)

	_, contract, _ := op.DeployContract()
	op.ContractAddr = contract

	op.Nonces = make(map[common.Address]*big.Int)

	op.CheckRecorder = recorder.NewCRecorder()

	return op, nil
}

func (op *Operator) GenCheck(
	value *big.Int,
	token common.Address,
	from common.Address,
	to common.Address) (*check.Check, error) {

	chk := new(check.Check)

	chk.Value = value
	chk.TokenAddr = token
	chk.FromAddr = from
	chk.ToAddr = to
	chk.Nonce = op.Nonces[to]
	// nonce = nonce + 1
	bigOne := big.NewInt(1)
	op.Nonces[to].Add(op.Nonces[to], bigOne)

	chk.OpAddr = com.KeyToAddr(op.OpSK)
	chk.ContractAddr = op.ContractAddr

	chk.Sign(op.OpSK)

	op.CheckRecorder.Record(chk)

	return chk, nil
}

// todo: called by NewOperator
func (op *Operator) DeployContract() (string, common.Address, error) {

	var contractAddr common.Address

	client, err := com.GetClient(op.Host)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return "", contractAddr, err
	}
	defer client.Close()

	// get sk
	priKeyECDSA, err := crypto.HexToECDSA(op.OpSK)
	if err != nil {
		fmt.Println("HexToECDSA err: ", err)
		return "", contractAddr, err
	}

	// get pubkey
	pubKey := priKeyECDSA.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
		return "", contractAddr, err
	}

	// pubkey to address
	opComAddr := crypto.PubkeyToAddress(*pubKeyECDSA)
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

	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	auth, err := bind.NewKeyedTransactorWithChainID(priKeyECDSA, big.NewInt(1337))
	if err != nil {
		log.Println("NewKeyedTransactorWithChainID err:", err)
		return "", contractAddr, err
	}

	// set nonce
	auth.Nonce = big.NewInt(int64(nonce))
	// string to bigint
	bn := new(big.Int)
	bn, ok2 := bn.SetString("10000000000000000000", 10) // deploy 10 eth
	if !ok2 {
		fmt.Println("SetString: error")
		fmt.Println("big number SetString error")
		return "", contractAddr, err
	}
	auth.Value = bn                 // deploy 100 eth
	auth.GasLimit = uint64(7000000) // in units
	auth.GasPrice = gasPrice

	contractAddr, tx, _, err := cash.DeployCash(auth, client)
	if err != nil {
		fmt.Println("deployCashErr:", err)
		return "", contractAddr, err
	}

	return tx.Hash().String(), contractAddr, nil
}
