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

// the process of pay: user send paycheck to provider
func (user *User) Pay(price *big.Int) error {
	return nil
}

// 如果paycheck队列为空，表示没有支票被用过，所有的check都能支付，则直接取出check池的第一张支票返回。
// 如果paycheck队列不为空，则在paycheck队列中取出末尾项的nonce（队列中的最大nonce）
// 然后在check池中从nonce+1开始向后找，一直找到存在支票的数据项返回。
// 如果一直找到切片末尾都是空值，表示当前已无可用支票，返回空。
// get a check that has not been used yet
func (user *User) GetNew(to common.Address) (*check.Check, error) {

	sc := user.PoolA.Data[to]
	sp := user.PoolB.Data[to]

	maxCNonce := len(sc) - 1
	maxPCNonce := len(sp) - 1

	// no check in pool
	if len(sc) == 0 {
		return nil, errors.New("no check in pool")
	}

	// if no paycheck exist, return the first check
	if len(sp) == 0 {
		return sc[1], nil
	}

	// last paycheck's nonce
	if maxCNonce <= maxPCNonce {
		return nil, errors.New("no usable check left, need buy more")
	}

	// search for an usable check in pool
	for i := maxPCNonce + 1; i <= maxCNonce; i++ {
		if sc[i] != nil {
			return sc[i], nil
		}
	}

	// no check found
	return nil, errors.New("no usable check in check pool")
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

// 验证支票签名(operator)。
// 支票的from字段是否等于user地址。
// 支票的nonce必须大于合约中节点地址对应的当前nonce。
// 支票在本地池中不能已存在（不能有相同nonce）。
// 验证通过返回true，否则返回false
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

	// nonce must bigger than contract nonce
	contractNonce, err := comn.QueryNonce(user.UserAddr, chk.ContractAddr, chk.ToAddr)
	if err != nil {
		return false, errors.New("query nonce error")
	}
	if chk.Nonce <= contractNonce {
		return false, errors.New("check nonce must larger than contract nonce")
	}

	// check must not exist in pool
	sc := user.PoolA.Data[chk.ToAddr]
	if sc[chk.Nonce] != nil {
		return false, errors.New("check already exist in check pool")
	}

	// store check into pool
	user.PoolA.Store(chk)

	return true, nil
}

type CheckPool struct {
	// to -> []check
	Data map[common.Address][]*check.Check //有序数组
}

// add a check into pool
func (p *CheckPool) Store(chk *check.Check) error {
	// get slice
	s := p.Data[chk.ToAddr]

	// check already exist
	if s[chk.Nonce] != nil {
		return errors.New("check already exist")
	}

	// put check into right position
	if chk.Nonce+1 > uint64(len(s)) {
		// padding nils
		for n := uint64(len(s)); n < chk.Nonce; n++ {
			s = append(s, nil)
		}
		// right position after nils, and append check
		s = append(s, chk)
		p.Data[chk.ToAddr] = s
		return nil
	}

	return errors.New("exception")
}

type PaycheckPool struct {
	// to -> []paycheck
	Data map[common.Address][]*check.Paycheck //按照nonce有序
}

// 如果paycheck池为空，则直接将paycheck添加到池中。
// 如果此nonce的paycheck已经存在，则比较payvalue以后，替换。
// 如果此nonce的paycheck不存在，则直接放到nonce对应的切片位置。
func (p *PaycheckPool) Store(pc *check.Paycheck) error {

	// get slice
	s := p.Data[pc.Check.ToAddr]

	// if paycheck pool is nil, append it
	if len(s) == 0 {
		s = append(s, pc)
		p.Data[pc.Check.ToAddr] = s
		return nil
	}

	// check already exist
	if s[pc.Check.Nonce] != nil {
		if (pc.PayValue.Cmp(s[pc.Check.Nonce].PayValue)) <= 0 {
			return errors.New("new payvalue must larger than old one")
		}
		// update old paycheck to new one
		s[pc.Check.Nonce] = pc
		return nil

	}

	// update old paycheck to new one
	s[pc.Check.Nonce] = pc
	return nil
}

// get the last paycheck, which has the biggest nonce
func (p *PaycheckPool) GetCurrent(to common.Address) (*check.Paycheck, error) {
	s := p.Data[to]
	return s[len(s)-1], nil
}
