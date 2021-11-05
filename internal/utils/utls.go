package utils

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
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
func SkToAddr(sk string) (common.Address, error) {
	skECDSA, err := crypto.HexToECDSA(sk)
	if err != nil {
		return common.Address{}, err
	}

	pubKey := skECDSA.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("error casting public key to ECDSA")
	}

	addr := crypto.PubkeyToAddress(*pubKeyECDSA)

	return addr, nil
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
	value *big.Int,
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
	auth.Value = value // 交易的金额
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

func WaitForMiner(tx *types.Transaction) error {
	// connect to geth
	ethClient, err := GetClient(HOST)
	if err != nil {
		return err
	}
	defer ethClient.Close()

	for {
		txReceipt, _ := ethClient.TransactionReceipt(context.Background(), tx.Hash())
		// receipt ok
		if txReceipt != nil {
			fmt.Println("receipt gas:", txReceipt.GasUsed)
			break
		}
		fmt.Println("waiting for miner, 5 seconds..")
		time.Sleep(time.Duration(5) * time.Second)
	}
	return nil
}

func GenerateSK() (string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	priv := hexutil.Encode(privateKeyBytes)[2:]

	return priv, nil
}
