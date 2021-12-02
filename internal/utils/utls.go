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

// get contract nonce
func GetCtNonce(contract common.Address, to common.Address) (uint64, error) {

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

	fmt.Println("Mining...")

	for {
		txReceipt, _ := ethClient.TransactionReceipt(context.Background(), tx.Hash())
		// receipt ok
		if txReceipt != nil {
			break
		}
		time.Sleep(time.Duration(1) * time.Second)
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

// send some coin
func SendCoin(senderSk string, receiverAddr common.Address, value *big.Int) (*types.Transaction, error) {

	//https://www.cxyzjd.com/article/mongo_node/89709286

	ethClient, err := GetClient(HOST)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	// get sk
	skECDSA, err := crypto.HexToECDSA(senderSk)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := skECDSA.Public()
	pkECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*pkECDSA)

	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000) // in units
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// create tx
	tx := types.NewTransaction(nonce, receiverAddr, value, gasLimit, gasPrice, nil)

	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), skECDSA)
	if err != nil {
		log.Fatal(err)
	}

	// send tx
	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0x77006fcb3938f648e2cc65bafd27dec30b9bfbe9df41f78498b9c8b7322a249e
	//fmt.Println("-> Now mine a block to complete tx.")
	return signedTx, nil
}

// get gas used, in wei
func GetGasUsed(tx *types.Transaction) (*big.Int, error) {
	ethClient, err := GetClient(HOST)
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()
	txReceipt, _ := ethClient.TransactionReceipt(context.Background(), tx.Hash())
	// receipt ok
	if txReceipt == nil {
		return nil, errors.New("txReceipt is nil")
	} else {
		//gwei to wei
		gasWei := new(big.Int).SetUint64(txReceipt.GasUsed * 1000)
		return gasWei, nil
	}
}

func Uint64ToByte(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}
func ByteToUint64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}

// get key from paycheck for db operation
// key = to address + nonce
func ToKey(a common.Address, n uint64) []byte {
	var key []byte
	key = append(key, a.Bytes()...)
	key = append(key, Uint64ToByte(n)...)

	return key
}

// parse key to provider address and nonce
func FromKey(key []byte) (common.Address, uint64) {
	// provider address, 20 bytes
	proAddr := common.BytesToAddress(key[:20])
	// nonce
	n := ByteToUint64(key[20:])
	return proAddr, n
}

// generate the batch check key
func ToBatchKey(a common.Address, min, max uint64) (key []byte) {
	key = append(key, a.Bytes()...)
	key = append(key, Uint64ToByte(min)...)
	key = append(key, Uint64ToByte(max)...)

	return
}

// get address ,nonce from a batch key
func FromBatchKey(key []byte) (to common.Address, min, max uint64) {
	// provider address, 20 bytes
	to = common.BytesToAddress(key[:20])
	// nonce, 8 bytes
	min = ByteToUint64(key[20:28])
	max = ByteToUint64(key[28:36])

	return
}
