package user

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
	comn "github.com/rockiecn/check/common"
	"github.com/rockiecn/check/recorder"
)

type User struct {
	UserSK   string
	UserAddr common.Address

	Recorder *recorder.Recorder

	Host string
}

type IUser interface {
	GenPaycheck(chk *check.Check, payValue *big.Int) (*check.Paycheck, error)
}

func New(sk string) (IUser, error) {
	user := new(User)
	user.UserSK = sk
	user.UserAddr = comn.KeyToAddr(sk)

	user.Recorder = recorder.New()

	user.Host = "http://localhost:8545"

	return user, nil
}

// generate Paycheck based on check, sig of Paycheck is updated
func (user *User) GenPaycheck(chk *check.Check, payValue *big.Int) (*check.Paycheck, error) {
	pchk := new(check.Paycheck)
	pchk.Check = *chk
	pchk.PayValue = payValue

	err := pchk.Sign(user.UserSK)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// todo: record pchk into data

	return pchk, nil
}
