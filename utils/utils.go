package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rockiecn/check/check"
)

// calc hash of check, used to sign check and verify
func CheckHash(Check *check.Check) []byte {

	valuePad32 := common.LeftPadBytes(Check.Value.Bytes(), 32)
	noncePad32 := common.LeftPadBytes(Check.Nonce.Bytes(), 32)

	tokenBytes, _ := hex.DecodeString(Check.TokenAddr)
	fromBytes, _ := hex.DecodeString(Check.From)
	toBytes, _ := hex.DecodeString(Check.To)
	operatorBytes, _ := hex.DecodeString(Check.OperatorAddr)
	contractBytes, _ := hex.DecodeString(Check.ContractAddr)

	// calc hash
	hash := crypto.Keccak256(
		valuePad32,
		tokenBytes,
		noncePad32,
		fromBytes,
		toBytes,
		operatorBytes,
		contractBytes,
	)

	return hash
}

// calc hash of check, used to sign check and verify
func PayCheckHash(PayCheck *check.PayCheck) []byte {

	valuePad32 := common.LeftPadBytes(PayCheck.Check.Value.Bytes(), 32)
	noncePad32 := common.LeftPadBytes(PayCheck.Check.Nonce.Bytes(), 32)
	payvaluePad32 := common.LeftPadBytes(PayCheck.PayValue.Bytes(), 32)

	tokenBytes, _ := hex.DecodeString(PayCheck.Check.TokenAddr)
	fromBytes, _ := hex.DecodeString(PayCheck.Check.From)
	toBytes, _ := hex.DecodeString(PayCheck.Check.To)
	operatorBytes, _ := hex.DecodeString(PayCheck.Check.OperatorAddr)
	contractBytes, _ := hex.DecodeString(PayCheck.Check.ContractAddr)

	// calc hash
	hash := crypto.Keccak256(
		valuePad32,
		tokenBytes,
		noncePad32,
		fromBytes,
		toBytes,
		operatorBytes,
		contractBytes,
		payvaluePad32,
	)

	return hash
}

// get address from private key
func KeyToAddr(sk string) string {
	skECDSA, err := crypto.HexToECDSA(sk)
	if err != nil {
		fmt.Println("hex to ecdsa failed:", err)
	}

	pubKey := skECDSA.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	addr := crypto.PubkeyToAddress(*pubKeyECDSA)

	big.NewInt(1231231231231231231)
	return addr.String()
}

// GetClient - dial to chain host
func GetClient(endPoint string) (*ethclient.Client, error) {
	rpcClient, err := rpc.Dial(endPoint)
	if err != nil {
		fmt.Println("rpc.Dial err:", err)
		return nil, err
	}

	conn := ethclient.NewClient(rpcClient)
	return conn, nil
}

// MakeAuth - make a transactOpts to call contract
func MakeAuth(
	hexSk string,
	moneyToContract *big.Int,
	nonce *big.Int,
	gasPrice *big.Int,
	gasLimit uint64) (*bind.TransactOpts, error) {

	auth := new(bind.TransactOpts)

	priKeyECDSA, err := crypto.HexToECDSA(hexSk)
	if err != nil {
		log.Println("HexToECDSA err: ", err)
		return auth, err
	}

	auth, err = bind.NewKeyedTransactorWithChainID(priKeyECDSA, big.NewInt(1337))
	if err != nil {
		fmt.Println("make auth failed:", err)
		return nil, err
	}
	auth.GasPrice = gasPrice
	auth.Value = moneyToContract //放进合约里的钱
	auth.Nonce = nonce
	auth.GasLimit = gasLimit
	return auth, nil
}
