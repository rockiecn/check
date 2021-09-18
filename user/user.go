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
	PreStore(pc *check.Paycheck) (bool, error)
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
		return nil, err
	}

	// todo: record pchk into data

	return pchk, nil
}

// tests before paycheck been stored
func (user *User) PreStore(pc *check.Paycheck) (bool, error) {
	// check signed by check.operator
	if ok, _ := pc.Check.Verify(); !ok {
		return false, errors.New("check not signed by check.operator")
	}

	// from address = user address
	if pc.Check.FromAddr != user.UserAddr {
		return false, errors.New("check's from address must be user")
	}

	// nonce >= contract.nonce
	nonceContract, _ := comn.GetNonce(pc.Check.ContractAddr, pc.Check.ToAddr)
	if pc.Check.Nonce < nonceContract {
		return false, errors.New("check is obsoleted, cannot withdraw")
	}

	// paycheck should not exist in recorder
	if ok, _ := user.Recorder.Exist(pc); ok {
		return false, errors.New("paycheck already exist")
	}

	return true, nil
}
