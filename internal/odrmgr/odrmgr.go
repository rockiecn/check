package odrmgr

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
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

	// 0 initial
	// 1 order paid
	// 2 created check
	State uint8 // 标记是否已付款;
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

// serialize an order with cbor
func (odr *Order) Marshal() ([]byte, error) {

	if odr == nil {
		return nil, errors.New("nil order")
	}

	b, err := cbor.Marshal(*odr)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// decode a buf into order
func (odr *Order) UnMarshal(buf []byte) error {
	if odr == nil {
		return errors.New("nil order")
	}
	if buf == nil {
		return errors.New("nil buf")
	}

	err := cbor.Unmarshal(buf, odr)
	if err != nil {
		return err
	}
	return nil
}

type Ordermgr struct {
	ID      uint64                  // ID used for create next order
	OdrPool map[uint64]*Order       // id -> order
	ChkPool map[uint64]*check.Check // id -> check
}

// create a new order manager
func New() *Ordermgr {
	om := &Ordermgr{
		ID:      0,
		OdrPool: make(map[uint64]*Order),
		ChkPool: make(map[uint64]*check.Check),
	}
	return om
}

// get ID for new order, and increase ID by 1
func (mgr *Ordermgr) NewID() uint64 {
	id := mgr.ID
	mgr.ID++
	return id
}

// get an order by id
func (mgr *Ordermgr) GetOrder(oid uint64) (*Order, error) {
	if mgr.OdrPool[oid] == nil {
		return nil, errors.New("order not exist")
	}
	return mgr.OdrPool[oid], nil
}

// store an order into pool
func (mgr *Ordermgr) PutOrder(odr *Order) error {
	if odr == nil {
		return errors.New("order is nil")
	}

	// put into pool
	mgr.OdrPool[odr.ID] = odr

	return nil
}

// delete an order from pool by ID
func (mgr *Ordermgr) DelOrder(oid uint64) {
	delete(mgr.OdrPool, oid)
}

// get a check from pool by oid
func (mgr *Ordermgr) GetCheck(oid uint64) *check.Check {
	return mgr.ChkPool[oid]
}

// store a check into pool by oid
func (mgr *Ordermgr) PutCheck(oid uint64, chk *check.Check) error {
	if chk == nil {
		return errors.New("chk is nil")
	}

	mgr.ChkPool[oid] = chk
	return nil
}

// delete a check from pool by ID
func (mgr *Ordermgr) DelCheck(oid uint64) {
	delete(mgr.ChkPool, oid)
}

// get order state
func (mgr *Ordermgr) GetState(oid uint64) (uint8, error) {
	odr, err := mgr.GetOrder(oid)
	if err != nil {
		return 0, err
	}
	return odr.State, nil
}

// set order state
func (mgr *Ordermgr) SetState(oid uint64, st uint8) error {
	odr, err := mgr.GetOrder(oid)
	if err != nil {
		return err
	}
	odr.State = st
	return nil
}

// pay process for an specific order
func (mgr *Ordermgr) UserPay(oid uint64) {
	// set order state to paid after user paid the order
	// create and store check
}
