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

	// address to checks
	Pool map[common.Address]Paychecks
}

type IUser interface {
	GenPaycheck(to common.Address, payValue *big.Int) (*check.Paycheck, error)
}

func New(sk string) (IUser, error) {
	user := &User{
		UserSK:   sk,
		UserAddr: utils.KeyToAddr(sk),
	}

	return user, nil
}

// first find a paycheck from pool, whose remain value is enough for paying.
// then generate a new paycheck with accumulated payvalue and new signature.
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
