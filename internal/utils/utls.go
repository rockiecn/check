package utils

import (
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/rockiecn/check/internal/cash"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Common struct {
	Host string
}

const HOST = "http://localhost:8545"

// get address from private key
func KeyToAddr(sk string) common.Address {
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

	return addr
}

// GetClient - dial to chain host
func GetClient(endPoint string) (*ethclient.Client, error) {
	rpcClient, err := rpc.Dial(endPoint)
	if err != nil {
		fmt.Println("rpc.Dial err:", err)
		return nil, err
	}

	ethClient := ethclient.NewClient(rpcClient)
	return ethClient, nil

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

func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// get contract nonce
func GetNonce(contract common.Address, to common.Address) (uint64, error) {

	cli, err := GetClient(HOST)
	if err != nil {
		return 0, errors.New("failed to dial geth")
	}
	defer cli.Close()

	// get contract instance from address
	cashInstance, err := cash.NewCash(contract, cli)
	if err != nil {
		return 0, errors.New("NewCash err")
	}

	// get nonce
	nonce, err := cashInstance.GetNonce(nil, to)
	if err != nil {
		return 0, errors.New("tx failed")
	}

	return nonce, nil
}

func String2BigInt(str string) *big.Int {
	v, _ := big.NewInt(0).SetString(str, 10)
	return v
}

func String2Byte(str string) []byte {
	v, _ := hex.DecodeString(str)
	return v
}

//
func BlockValue(s *big.Int, factor int64) *big.Int {
	bigF := big.NewInt(factor)
	return s.Mul(s, bigF)
}
