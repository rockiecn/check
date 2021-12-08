package operator

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fxamacker/cbor/v2"
	"github.com/rockiecn/check/internal/cash"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/store"
	"github.com/rockiecn/check/internal/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

type Operator struct {
	OpSK    string
	OpAddr  common.Address
	CtrAddr common.Address

	// nonce for next check
	Nonces map[common.Address]uint64

	// oid -> order
	OdrStorer store.Storer
	// oid -> check
	ChkStorer store.Storer

	OM *Ordermgr
}

type IOperator interface {
	Deploy(value *big.Int) (*types.Transaction, common.Address, error)
	SetCtrAddr(addr common.Address)
	QueryBalance() (*big.Int, error)
	Deposit(value *big.Int) (*types.Transaction, error)
	GetNonce(to common.Address) (uint64, error)
	CreateCheck(oid uint64) (*check.Check, error)
	Aggregate(pcs []*check.Paycheck) (*check.BatchCheck, error)
}

// create an operator out of sk, order dbfile, check dbfile
func New(sk string, odrDBfile string, chkDBfile string) (IOperator, error) {
	opAddr, err := utils.SkToAddr(sk)
	if err != nil {
		return nil, err
	}

	op := &Operator{
		OpSK:   sk,
		OpAddr: opAddr,
		Nonces: make(map[common.Address]uint64),
		OM:     NewMgr(),
	}

	// open order db
	var OdrStore = &store.Store{}
	OdrStore.DB, err = leveldb.OpenFile(odrDBfile, nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return nil, err
	}
	// open check db
	var ChkStore = &store.Store{}
	ChkStore.DB, err = leveldb.OpenFile(chkDBfile, nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return nil, err
	}

	op.ChkStorer = OdrStore
	op.OdrStorer = ChkStore

	return op, nil
}

// Store an order into order db
func (op *Operator) StoreOrder(odr *Order) error {
	// serialize
	b, err := odr.Serialize()
	if err != nil {
		return err
	}

	// write into db
	err = op.OdrStorer.Put(utils.Uint64ToByte(odr.ID), b)
	if err != nil {
		return err
	}

	return nil
}

// restore an order from order db
func (op *Operator) RestoreOrder(oid uint64) error {

	// get order from db with oid
	k := utils.Uint64ToByte(oid)
	v, err := op.OdrStorer.Get(k)
	if err != nil {
		return err
	}
	// deserialize order
	odr := &Order{}
	err = odr.DeSerialize(v)
	if err != nil {
		return err
	}
	// put into pool
	err = op.PutOrder(odr)
	if err != nil {
		return err
	}

	return nil
}

// store a check into check db
func (op *Operator) StoreChk(oid uint64, chk *check.Check) error {
	// serialize
	b, err := chk.Serialize()
	if err != nil {
		return err
	}

	// write db
	err = op.ChkStorer.Put(utils.Uint64ToByte(oid), b)
	if err != nil {
		return err
	}

	return nil
}

// restore a check from check db with oid
func (op *Operator) RestoreChk(oid uint64) error {

	// get check from db with oid
	k := utils.Uint64ToByte(oid)
	v, err := op.ChkStorer.Get(k)
	if err != nil {
		return err
	}
	// deserialize check
	chk := &check.Check{}
	err = chk.DeSerialize(v)
	if err != nil {
		return err
	}
	// put into pool
	err = op.PutCheck(oid, chk)
	if err != nil {
		return err
	}

	return nil
}

// value: the money given to new contract
func (op *Operator) Deploy(value *big.Int) (*types.Transaction, common.Address, error) {

	// connect to node
	ethClient, err := utils.GetClient(utils.HOST)
	if err != nil {
		return nil, common.Address{}, err
	}
	defer ethClient.Close()

	// get nonce
	nonce, err := ethClient.PendingNonceAt(context.Background(), op.OpAddr)
	if err != nil {
		return nil, common.Address{}, err
	}

	// get gas price
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, common.Address{}, err
	}

	// transfer to big.Int for contract
	bigNonce := new(big.Int).SetUint64(nonce)
	auth, err := utils.MakeAuth(op.OpSK, value, bigNonce, gasPrice, utils.GasL)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("MakeAuth- %v", err)
	}

	addr, tx, _, err := cash.DeployCash(auth, ethClient)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("DeployCash- %v", err)
	}

	return tx, addr, nil
}

