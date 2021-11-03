package operator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rockiecn/check/internal/cash"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/utils"
)

type Operator struct {
	OpSK         string
	OpAddr       common.Address
	ContractAddr common.Address
	Nonces       map[common.Address]uint64 // nonce for next check
	OdrMgr       *OrderMgr
}

type IOperator interface {
	Deploy(value *big.Int) (*types.Transaction, common.Address, error)
	QueryBalance() (*big.Int, error)
	Deposit(value *big.Int) (*types.Transaction, error)
	GetNonce(to common.Address) (uint64, error)

	QueryOrder(order)
	GenCheck(oid uint64) (*check.Check, error))
}

// create an operator without contract.
// a contract should be deployed after this.
func NewOperator(sk string, token string) (IOperator, error) {
	addr, err := utils.KeyToAddr(sk)
	if err != nil {
		return nil, err
	}
	op := &Operator{
		OpSK:   sk,
		OpAddr: addr,
		Nonces: make(map[common.Address]uint64),
		OdrMgr: new(OrderMgr),
	}

	return op, nil
}

// value: money to new contract
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

	op.SetCtrAddr(addr)
	/*
		go func() {
			// deploy contract, wait for mining.
			for {
				txReceipt, _ := ethClient.TransactionReceipt(context.Background(), tx.Hash())
				// receipt ok
				if txReceipt != nil {
					break
				}
				fmt.Println("deploy mining..")
				time.Sleep(time.Duration(5) * time.Second)
			}
		}()
	*/
	return tx, addr, nil
}

// query balance of contract
func (op *Operator) QueryBalance() (*big.Int, error) {
	ethClient, err := utils.GetClient(utils.HOST)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	auth := new(bind.CallOpts)
	auth.From = op.OpAddr

	// get contract instance from address
	cashInstance, err := cash.NewCash(op.ContractAddr, ethClient)
	if err != nil {
		return nil, errors.New("newcash failed")
	}

	bal, err := cashInstance.GetBalance(auth)
	if err != nil {
		return nil, errors.New("tx failed")
	}

	return bal, nil
}

// get the nonce of a given provider in contract
func (op *Operator) GetNonce(to common.Address) (uint64, error) {
	nonce, err := utils.GetNonce(op.ContractAddr, to)
	if err != nil {
		return 0, err
	}

	return nonce, err
}

// set nonce of contract
func (op *Operator) SetNonce(to common.Address, nonce uint64) (*types.Transaction, error) {
	ethClient, err := utils.GetClient(utils.HOST)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	auth, err := utils.MakeAuth(op.OpSK, nil, nil, nil, 9000000)
	if err != nil {
		return nil, err
	}

	// get contract instance from address
	cashInstance, err := cash.NewCash(op.ContractAddr, ethClient)
	if err != nil {
		return nil, errors.New("newcash failed")
	}

	tx, err := cashInstance.SetNonce(auth, to, nonce)
	if err != nil {
		return nil, errors.New("tx failed")
	}

	return tx, nil
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
	cashInstance, err := cash.NewCash(op.ContractAddr, ethClient)
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
// last, put the check into order
func (op *Operator) GenCheck(oid uint64) (*check.Check, error) {

	odr := op.OdrMgr.GetOrder(oid)
	nonce := op.Nonces[odr.To]

	chk := &check.Check{
		Value:        odr.Value,
		TokenAddr:    odr.Token,
		FromAddr:     odr.From,
		ToAddr:       odr.To,
		Nonce:        nonce,
		OpAddr:       op.OpAddr,
		ContractAddr: op.ContractAddr,
	}

	// signed by operator
	err := chk.Sign(op.OpSK)
	if err != nil {
		return nil, err
	}

	// assign check for this order
	op.OdrMgr.PutCheck(oid, chk)

	// update nonce for next check
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

		if v.Check.Nonce < minNonce {
			minNonce = v.Check.Nonce
		}
		if v.Check.Nonce > maxNonce {
			maxNonce = v.Check.Nonce
		}

	}

	batch := new(check.BatchCheck)
	batch.OpAddr = op.OpAddr
	batch.ToAddr = toAddr
	batch.BatchValue = total
	batch.MinNonce = minNonce
	batch.MaxNonce = maxNonce
	err := batch.Sign(op.OpSK)
	if err != nil {
		return nil, err
	}

	return batch, nil
}

func (op *Operator) SetCtrAddr(addr common.Address) {
	op.ContractAddr = addr
}

func (op *Operator) WaitForMiner(txHash *types.Transaction) error {
	// connect to geth
	ethClient, err := utils.GetClient(utils.HOST)
	if err != nil {
		return err
	}
	defer ethClient.Close()

	for {
		txReceipt, _ := ethClient.TransactionReceipt(context.Background(), txHash.Hash())
		// receipt ok
		if txReceipt != nil {
			break
		}
		fmt.Println("waiting for miner, 5 seconds..")
		time.Sleep(time.Duration(5) * time.Second)
	}
	return nil
}

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
