package user

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
	"github.com/rockiecn/check/internal"
)

// nonce to check
type Paychecks map[uint64]*check.Paycheck

type User struct {
	UserSK   string
	UserAddr common.Address
	Host     string

	// address to checks
	Pool map[common.Address]Paychecks
}

type IUser interface {
	GenPaycheck(to common.Address, payValue *big.Int) (*check.Paycheck, error)
}

func New(sk string) (IUser, error) {
	user := &User{
		UserSK:   sk,
		UserAddr: internal.KeyToAddr(sk),
	}

	return user, nil
}

// find a paycheck from pool first, whose remain value is enouph for payValue
// then generate a new paycheck with aggregated payvalue and new signature
func (user *User) GenPaycheck(to common.Address, payValue *big.Int) (*check.Paycheck, error) {
	for _, v := range user.Pool[to] {
		remain := v.Check.Value.Sub(v.Check.Value, v.PayValue)
		if remain.Cmp(payValue) >= 0 {
			// aggregate
			v.PayValue = v.PayValue.Add(v.PayValue, payValue)
			// sign
			err := v.Sign(user.UserSK)
			if err != nil {
				return nil, errors.New("sign payckeck failed")
			}
			return v, nil
		}
	}
	// not found an usable paycheck
	return nil, nil
}
