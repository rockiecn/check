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
	Host     string

	PoolA CheckPool
	PoolB PaycheckPool
}

type IUser interface {
	GenPaycheck(chk *check.Check, payValue *big.Int) (*check.Paycheck, error)

	PreStore(chk *check.Check) (bool, error)
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

// import check from a check file
func (user *User) ImportCheck(path string) (*check.Check, error) {
	//用户从接收到的支票文件导入支票并返回。
	return nil, nil
}

// get a check that has not been used yet
func (user *User) GetNew(to common.Address) (*check.Check, error) {
	/*
		如果paycheck数组为空，表示没有支票被用过，所有的check都能支付，则直接取出check池的第一张支票返回。

		否则，在paycheck数组中取出末尾项的nonce（队列中的最大nonce），然后在check池中找出第一张大于此nonce的支票返回。

		如果没找到，表示当前已无可用支票，返回空。
	*/
	return nil, nil
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

func (user *User) SendPaycheck(to common.Address, pc *check.Paycheck) error {
	//将paycheck发送给provider节点，以支付本次数据块的费用。
	return nil
}

// verify received check
func (user *User) PreStore(chk *check.Check) (bool, error) {

	// verify signature
	ok, err := chk.Verify()
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errors.New("check signature verify failed")
	}

	// from = user
	if chk.FromAddr != user.UserAddr {
		return false, errors.New("check's from address not user")
	}

	// nonce
	contractNonce, err := comn.QueryNonce(user.UserAddr, chk.ContractAddr, chk.ToAddr)
	if err != nil {
		return false, errors.New("query nonce error")
	}
	if chk.Nonce <= contractNonce {
		return false, errors.New("check nonce should larger than contract nonce")
	}

	// check must not exist in pool
	for _, v := range user.PoolA.Data[chk.ToAddr] {
		if chk.Nonce == v.Nonce {
			return false, errors.New("check already exist in check pool")
		}
	}

	// store check into pool
	user.PoolA.Store(chk)

	return true, nil
}

// get a new check for pay, the first check next to current nonce
func (user *User) GetVirgin(to common.Address) (*check.Check, error) {

	checks := user.PoolA.Data[to]
	// no check left
	if len(checks) == 0 {
		return nil, errors.New("check pool is nil")
	}

	cur, err := user.PoolB.GetCurrent(to)
	// get current paycheck failed
	if err != nil {
		return nil, err
	}

	// no current paycheck exist, means no check is in use, return fist check to pay
	if cur == nil {
		return checks[0], nil
	}

	// find virgin
	for k, v := range checks {
		if v.Nonce > cur.Check.Nonce {
			return checks[k], nil
		}
	}

	// not found
	return nil, errors.New("virgin check is not found")
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
