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

// first find a paycheck from pool, whose remain value is enough for paying.
// then generate a new paycheck with accumulated payvalue and new signature.
// finally, store the newly created paycheck into user pool
func (user *User) Pay(proAddr common.Address, dataValue *big.Int) (*check.Paycheck, error) {

	for _, v := range user.Pool[proAddr] {
		remain := new(big.Int).Sub(v.Check.Value, v.PayValue)
		if remain.Cmp(dataValue) >= 0 {
			// accumulate
			v.PayValue = new(big.Int).Add(v.PayValue, dataValue)
			// sign
			err := v.Sign(user.UserSK)
			if err != nil {
				return nil, errors.New("sign payckeck failed")
			}
			// store new paycheck into pool
			user.Pool[proAddr][v.Check.Nonce] = v
			return v, nil
		}
	}
	// usable paycheck not found
	return nil, errors.New("usable paycheck not found")
}
