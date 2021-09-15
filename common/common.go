package common

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
