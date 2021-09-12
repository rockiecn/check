package operator

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rockiecn/check/check"
	"github.com/rockiecn/check/utils"
)

type Operator struct {
	OperatorSK   string
	OperatorAddr string
	ContractAddr string

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

func (op *Operator) DeployContract() (txHash string, err error) {
	op.ContractAddr = "asdf"
	return "0", nil
}

func NewOperator(sk string, token string) (*Operator, error) {
	op := new(Operator)

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

	check := new(check.Check)

	check.Value = value
	check.TokenAddr = token
	check.From = from
	check.To = to
	check.OperatorAddr = utils.KeyToAddr(op.OperatorSK)
	check.Nonce = nonce
	check.ContractAddr = op.ContractAddr

	return check, nil

}

// Sign check by operator's sk
func Sign(check *check.Check, skByte []byte) ([]byte, error) {

	hash := utils.CheckHash(check)

	//
	priKeyECDSA, err := crypto.HexToECDSA(string(skByte))
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