// query the balance of contract
func (op *Operator) QueryBalance() (*big.Int, error) {
	ethClient, err := utils.GetClient(utils.HOST)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	auth := new(bind.CallOpts)
	auth.From = op.OpAddr

	// get contract instance from address
	cashInstance, err := cash.NewCash(op.CtrAddr, ethClient)
	if err != nil {
		return nil, errors.New("newcash failed")
	}

	bal, err := cashInstance.GetBalance(auth)
	if err != nil {
		return nil, errors.New("tx failed")
	}

	return bal, nil
}

// GetNonce: get the nonce of a given provider in contract
func (op *Operator) GetNonce(to common.Address) (uint64, error) {
	nonce, err := utils.GetCtNonce(op.CtrAddr, to)
	if err != nil {
		return 0, err
	}

	return nonce, err
}

// deposit some money to contract
func (op *Operator) Deposit(value *big.Int) (*types.Transaction, error) {
	ethClient, err := utils.GetClient(utils.HOST)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	auth, err := utils.MakeAuth(op.OpSK, nil, nil, big.NewInt(1000), 9000000)
	if err != nil {
		return nil, errors.New("make auth failed")
	}
	// money to deposit
	auth.Value = value

	// get contract instance from address
	cashInstance, err := cash.NewCash(op.CtrAddr, ethClient)
	if err != nil {
		return nil, errors.New("newcash failed")
	}

	// call contract
	tx, err := cashInstance.Deposit(auth)
	if err != nil {
		return nil, errors.New("tx failed")
	}

	return tx, nil
}

// CreateCheck - generate a new check for an order
// 1 get order from pool with oid
// 2 generate a check from order
// 3 update nonce by 1
// Things should be followed after create a check:
// 1 put the check into pool
// 2 store the check into db
func (op *Operator) CreateCheck(oid uint64) (*check.Check, error) {

	// get order by id
	odr, err := op.GetOrder(oid)
	if err != nil {
		return nil, err
	}

	nonce := op.Nonces[odr.To]

	// create check
	chk := &check.Check{
		CheckInfo: check.CheckInfo{
			Value:     odr.Value,
			TokenAddr: odr.Token,
			FromAddr:  odr.From,
			ToAddr:    odr.To,
			Nonce:     nonce,
			OpAddr:    op.OpAddr,
			CtrAddr:   op.CtrAddr,
		},
	}

	// sign by operator
	err = chk.Sign(op.OpSK)
	if err != nil {
		return nil, err
	}

	// update nonces
	op.Nonces[odr.To] = nonce + 1

	return chk, nil
}

// Aggregate - aggregate a batch of paychecks into a single BatchCheck
func (op *Operator) Aggregate(pcs []*check.Paycheck) (*check.BatchCheck, error) {
	if len(pcs) == 0 {
		return nil, errors.New("no paycheck in data")
	}

	toAddr := pcs[0].Check.ToAddr
	total := new(big.Int)
	minNonce := ^uint64(0)
	maxNonce := uint64(0)

	for _, v := range pcs {
		// verify operator address
		if v.Check.OpAddr != op.OpAddr {
			return nil, errors.New("illegal operator address detected")
		}

		// verify check sig
		ok := v.Check.Verify()
		if !ok {
			return nil, errors.New("check sig verify failed")
		}

		// verify paycheck sig
		ok = v.Verify()
		if !ok {
			return nil, errors.New("paycheck sig verify failed")
		}

		// payvalue must not bigger than value
		if v.PayValue.Cmp(v.Check.Value) > 0 {
			return nil, errors.New("payvalue exceed value")
		}

		// verify toaddr identical
		if v.Check.ToAddr != toAddr {
			return nil, errors.New("to address not identical")
		}

		// aggregate payvalue
		total.Add(total, v.PayValue)

		// update minNonce and maxNonce
		if v.Check.Nonce < minNonce {
			minNonce = v.Check.Nonce
		}
		if v.Check.Nonce > maxNonce {
			maxNonce = v.Check.Nonce
		}

	}

	batch := &check.BatchCheck{
		OpAddr:     op.OpAddr,
		ToAddr:     toAddr,
		CtrAddr:    op.CtrAddr,
		BatchValue: total,
		MinNonce:   minNonce,
		MaxNonce:   maxNonce,
	}
	err := batch.Sign(op.OpSK)
	if err != nil {
		return nil, err
	}

	return batch, nil
}

