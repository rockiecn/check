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

// 如果paycheck队列为空，表示所有的check都能支付，则直接取出check池的第一张支票返回（如果有的话）。
// 如果paycheck队列不为空，则在paycheck队列中取出尾项的nonce（队列中最大，命名为maxNonce）
// 然后在check队列中从maxNonce+1开始向后找，返回找到的第一张支票
// 如果直到队列最后都没找到一张支票，表示已无可用支票，返回空
// get a check that has not been used yet
func (user *User) GetNew(to common.Address) (*check.Check, error) {

	sc := user.PoolA.Data[to]
	sp := user.PoolB.Data[to]

	// check pool is empty
	if len(sc) == 0 {
		return nil, errors.New("check pool is empty")
	}

	// if paycheck pool is empty, return the first check
	if len(sp) == 0 {
		if sc[0] != nil {
			return sc[0], nil
		} else {
			return nil, errors.New("first check is nil")
		}
	}

	// here we can sure len(sc) and len(sp) not 0
	maxCNonce := len(sc) - 1
	maxPCNonce := len(sp) - 1

	// last check is used for paycheck
	if maxCNonce <= maxPCNonce {
		return nil, errors.New("no usable check in pool")
	}

	// search for an usable check in check pool
	for i := maxPCNonce + 1; i <= maxCNonce; i++ {
		if sc[i] != nil {
			return sc[i], nil
		}
	}

	// no check found in check pool
	return nil, errors.New("check nonce larger than paycheck nonce, but no check is found")
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

	return pchk, nil
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
		return false, nil
	}

	// from = user
	if chk.FromAddr != user.UserAddr {
		return false, nil
	}

	// nonce must bigger than contract nonce
	contractNonce, err := comn.QueryNonce(user.UserAddr, chk.ContractAddr, chk.ToAddr)
	if err != nil {
		return false, err
	}
	if chk.Nonce <= contractNonce {
		return false, nil
	}

	// check must not exist in pool
	sc := user.PoolA.Data[chk.ToAddr]
	if sc[chk.Nonce] != nil {
		return false, nil
	}

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

	// if nonce is out of boundary, extend pool and put check into right position
	if chk.Nonce >= uint64(len(s)) {
		// padding nils
		pad := chk.Nonce - uint64(len(s))
		for i := uint64(0); i < pad; i++ {
			s = append(s, nil)
		}
		// right position after nils, and append check
		if chk.Nonce == uint64(len(s)) {
			s = append(s, chk)
			p.Data[chk.ToAddr] = s
			return nil
		} else {
			return errors.New("bad check nonce")
		}
	}

	// check in boundary, put check into nonce position
	s[chk.Nonce] = chk
	return nil
}

type PaycheckPool struct {
	// to -> []paycheck
	Data map[common.Address][]*check.Paycheck //按照nonce有序
}

// 如果nonce越界（使用新check支付），则先使用nil值填充队列，直到nonce所在的位置，然后把paycheck放到这里
// 如果nonce没有越界（使用current支票支付），则直接用它替换nonce位置的paycheck
func (p *PaycheckPool) Store(pc *check.Paycheck) error {

	// get slice
	s := p.Data[pc.Check.ToAddr]

	// if nonce is out of boundary, extend pool and put check into right position
	if pc.Check.Nonce >= uint64(len(s)) {
		// padding nils
		pad := pc.Check.Nonce - uint64(len(s))
		for i := uint64(0); i < pad; i++ {
			s = append(s, nil)
		}
		// right position to append check
		if pc.Check.Nonce == uint64(len(s)) {
			s = append(s, pc)
			p.Data[pc.Check.ToAddr] = s
			return nil
		} else {
			return errors.New("bad check nonce")
		}
	}

	// check in boundary, put check into nonce position
	s[pc.Check.Nonce] = pc
	return nil
}

// get the last paycheck, which has the biggest nonce
func (p *PaycheckPool) GetCurrent(to common.Address) (*check.Paycheck, error) {
	s := p.Data[to]
	if len(s) == 0 {
		return nil, nil
	} else {
		return s[len(s)-1], nil
	}
}
