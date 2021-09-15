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

	// recorder for paycheck
	PaycheckRecorder *recorder.PRecorder
}

type IUser interface {
	GeneratePaycheck(check *check.Check, payValue *big.Int) (*check.Paycheck, error)
}

func NewUser(sk string) (*User, error) {
	user := new(User)
	user.UserSK = sk
	user.UserAddr = comn.KeyToAddr(sk)

	user.PaycheckRecorder = recorder.NewPRecorder()

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

	user.PaycheckRecorder.Record(pchk)

	return pchk, nil
}
