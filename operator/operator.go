package operator

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
	DeployContract(value *big.Int) (*types.Transaction, common.Address, error)
	QueryBalance() (*big.Int, error)
	GetNonce(to common.Address) (uint64, error)
	Deposit(value *big.Int) (*types.Transaction, error)

	GenCheck(oid uint64) (*check.Check, error)
	Aggregate(pcs []*check.Paycheck) (*check.BatchCheck, error)
}

// create an operator, and a contract is deployed.
func NewOperator(sk string, token string) (IOperator, *types.Transaction, error) {
	op := &Operator{
		OpSK:   sk,
		OpAddr: utils.KeyToAddr(sk),
		Nonces: make(map[common.Address]uint64),
		OdrMgr: new(OrderMgr),
	}

	/*
		// deploy new contract, give 20 eth to it
		tx, addr, err := op.DeployContract(comn.String2BigInt("20000000000000000000"))
		if err != nil {
			return nil, nil, err
		}
		op.ContractAddr = addr

		return op, tx, nil
	*/
	return op, nil, nil
}

// value: money to new contract
func (op *Operator) DeployContract(value *big.Int) (tx *types.Transaction, contractAddr common.Address, err error) {

	// connect to node
	ethClient, err := utils.GetClient(utils.HOST)
	if err != nil {
		return nil, common.Address{}, err
	}
	defer ethClient.Close()

	// string to ecdsa
	priKeyECDSA, err := crypto.HexToECDSA(op.OpSK)
	if err != nil {
		return nil, common.Address{}, err
	}

	// get pubkey
	pubKey := priKeyECDSA.Public()
	// ecdsa
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, common.Address{}, errors.New("error casting public key to ECDSA")
	}
	// get operator address
	opComAddr := crypto.PubkeyToAddress(*pubKeyECDSA)
	// get nonce
	nonce, err := ethClient.PendingNonceAt(context.Background(), opComAddr)
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

	contractAddr, tx, _, err = cash.DeployCash(auth, ethClient)
	if err != nil {
		return nil, common.Address{}, err
	}
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
	return tx, contractAddr, nil
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

	var toAddr common.Address
	total := new(big.Int)
	minNonce := ^uint64(0)
	maxNonce := uint64(0)
	batch := new(check.BatchCheck)
	for _, v := range pcs {
		// verify paycheck sig
		ok, err := v.Verify()
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("paycheck sig verify failed")
		}

		// verify check sig
		ok, err = v.Check.Verify()
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("check sig verify failed")
		}

		// payvalue must not bigger than value
		if v.PayValue.Cmp(v.Check.Value) > 0 {
			return nil, errors.New("payvalue bigger than value")
		}

		if toAddr.String() == "" {
			toAddr = v.Check.ToAddr
		} else {
			if toAddr != v.Check.ToAddr {
				return nil, errors.New("to address not identical")
			}
		}

		// aggregate payvalue
		total.Add(total, v.PayValue)

		if v.Check.Nonce < minNonce {
			minNonce = v.Check.Nonce
		}
		if v.Check.Nonce > maxNonce {
			maxNonce = v.Check.Nonce
		}

		batch.OpAddr = op.OpAddr
		batch.ToAddr = v.Check.ToAddr
		batch.BatchValue = total
		batch.MinNonce = minNonce
		batch.MaxNonce = maxNonce
		err = batch.Sign(op.OpSK)
		if err != nil {
			return nil, err
		}
	}

	return batch, nil
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
