package user

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
)

type User struct {
	UserSK string

	PayChecks map[string]*check.PayCheck // (operator,nonce) to paycheck
}

type IUser interface {
	NewUser(sk string) (*User, error)
	VerifyCheck(check *check.Check, sig []byte, opAddr common.Address) (bool, error)
	GeneratePayCheck(check *check.Check, payValue *big.Int) (*check.PayCheck, error)
}
