package user

import (
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
	comn "github.com/rockiecn/check/common"
)

type User struct {
	UserSK   string
	UserAddr common.Address
	Host     string

	PoolA CheckPool
	PoolB PaycheckPool
}

type IUser interface {
	GenPaycheck(chk *check.Check, payValue *big.Int) (*check.Paycheck, error)

	Verify(chk *check.Check) (uint64, error)
	Pay(dataValue *big.Int) error
}

func New(sk string) (IUser, error) {
	user := &User{
		UserSK:   sk,
		UserAddr: comn.KeyToAddr(sk),
		//ChkRecorder:  NewChkRec(),
		//PChkRecorder: NewPChkRec(),
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

	// store paycheck into pool
	user.PoolB.Store(pchk)

	return pchk, nil
}

// verify received check
func (user *User) Verify(chk *check.Check) (uint64, error) {

	// verify signature
	ok, err := chk.Verify()
	if err != nil {
		return 1, err
	}
	if !ok {
		return 1, errors.New("check signature verify failed")
	}

	// from = user
	if chk.FromAddr != user.UserAddr {
		return 2, errors.New("check's from address not user")
	}

	// nonce
	contractNonce, err := comn.QueryNonce(user.UserAddr, chk.ContractAddr, chk.ToAddr)
	if err != nil {
		return 3, errors.New("query nonce error")
	}
	if chk.Nonce <= contractNonce {
		return 3, errors.New("check nonce should larger than contract nonce")
	}

	// check must not exist in pool
	for _, v := range user.PoolA.Data[chk.ToAddr] {
		if chk.Nonce == v.Nonce {
			return 4, errors.New("check already exist in check pool")
		}
	}

	// store check into pool
	user.PoolA.Store(chk)

	return 0, nil
}

// get a new check for pay, the first check next to current nonce
func (user *User) GetVirgin(to common.Address) (*check.Check, error) {
	// current paycheck
	cur, err := user.PoolB.GetCurrent(to)
	if err != nil {
		return nil, err
	}

	// checks
	checks := user.PoolA.Data[to]
	// no check exist in check pool
	if len(checks) == 0 {
		return nil, errors.New("check pool is nil")
	}

	// no paycheck exist, means no check is in use, return fist check to pay
	if cur == nil {
		return checks[0], nil
	}

	lastChk := checks[len(checks)-1]

	// the last check is used in current paycheck
	if lastChk.Nonce == cur.Check.Nonce {
		return nil, errors.New("no usable check already, need to buy more")
	}

	// the check in use is not exist
	if lastChk.Nonce < cur.Check.Nonce {
		return nil, errors.New("currently used check is not exist in check pool")
	}

	// lastcheck > current
	// find the used check first, then return the one next to it
	for k, v := range checks {
		if v.Nonce == cur.Check.Nonce {
			if checks[k+1] == nil {
				return nil, errors.New("last check > current, but no check available found")
			} else {
				// get virgin check
				return checks[k+1], nil
			}
		}
	}

	// using nonce is not exist in check pool?
	return nil, errors.New("the using check nonce is not found in check pool")
}

// price = size * factor
func (user *User) Price(size *big.Int, factor *big.Int) *big.Int {
	return size.Mul(size, factor)
}

// the process of pay: user send paycheck to provider
func (user *User) Pay(price *big.Int) error {
	return nil
}

type CheckPool struct {
	// to -> []check
	Data map[common.Address][]*check.Check //有序数组
}

// append a check into pool
func (p *CheckPool) Store(c *check.Check) error {

	checks := p.Data[c.ToAddr]

	// new nonce must be biggest
	if len(checks) > 0 && c.Nonce <= checks[len(checks)-1].Nonce {
		return errors.New("new nonce must be the biggest one")
	}

	// if check pool is nil or new check with the biggest nonce, ok to append
	p.Data[c.ToAddr] = append(p.Data[c.ToAddr], c)

	return nil
}

type PaycheckPool struct {
	// to -> []paycheck
	Data map[common.Address][]*check.Paycheck //按照nonce和payvalue有序
}

//
func (p *PaycheckPool) Store(pc *check.Paycheck) error {

	paychecks := p.Data[pc.Check.ToAddr]

	// slice is nil, just append it
	if len(paychecks) == 0 {
		p.Data[pc.Check.ToAddr] = append(paychecks, pc)
		return nil
	}

	// substitude last one with new paycheck
	last := paychecks[len(paychecks)-1]
	if pc.Check.Nonce == last.Check.Nonce {
		*last = *pc
		return nil
	}

	// append new paycheck into the slice
	if pc.Check.Nonce > last.Check.Nonce {
		p.Data[pc.Check.ToAddr] = append(paychecks, pc)
		return nil
	}

	// new paycheck's nonce too small
	if pc.Check.Nonce < last.Check.Nonce {
		return errors.New("new paycheck's nonce is too small, cannot withdraw")
	}

	return errors.New("exception occurd, should not see this")
}

// get the last paycheck, which has the biggest nonce
func (p *PaycheckPool) GetCurrent(to common.Address) (*check.Paycheck, error) {
	paychecks := p.Data[to]
	if len(paychecks) == 0 {
		return nil, nil
	} else {
		return paychecks[len(paychecks)-1], nil
	}
}

type Receipt struct {
	Dt    time.Time      // 购买日期
	Value *big.Int       // 购买金额
	Token common.Address // 货币类型
	Op    common.Address // 运营商地址
	From  common.Address // 付款方地址
	To    common.Address // 收款方地址
	Nonce uint64         // nonce
	Sig   string         // 运营商的签名
}

type ReceiptPool struct {
	Pool []*Receipt
}

// 存储一张收据到收据池
func (p *ReceiptPool) Store(r *Receipt) error {
	p.Pool = append(p.Pool, r)
	return nil
}
