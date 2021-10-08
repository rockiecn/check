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
}

type IUser interface {
	GenPaycheck(chk *check.Check, payValue *big.Int) (*check.Paycheck, error)

	// TODO:
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

	// record pchk into data
	// err = user.PChkRecorder.Record(pchk)
	// if err != nil {
	// 	return nil, errors.New("record paycheck failed")
	// }

	return pchk, nil
}

// verify received check
func (user *User) Verify(chk *check.Check) (uint64, error) {
	return 0, nil
}

func (user *User) Pay(dataValue *big.Int) error {
	return nil
}

type CheckPool struct {
	// to -> []check
	Pool map[common.Address][]*check.Check //有序数组
}

func (p *CheckPool) Store(c *check.Check) error {
	return nil
}

func (p *CheckPool) GetVirgin(to common.Address) (*check.Check, error) {
	return nil, nil
}

type PaycheckPool struct {
	// to -> []paycheck
	Pool map[common.Address][]*check.Paycheck //按照nonce和payvalue有序
}

func (p *PaycheckPool) Store(pc *check.Paycheck) error {
	return nil
}

func (p *PaycheckPool) GetLast(to common.Address) (*check.Paycheck, error) {
	return nil, nil
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
