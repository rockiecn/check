package check

import "math/big"

type Check struct {
	Value     *big.Int
	TokenAddr string

	Nonce *big.Int
	From  string
	To    string

	OperatorAddr string
	ContractAddr string

	CheckSig string
}

type PayCheck struct {
	Check    *Check
	PayValue *big.Int

	PayCheckSig []byte
}