// set operator's contract address
func (op *Operator) SetCtrAddr(addr common.Address) {
	op.CtrAddr = addr
}

/*
// show all checks in pool
func (op *Operator) ShowChkPool() {
	for k, v := range op.OM.ChkPool {
		fmt.Println("-> oid:", k)
		fmt.Println("check info:")
		fmt.Println(v)
	}
}

// show all orders in pool
func (op *Operator) ShowOdrPool() {
	for k, v := range op.OM.OdrPool {
		fmt.Println("-> oid:", k)
		fmt.Println("order info:")
		fmt.Println(v)
	}
}
*/

// get an order by id
func (op *Operator) GetOrder(oid uint64) (*Order, error) {
	// if order not in pool, read it from db
	if op.OM.OdrPool[oid] == nil {
		k := utils.Uint64ToByte(oid)
		b, err := op.OdrStorer.Get(k)
		if err != nil {
			return nil, err
		}
		odr := &Order{}
		err = odr.DeSerialize(b)
		if err != nil {
			return nil, err
		}

		//put into pool
		err = op.PutOrder(odr)
		if err != nil {
			return nil, err
		}
		return odr, nil
	}

	// get from pool
	return op.OM.OdrPool[oid], nil
}

// put an order into pool, then store into db
func (op *Operator) PutOrder(odr *Order) error {
	if odr == nil {
		return errors.New("order is nil")
	}

	// put order into pool
	op.OM.OdrPool[odr.ID] = odr

	return nil
}

// delete an order by ID
func (op *Operator) DelOrder(oid uint64) error {
	// delete from pool
	delete(op.OM.OdrPool, oid)

	// delete from db
	err := op.OdrStorer.Delete(utils.Uint64ToByte(oid))
	if err != nil {
		return err
	}

	return nil
}

// get a check from pool by oid
func (op *Operator) GetCheck(oid uint64) (*check.Check, error) {
	// if check not in pool, read it from db
	if op.OM.ChkPool[oid] == nil {
		k := utils.Uint64ToByte(oid)
		b, err := op.ChkStorer.Get(k)
		if err != nil {
			return nil, err
		}
		chk := &check.Check{}
		err = chk.DeSerialize(b)
		if err != nil {
			return nil, err
		}

		//put check into pool
		err = op.PutCheck(oid, chk)
		if err != nil {
			return nil, err
		}
		return chk, nil
	}

	// get from pool
	return op.OM.ChkPool[oid], nil
}

// store a check into pool by oid
func (op *Operator) PutCheck(oid uint64, chk *check.Check) error {
	if chk == nil {
		return errors.New("check is nil")
	}

	// put check into pool
	op.OM.ChkPool[oid] = chk

	return nil
}

// delete a check by ID
func (op *Operator) DelCheck(oid uint64) error {
	// delete from pool
	delete(op.OM.ChkPool, oid)

	// delete from db
	err := op.ChkStorer.Delete(utils.Uint64ToByte(oid))
	if err != nil {
		return err
	}

	return nil
}

// get order state
func (op *Operator) GetState(oid uint64) (uint8, error) {
	odr, err := op.GetOrder(oid)
	if err != nil {
		return 0, err
	}
	return odr.State, nil
}

// set order state
func (op *Operator) SetState(oid uint64, st uint8) error {
	odr, err := op.GetOrder(oid)
	if err != nil {
		return err
	}
	odr.State = st
	return nil
}

// pay process for an specific order
func (op *Operator) UserPay(oid uint64) {
	// set order state to paid after user paid the order
	// create and store check
}

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

// compare orders
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
func (odr *Order) Serialize() ([]byte, error) {

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
func (odr *Order) DeSerialize(buf []byte) error {
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
func NewMgr() *Ordermgr {
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
