package user

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/utils"
	"github.com/rockiecn/check/order"
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
	StoreCheck(om *order.OrderMgr, chk *check.Check) error
	NewPaycheck(to common.Address, payValue *big.Int) (*check.Paycheck, error)
}

func New(sk string) (IUser, error) {
	addr, err := utils.SkToAddr(sk)
	if err != nil {
		return nil, err
	}
	user := &User{
		UserSK:   sk,
		UserAddr: addr,
	}

	return user, nil
}

// store a check into pool
func (user *User) StoreCheck(om *order.OrderMgr, chk *check.Check) error {
	return nil
}

// first find a paycheck from pool, whose remain value is enough for paying.
// then generate a new paycheck with accumulated payvalue and new signature.
func (user *User) NewPaycheck(proAddr common.Address, payValue *big.Int) (*check.Paycheck, error) {
	for _, v := range user.Pool[proAddr] {
		remain := new(big.Int).Sub(v.Check.Value, v.PayValue)
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
	// usable paycheck not found
	return nil, errors.New("not found")
}
