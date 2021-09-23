package user

import (
	"errors"
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
	Store(pc *check.Check) (bool, error)
}

func New(sk string) (IUser, error) {
	user := &User{
		UserSK:   sk,
		UserAddr: comn.KeyToAddr(sk),
		Recorder: recorder.New(),
	}

	return user, nil
}

// generate Paycheck based on check, sig of Paycheck is updated
func (user *User) GenPaycheck(chk *check.Check, payValue *big.Int) (*check.Paycheck, error) {
	pchk := &check.Paycheck{
		Check:    *chk,
		PayValue: payValue,
	}

	err := pchk.Sign(user.UserSK)
	if err != nil {
		return nil, errors.New("sign payckeck failed")
	}

	// record pchk into data
	err = user.Recorder.Record(pchk)
	if err != nil {
		return nil, errors.New("record paycheck failed")
	}

	return pchk, nil
}

// tests before check been stored
func (user *User) Store(chk *check.Check) (bool, error) {
	// check signed by check.operator
	if ok, _ := chk.Verify(); !ok {
		return false, errors.New("check not signed by check.operator")
	}

	// from address = user address
	if chk.FromAddr != user.UserAddr {
		return false, errors.New("check's from address must be user")
	}

	// nonce >= contract.nonce
	nonceContract, _ := comn.GetNonce(chk.ContractAddr, chk.ToAddr)
	if chk.Nonce < nonceContract {
		return false, errors.New("check is obsoleted, cannot withdraw")
	}

	// paycheck should not exist in recorder
	if ok, _ := user.Recorder.IsValid(chk); ok {
		return false, errors.New("check already exist")
	}

	// ok to store
	user.Recorder.Record(chk)

	return true, nil
}
