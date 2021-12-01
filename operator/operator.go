package operator

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rockiecn/check/internal/cash"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/db"
	"github.com/rockiecn/check/internal/odrmgr"
	"github.com/rockiecn/check/internal/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

type Operator struct {
	OpSK    string
	OpAddr  common.Address
	CtrAddr common.Address

	Nonces map[common.Address]uint64 // nonce for next check

	orderDB string // dbfile for order
	checkDB string // dbfile for check

	OM *odrmgr.Ordermgr
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

// create an operator out of sk
func New(sk string) (IOperator, error) {
	opAddr, err := utils.SkToAddr(sk)
	if err != nil {
		return nil, err
	}

	op := &Operator{
		OpSK:    sk,
		OpAddr:  opAddr,
		Nonces:  make(map[common.Address]uint64),
		orderDB: "./order.db",
		checkDB: "./check.db",
		OM:      odrmgr.New(),
	}

	return op, nil
}

// Store an order into order db
func (op *Operator) StoreOrder(odr *odrmgr.Order) error {
	// serialize
	b, err := odr.Marshal()
	if err != nil {
		return err
	}
	// write db
	err = db.WriteDB(op.orderDB, utils.Uint64ToByte(odr.ID), b)
	if err != nil {
		return err
	}

	return nil
}

// restore orders from order db
func (op *Operator) RestoreOrder() error {
	db, err := leveldb.OpenFile(op.orderDB, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	// read data from db
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		//k := iter.Key()
		v := iter.Value()
		odr := &odrmgr.Order{}
		err := odr.UnMarshal(v)
		if err != nil {
			return err
		}

		// put into memory
		err = op.OM.PutOrder(odr)
		if err != nil {
			return err
		}
	}

	iter.Release()
	err = iter.Error()
	if err != nil {
		return err
	}

	return nil
}

// store a check into check db
func (op *Operator) StoreChk(oid uint64, chk *check.Check) error {
	// serialize
	b, err := chk.Marshal()
	if err != nil {
		return err
	}
	// write db
	err = db.WriteDB(op.checkDB, utils.Uint64ToByte(oid), b)
	if err != nil {
		return err
	}

	return nil
}

// restore checks from check db
func (op *Operator) RestoreChk() error {
	db, err := leveldb.OpenFile(op.checkDB, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	// read data from db
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		chk := &check.Check{}
		err := chk.UnMarshal(v)
		if err != nil {
			return err
		}

		oid := utils.ByteToUint64(k)
		// put into memory
		err = op.OM.PutCheck(oid, chk)
		if err != nil {
			return err
		}
	}

	iter.Release()
	err = iter.Error()
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
	auth, err := utils.MakeAuth(op.OpSK, value, bigNonce, gasPrice, 9000000)
	if err != nil {
		return nil, common.Address{}, err
	}

	addr, tx, _, err := cash.DeployCash(auth, ethClient)
	if err != nil {
		return nil, common.Address{}, err
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

// generate a new check for an order
// first get order with oid
// then generate a check from order info
// increase next check nonce by 1
func (op *Operator) CreateCheck(oid uint64) (*check.Check, error) {

	odr, err := op.OM.GetOrder(oid)
	if err != nil {
		return nil, err
	}

	nonce := op.Nonces[odr.To]

	chk := &check.Check{
		Value:     odr.Value,
		TokenAddr: odr.Token,
		FromAddr:  odr.From,
		ToAddr:    odr.To,
		Nonce:     nonce,
		OpAddr:    op.OpAddr,
		CtrAddr:   op.CtrAddr,
	}

	// signed by operator
	err = chk.Sign(op.OpSK)
	if err != nil {
		return nil, err
	}

	// increase nonce by 1
	op.Nonces[odr.To] = nonce + 1

	return chk, nil
}

// aggregate a batch of paychecks into a single BatchCheck
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

	batch := &check.BatchCheck{}
	batch.OpAddr = op.OpAddr
	batch.ToAddr = toAddr
	batch.CtrAddr = op.CtrAddr
	batch.BatchValue = total
	batch.MinNonce = minNonce
	batch.MaxNonce = maxNonce
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
