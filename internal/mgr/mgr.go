package mgr

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/check"
)

// order info
type Order struct {
	ID uint64 // 订单ID

	Token common.Address // 货币类型
	Value *big.Int       // 货币数量
	From  common.Address // user地址
	To    common.Address // provider地址

	Time int64 // 订单提交时间

	Name  string // 购买人姓名
	Tel   string // 购买人联系方式
	Email string // 接收支票的邮件地址

	State uint8 // 标记是否已付款; 0,1 paid,2 check
}

func (odr *Order) Equal(o2 *Order) (bool, error) {
	if odr.ID != o2.ID {
		return false, errors.New("id not equal")
	}
	if odr.Token != o2.Token {
		return false, errors.New("token not equal")
	}
	if odr.Value.String() != o2.Value.String() {
		return false, errors.New("value not equal")
	}
	if odr.From != o2.From {
		return false, errors.New("from not equal")
	}
	if odr.To != o2.To {
		return false, errors.New("to not equal")
	}
	if odr.Time != o2.Time {
		return false, errors.New("time not equal")
	}
	if odr.Name != o2.Name {
		return false, errors.New("name not equal")
	}
	if odr.Tel != o2.Tel {
		return false, errors.New("tel not equal")
	}
	if odr.Email != o2.Email {
		return false, errors.New("email not equal")
	}
	if odr.State != o2.State {
		return false, errors.New("state not equal")
	}

	return true, nil
}

type OrderMgr struct {
	ID      uint64                  // ID used for create next order
	OdrPool map[uint64]*Order       // id -> order
	ChkPool map[uint64]*check.Check // id -> check
}

// create a new order manager
func New() *OrderMgr {
	om := &OrderMgr{
		ID:      0,
		OdrPool: make(map[uint64]*Order),
		ChkPool: make(map[uint64]*check.Check),
	}
	return om
}

// get ID for new order, and increase ID by 1
func (odrMgr *OrderMgr) NewID() uint64 {
	id := odrMgr.ID
	odrMgr.ID++
	return id
}

// get an order by id
func (odrMgr *OrderMgr) GetOrder(oid uint64) (*Order, error) {
	if odrMgr.OdrPool[oid] == nil {
		return nil, errors.New("order not exist")
	}
	return odrMgr.OdrPool[oid], nil
}

// store an order into pool
func (odrMgr *OrderMgr) PutOrder(odr *Order) error {
	if odr == nil {
		return errors.New("order is nil")
	}
	odrMgr.OdrPool[odr.ID] = odr
	return nil
}

// get a check from pool by oid
func (odrMgr *OrderMgr) GetCheck(oid uint64) *check.Check {
	return odrMgr.ChkPool[oid]
}

// store a check into pool by oid
func (odrMgr *OrderMgr) PutCheck(oid uint64, chk *check.Check) {
	odrMgr.ChkPool[oid] = chk
}

// get order state
func (odrMgr *OrderMgr) GetState(oid uint64) (uint8, error) {
	odr, err := odrMgr.GetOrder(oid)
	if err != nil {
		return 0, err
	}
	return odr.State, nil
}

// set order state
func (odrMgr *OrderMgr) SetState(oid uint64, st uint8) error {
	odr, err := odrMgr.GetOrder(oid)
	if err != nil {
		return err
	}
	odr.State = st
	return nil
}

// pay process for an specific order
func (odrMgr *OrderMgr) UserPay(oid uint64) {
	// set state paid after user pay money
	// generate a check for user
	// set check to odr.Check
}
