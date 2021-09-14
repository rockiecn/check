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
	"github.com/rockiecn/check/utils"
)

type Operator struct {
	OperatorSK   string
	OperatorAddr string
	ContractAddr string
	Host         string

	Nonces map[string]*big.Int // provider to nonce, every provider's nonce

	//
	History map[string]*check.Check // keyHash -> key, check, key: "operator:xxx, provider:xxx, nonce:xxx"

}

type IOperator interface {
	DeployContract() (txHash string, err error)
	GenerateCheck(from string, to string, value *big.Int, nonce *big.Int) (*check.Check, error)
	RecordCheck(check *check.Check) error
	Sign(check *check.Check, skByte []byte) ([]byte, error)
}

func NewOperator(sk string, token string) (*Operator, error) {
	op := new(Operator)

	op.Host = "http://localhost:8545"

	op.OperatorSK = sk
	op.OperatorAddr = utils.KeyToAddr(sk)

	return op, nil
}

func (op *Operator) GenerateCheck(
	value *big.Int, // value of this check
	token string, // token address
	from string, // from address
	to string, // to address
	nonce *big.Int, // nonce
) (*check.Check, error) {

	chk := new(check.Check)

	chk.Value = value
	chk.TokenAddr = token
	chk.From = from
	chk.To = to
	chk.Nonce = nonce

	chk.OperatorAddr = utils.KeyToAddr(op.OperatorSK)
	chk.ContractAddr = op.ContractAddr

	sigByte, _ := op.Sign(chk)
	chk.CheckSig = sigByte

	return chk, nil

}

// Sign check by operator's sk
func (op *Operator) Sign(check *check.Check) ([]byte, error) {

	hash := utils.CheckHash(check)

	//
	priKeyECDSA, err := crypto.HexToECDSA(op.OperatorSK)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// sign to bytes
	sigByte, err := crypto.Sign(hash, priKeyECDSA)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return sigByte, nil

}

func (op *Operator) DeployContract() (string, common.Address, error) {

	var contractAddr common.Address

	client, err := utils.GetClient(op.Host)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return "", contractAddr, err
	}
	defer client.Close()

	// get sk
	priKeyECDSA, err := crypto.HexToECDSA(op.OperatorSK)
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
