package operator

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rockiecn/check/check"
)

type Operator struct {
	OperatorSK string
	TokenAddr  string

	Nonces map[string]*big.Int // provider to nonce, every provider's nonce

	//
	History map[string]*check.Check // keyHash -> key, check, key: "operator:xxx, provider:xxx, nonce:xxx"

	IOperator
}

type IOperator interface {
	NewOperator(sk string, tokenAddr string) (*Operator, error)
	GenerateCheck(from string, to string, value *big.Int, nonce *big.Int) (*check.Check, error)
	DeployContract() (txHash string, err error)
	Sign(check *check.Check, skByte []byte) ([]byte, error)
}

func (*Operator) NewOperator(sk string, token string) (*Operator, error) {
	op := new(Operator)

	op.OperatorSK = sk
	op.TokenAddr = token

	return op, nil
}

func (op *Operator) GenerateCheck(
	value *big.Int, // value of this check
	from string, // from address
	to string, // to address
	nonce *big.Int, // nonce
	conAddr string, // contract address
) (*check.Check, error) {

	check := new(check.Check)

	check.Value = value
	check.TokenAddr = op.TokenAddr
	check.From = from
	check.To = to
	check.OperatorAddr = op.KeyToAddr(op.OperatorSK)
	check.Nonce = nonce
	check.ContractAddr = conAddr

	return check, nil

}

// get address from private key
func (op *Operator) KeyToAddr(sk string) string {
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

// calc hash of check, used to sign check and verify
func CalcHash(Check *check.Check) []byte {

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

// Sign check by operator
func Sign(check *check.Check, skByte []byte) ([]byte, error) {

	hash := CalcHash(check)

	// byte to string, then string to ecdsa
	//privateKeyECDSA, err := crypto.HexToECDSA(utils.Byte2Str(skByte))
	privateKeyECDSA, err := crypto.HexToECDSA(string(skByte))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// compute digest
	//digest := crypto.Keccak256Hash(msg)

	// sign to bytes
	sigByte, err := crypto.Sign(hash, privateKeyECDSA)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	//fmt.Println("len sigByte:", len(sigByte))
	//fmt.Println("len skByte:", len(skByte))

	return sigByte, nil

}
