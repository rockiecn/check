package check

import "math/big"

type Check struct {
	Value        *big.Int
	TokenAddress string

	Nonce *big.Int
	From  string
	To    string

	OperatorAddr string
	ContractAddr string
}

type PayCheck struct {
	Check    *Check
	CheckSig []byte
	PayValue *big.Int
}
