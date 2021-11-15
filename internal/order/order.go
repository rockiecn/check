package order

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
	"github.com/rockiecn/check/internal/check"
	"github.com/syndtr/goleveldb/leveldb"
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
}

type OrderMgr struct {
	ID      uint64                  // ID used for create next order
	OdrPool map[uint64]*Order       // id -> order
	ChkPool map[uint64]*check.Check // id -> check
}

// create a new order
func NewOdr(
	ID uint64,
	token common.Address,
	from common.Address,
	to common.Address,
	value *big.Int,
	t time.Time,
	name string,
	tel string,
	email string,
	state uint8,
) *Order {
	odr := &Order{
		ID:    ID,
		Token: token,
		From:  from,
		To:    to,
		Value: value,
		Time:  t,
		Name:  name,
		Tel:   tel,
		Email: email,
		State: state,
	}
	return odr
}

// create a new order manager
func NewMgr() *OrderMgr {
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

// serialize an order with cbor
func (odrMgr *OrderMgr) MarshOdr(odr *Order) ([]byte, error) {

	if odr == nil {
		return nil, errors.New("nil order")
	}

	b, err := cbor.Marshal(*odr)
	if err != nil {
		fmt.Println("error:", err)
	}
	return b, nil
}

// decode a buf into an order
func (odrMgr *OrderMgr) UnMarshOdr(buf []byte) (*Order, error) {
	if buf == nil {
		return nil, errors.New("nil buf")
	}

	odr := new(Order)
	err := cbor.Unmarshal(buf, odr)
	if err != nil {
		fmt.Println("error:", err)
	}
	return odr, nil
}

func (odrMgr *OrderMgr) WriteDB(oid uint64, buf []byte) error {
	db, err := leveldb.OpenFile("./order.db", nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return err
	}
	defer db.Close()

	// uint64 to []byte
	k := make([]byte, 8)
	binary.LittleEndian.PutUint64(k, oid)

	err = db.Put(k, buf, nil)
	if err != nil {
		return err
	}
	return nil
}

func (odrMgr *OrderMgr) ReadDB(oid uint64) ([]byte, error) {
	db, err := leveldb.OpenFile("./order.db", nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return nil, err
	}
	defer db.Close()

	// uint64 to []byte
	k := make([]byte, 8)
	binary.LittleEndian.PutUint64(k, oid)

	buf, err := db.Get(k, nil)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
