package user

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/utils"
)

// nonce to check
type Paychecks map[uint64]*check.Paycheck

type User struct {
	UserSK   string
	UserAddr common.Address
	Host     string

	// address to paychecks
	Pool map[common.Address]Paychecks
}

type IUser interface {
	StoreCheck(chk *check.Check) error
	Pay(to common.Address, dataValue *big.Int) (*check.Paycheck, error)
}

func New(sk string) (IUser, error) {
	addr, err := utils.SkToAddr(sk)
	if err != nil {
		return nil, err
	}
	user := &User{
		UserSK:   sk,
		UserAddr: addr,
		Pool:     make(map[common.Address]Paychecks),
	}

	return user, nil
}

// generate a paycheck from check ,and store it into pool
func (user *User) StoreCheck(chk *check.Check) error {
	if chk == nil {
		return errors.New("check nil")
	}

	pchk := &check.Paycheck{
		Check:    chk,
		PayValue: big.NewInt(0),
	}

	err := pchk.Sign(user.UserSK)
	if err != nil {
		return errors.New("paycheck sign error")
	}

	// create paychecks for a new provider
	if user.Pool[chk.ToAddr] == nil {
		user.Pool[chk.ToAddr] = make(Paychecks)
		// store paycheck into pool
		user.Pool[chk.ToAddr][chk.Nonce] = pchk
		return nil
	} else {
		if user.Pool[chk.ToAddr][chk.Nonce] != nil {
			return errors.New("paycheck with same nonce already exist in pool")
		} else {
			// store paycheck into pool
			user.Pool[chk.ToAddr][chk.Nonce] = pchk
			return nil
		}
	}

}

// Pay- Create a paycheck legal for paying dataValue
// First find the legal paycheck in pool.
// 1. Remain value is enough for paying dataValue.
// 2. Paycheck's nonce no less than nonce in contract
// 3. The one with the minimum nonce in the result
// Then update it with accumulated payvalue and new signature.
func (user *User) Pay(proAddr common.Address, dataValue *big.Int) (*check.Paycheck, error) {

	var (
		theOne   = (*check.Paycheck)(nil)
		minNonce = ^uint64(0)
	)

	// check each payvalue in user pool
	for _, v := range user.Pool[proAddr] {
		// get nonce in contract
		ctNonce, err := utils.GetCtNonce(v.Check.CtrAddr, v.Check.ToAddr)
		if err != nil {
			return nil, err
		}
		// nonce too old
		if v.Check.Nonce < ctNonce {
			continue
		}

		// remain value must no less than dataValue
		remain := new(big.Int).Sub(v.Check.Value, v.PayValue)
		if remain.Cmp(dataValue) < 0 {
			continue
		} else {
			// got one
			if v.Check.Nonce < minNonce {
				minNonce = v.Check.Nonce
				theOne = v
			} else {
				continue
			}
		}
	}

	// usable paycheck not found
	if theOne == nil {
		return nil, errors.New("usable paycheck not found")
	} else {
		// a tempor paycheck for sign
		newPchk := new(check.Paycheck)
		*newPchk = *theOne
		newPchk.PayValue = new(big.Int).Add(theOne.PayValue, dataValue)
		// sign
		err := newPchk.Sign(user.UserSK)
		if err != nil {
			return nil, errors.New("sign payckeck failed")
		}

		// update data in pool with new paycheck
		*theOne = *newPchk

		return theOne, nil
	}
}
