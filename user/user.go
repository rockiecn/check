package user

import (
	"errors"
	"fmt"
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

	// store paycheck into pool
	if user.Pool[chk.ToAddr] == nil {
		user.Pool[chk.ToAddr] = make(Paychecks)
	}
	if user.Pool[chk.ToAddr][chk.Nonce] != nil {
		return errors.New("paycheck already exist in pool")
	}

	fmt.Printf("pro:%s\n", pchk.Check.ToAddr)
	fmt.Printf("nonce:%d\n", pchk.Check.Nonce)
	fmt.Printf("pchk.payvalue:%s\n", pchk.PayValue)
	fmt.Printf("pchk.paysig:%x\n", pchk.PaycheckSig)

	user.Pool[chk.ToAddr][chk.Nonce] = pchk

	return nil
}

// first find a paycheck from pool, whose remain value is enough for paying.
// then generate a new paycheck with accumulated payvalue and new signature.
func (user *User) NewPaycheck(proAddr common.Address, payValue *big.Int) (*check.Paycheck, error) {

	fmt.Println("pro:", proAddr)
	fmt.Printf("pool:%v\n", user.Pool[proAddr])

	for _, v := range user.Pool[proAddr] {
		fmt.Println("nonce", v.Check.Nonce)
		fmt.Printf("sig:%x\n", user.Pool[proAddr][0].PaycheckSig)
		remain := new(big.Int).Sub(v.Check.Value, v.PayValue)
		if remain.Cmp(payValue) >= 0 {
			// aggregate
			v.PayValue = new(big.Int).Add(v.PayValue, payValue)
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
