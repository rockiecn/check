package user

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
	comn "github.com/rockiecn/check/common"
)

type User struct {
	UserSK   string
	UserAddr common.Address

	ChkRecorder  *ChkRecorder
	PChkRecorder *PChkRecorder

	Host string
}

type IUser interface {
	GenPaycheck(chk *check.Check, payValue *big.Int) (*check.Paycheck, error)
	VerifyCheck(pc *check.Check) (bool, error)
}

func New(sk string) (IUser, error) {
	user := &User{
		UserSK:       sk,
		UserAddr:     comn.KeyToAddr(sk),
		ChkRecorder:  NewChkRec(),
		PChkRecorder: NewPChkRec(),
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
	err = user.PChkRecorder.Record(pchk)
	if err != nil {
		return nil, errors.New("record paycheck failed")
	}

	return pchk, nil
}

// tests before check been stored
func (user *User) VerifyCheck(chk *check.Check) (bool, error) {
	// check signed by check.operator
	if ok, _ := chk.Verify(); !ok {
		return false, errors.New("check not signed by check.operator")
	}

	// from address = user address
	if chk.FromAddr != user.UserAddr {
		return false, errors.New("check's from address must be user")
	}

	// nonce > contract.nonce
	ctrNonce, _ := comn.GetNonce(chk.ContractAddr, chk.ToAddr)
	if chk.Nonce <= ctrNonce {
		return false, errors.New("check is obsoleted, cannot withdraw")
	}

	// check should not exist in recorder
	if ok, _ := user.ChkRecorder.ChkExist(chk); ok {
		return false, errors.New("check already exist")
	}

	// ok to store
	//user.ChkRecorder.Record(chk)

	return true, nil
}

type Key struct {
	Operator common.Address
	Provider common.Address
	Nonce    uint64
}

type ChkRecorder struct {
	Checks map[Key]*check.Check
}

// generate a recorder for operator
func NewChkRec() *ChkRecorder {

	r := &ChkRecorder{
		Checks: make(map[Key]*check.Check),
	}

	return r
}

// put a check into Checks
func (r *ChkRecorder) Record(chk *check.Check) error {

	key := Key{
		Operator: chk.OpAddr,
		Provider: chk.ToAddr,
		Nonce:    chk.Nonce,
	}
	r.Checks[key] = chk
	return nil
}

// if a check is valid to store
func (r *ChkRecorder) ChkExist(chk *check.Check) (bool, error) {

	k := Key{
		Operator: chk.OpAddr,
		Provider: chk.ToAddr,
		Nonce:    chk.Nonce,
	}
	v := r.Checks[k]

	if v == nil {
		return true, nil // not exist, ok to store
	} else {
		return false, nil // already exist
	}
}

type PChkRecorder struct {
	Paychecks map[Key]*check.Paycheck
}

// generate a recorder for operator
func NewPChkRec() *PChkRecorder {

	r := &PChkRecorder{
		Paychecks: make(map[Key]*check.Paycheck),
	}

	return r
}

// put a paycheck into Checks
func (r *PChkRecorder) Record(pchk *check.Paycheck) error {

	key := Key{
		Operator: pchk.Check.OpAddr,
		Provider: pchk.Check.ToAddr,
		Nonce:    pchk.Check.Nonce,
	}

	r.Paychecks[key] = pchk
	return nil
}

// if a check is valid to store
func (r *PChkRecorder) IsValid(pchk *check.Paycheck) (bool, error) {

	k := Key{
		Operator: pchk.Check.OpAddr,
		Provider: pchk.Check.ToAddr,
		Nonce:    pchk.Check.Nonce,
	}
	v := r.Paychecks[k]

	if v == nil {
		return true, nil // not exist, ok to store
	} else {
		return false, nil // already exist
	}
}
