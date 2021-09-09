package user

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
)

type User struct {
	UserSK string

	//
	History map[string]*check.PayCheck // keyHash -> key, paycheck, key: "operator:xxx, provider:xxx, nonce:xxx"
}

type IUser interface {
	NewUser(sk string) (*User, error)
	VerifyCheck(check *check.Check, sig []byte, opAddr common.Address) (bool, error)
	GeneratePayCheck(check *check.Check, payValue *big.Int) (*check.PayCheck, error)
}
