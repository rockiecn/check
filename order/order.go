package order

import (
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
)

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

func (odrMgr *OrderMgr) CurrentID() uint64 {
	return odrMgr.ID
}

func (odrMgr *OrderMgr) UpdateID() {
	odrMgr.ID++
}

func (odrMgr *OrderMgr) GetOrderByID(oid uint64) *Order {
	return odrMgr.Pool[oid]
}

func (odrMgr *OrderMgr) PutOrder(odr *Order) error {
	if odr != nil {
		odrMgr.Pool[odr.ID] = odr
		return nil
	}
	return errors.New("order is nil")
}

func (odrMgr *OrderMgr) GetCheckByID(oid uint64) *check.Check {
	return odrMgr.GetOrderByID(oid).Check
}
func (odrMgr *OrderMgr) SetCheckByID(oid uint64, chk *check.Check) {
	odrMgr.GetOrderByID(oid).Check = chk
}

func (odrMgr *OrderMgr) GetStateByID(oid uint64) uint8 {
	return odrMgr.GetOrderByID(oid).State
}

func (odrMgr *OrderMgr) SetStateByID(oid uint64, s uint8) {
	odrMgr.GetOrderByID(oid).State = s
}

func (odrMgr *OrderMgr) UserPay(oid uint64) {
	// paid
	// check
	// add check to om.chks
}
