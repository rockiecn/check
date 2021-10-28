package operator

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rockiecn/check/cash"
	"github.com/rockiecn/check/check"
	"github.com/rockiecn/check/internal"
	"github.com/rockiecn/check/order"
)

type Operator struct {
	OpSK         string
	OpAddr       common.Address
	ContractAddr common.Address
	// to -> nonce
	Nonces map[common.Address]uint64

	OdrMgr *order.OrderMgr
}

type IOperator interface {
	DeployContract(value *big.Int) (*types.Transaction, common.Address, error)
	// query current balance of contract
	QueryBalance() (*big.Int, error)
	// query contract nonce of a provider
	QueryNonce(to common.Address) (uint64, error)
	// give money to contract
	Deposit(value *big.Int) (*types.Transaction, error)

	// generate a check from order id
	GenCheck(oid uint64) (*check.Check, error)
}

// new operator, a contract is deployed.
func New(sk string, token string) (IOperator, *types.Transaction, error) {
	op := &Operator{
		OpSK:   sk,
		OpAddr: internal.KeyToAddr(sk),
		Nonces: make(map[common.Address]uint64),
		OdrMgr: new(order.OrderMgr),
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
	ethClient, err := internal.GetClient(internal.HOST)
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
	auth, err := internal.MakeAuth(op.OpSK, value, bigNonce, gasPrice, 9000000)
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
	ethClient, err := internal.GetClient(internal.HOST)
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

// query nonce of a given provider
func (op *Operator) QueryNonce(to common.Address) (uint64, error) {
	nonce, err := internal.QueryNonce(op.OpAddr, op.ContractAddr, to)
	return nonce, err
}

// deposit some money into contract
func (op *Operator) Deposit(value *big.Int) (*types.Transaction, error) {
	ethClient, err := internal.GetClient(internal.HOST)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	auth, err := internal.MakeAuth(op.OpSK, nil, nil, big.NewInt(1000), 9000000)
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
