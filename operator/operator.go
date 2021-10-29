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

type Operator struct {
	OpSK         string
	OpAddr       common.Address
	ContractAddr common.Address
	// to -> nonce
	Nonces map[common.Address]uint64

	OdrMgr *OrderMgr
}

type IOperator interface {
	DeployContract(value *big.Int) (*types.Transaction, common.Address, error)
	// query current balance of contract
	QueryBalance() (*big.Int, error)
	// get contract nonce of a provider
	GetNonce(to common.Address) (uint64, error)
	// give money to contract
	Deposit(value *big.Int) (*types.Transaction, error)

	// generate a check from order id
	GenCheck(oid uint64) (*check.Check, error)
}

// new operator, a contract is deployed.
func New(sk string, token string) (IOperator, *types.Transaction, error) {
	op := &Operator{
		OpSK:   sk,
		OpAddr: utils.KeyToAddr(sk),
		Nonces: make(map[common.Address]uint64),
		OdrMgr: new(OrderMgr),
	}

	/*
		// give 20 eth to new contract
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

// query balance of the contract
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

// get nonce of a given provider
func (op *Operator) GetNonce(to common.Address) (uint64, error) {
	nonce, err := utils.GetNonce(op.ContractAddr, to)
	return nonce, err
}

// deposit some money into contract
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

// generate a new check with order id
// first get order with oid
// then generate a check from order info
// last, put the check into order
func (op *Operator) GenCheck(oid uint64) (*check.Check, error) {

	odr := op.OdrMgr.GetOrderByID(oid)
	newNonce := op.Nonces[odr.To]

	chk := &check.Check{
		Value:        odr.Value,
		TokenAddr:    odr.Token,
		FromAddr:     odr.From,
		ToAddr:       odr.To,
		Nonce:        newNonce,
		OpAddr:       op.OpAddr,
		ContractAddr: op.ContractAddr,
	}

	// signed by operator
	err := chk.Sign(op.OpSK)
	if err != nil {
		return nil, err
	}

	// put new check into order
	op.OdrMgr.SetCheckByID(oid, chk)

	// update nonce
	op.Nonces[odr.To] = newNonce + 1

	return chk, nil
}

// mutli paycheck to a bacthCheck
func (op *Operator) Aggregate(pcs []*check.Paycheck) (*check.BatchCheck, error) {
	return nil, nil
}
