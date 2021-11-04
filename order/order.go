package order

import (
	"errors"
	"math/big"
	"time"

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

	Time time.Time // 订单提交时间

	Name  string // 购买人姓名
	Tel   string // 购买人联系方式
	Email string // 接收支票的邮件地址

	State uint8 // 标记是否已付款; 0,1 paid,2 check

	Check *check.Check // 根据此订单生成的支票
}

type OrderMgr struct {
	ID   uint64            // ID used for next order
	Pool map[uint64]*Order // id -> order
}

// get ID for new order, and increase ID by 1
func (odrMgr *OrderMgr) NewID() uint64 {
	id := odrMgr.ID
	odrMgr.ID++
	return id
}

// get an order by id
func (odrMgr *OrderMgr) GetOrder(oid uint64) *Order {
	return odrMgr.Pool[oid]
}

// put an order into pool
func (odrMgr *OrderMgr) PutOrder(odr *Order) error {
	if odr != nil {
		odrMgr.Pool[odr.ID] = odr
		return nil
	}
	return errors.New("order is nil")
}

// get the check of an order
func (odrMgr *OrderMgr) GetCheck(oid uint64) *check.Check {
	return odrMgr.GetOrder(oid).Check
}

// assign a check for an order by order id
func (odrMgr *OrderMgr) PutCheck(oid uint64, chk *check.Check) {
	odrMgr.GetOrder(oid).Check = chk
}

// get order state
func (odrMgr *OrderMgr) GetState(oid uint64) uint8 {
	return odrMgr.GetOrder(oid).State
}

// set order state
func (odrMgr *OrderMgr) SetState(oid uint64, st uint8) {
	odrMgr.GetOrder(oid).State = st
}

// pay process for an specific order
func (odrMgr *OrderMgr) UserPay(oid uint64) {
	// set state paid after user pay money
	// generate a check for user
	// set check to odr.Check
}
