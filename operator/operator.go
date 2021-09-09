package operator

import (
	"math/big"

	"github.com/rockiecn/check/check"
)

type Operator struct {
	OperatorSK string
	TokenAddr  string

	Nonce  *big.Int
	Checks map[string]*check.Check // nonce to check

	IOperator
}

type IOperator interface {
	NewOperator(sk string, tokenAddr string) (*Operator, error)
	GenerateCheck(from string, to string, value *big.Int, nonce *big.Int) (*check.Check, error)
	DeployContract() (txHash string, err error)
}
