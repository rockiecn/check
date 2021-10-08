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
	"github.com/rockiecn/check/cash"
	"github.com/rockiecn/check/check"
	comn "github.com/rockiecn/check/common"
)

type Operator struct {
	OpSK         string
	OpAddr       common.Address
	ContractAddr common.Address
	// to -> nonce
	Nonces map[common.Address]uint64
}

type IOperator interface {
	GenCheck(value *big.Int, token common.Address, from common.Address, to common.Address) (*check.Check, error)
	DeployContract(value *big.Int) (*types.Transaction, common.Address, error)

	//TODO:

	// query current balance of contract
	QueryBalance() (*big.Int, error)
	// query contract nonce of a provider
	QueryNonce(to common.Address) (uint64, error)
	// give money to contract
	Deposit(value *big.Int) (*types.Transaction, error)
	// sell a check to a user based on an apply
	Sell(a *Apply) (*Receipt, error)
}

// new operator, a contract is deployed.
// tx's receipt should be checked to make sure contract deploying is completed.
func New(sk string, token string) (IOperator, *types.Transaction, error) {
	op := &Operator{
		OpSK:   sk,
		OpAddr: comn.KeyToAddr(sk),
		Nonces: make(map[common.Address]uint64),
		//Recorder: NewRec(),
	}

	// give 20 eth to new contract
	tx, addr, err := op.DeployContract(comn.String2BigInt("20000000000000000000"))
	if err != nil {
		return nil, nil, err
	}
	op.ContractAddr = addr

	return op, tx, nil
}

// generate a check
func (op *Operator) GenCheck(value *big.Int, token common.Address, from common.Address, to common.Address) (*check.Check, error) {

	// construct check
	chk := &check.Check{
		Value:        value,
		TokenAddr:    token,
		FromAddr:     from,
		ToAddr:       to,
		Nonce:        op.Nonces[to],
		OpAddr:       op.OpAddr,
		ContractAddr: op.ContractAddr,
	}

	// signed by operator
	err := chk.Sign(op.OpSK)
	if err != nil {
		return nil, err
	}

	// store check
	// err = op.Recorder.Record(chk)
	// if err != nil {
	// 	return nil, err
	// }

	// update nonce
	op.Nonces[to] = op.Nonces[to] + 1

	return chk, nil
}

// value: money to new contract
func (op *Operator) DeployContract(value *big.Int) (tx *types.Transaction, contractAddr common.Address, err error) {

	ethClient, err := comn.GetClient(comn.HOST)
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
	auth, err := comn.MakeAuth(op.OpSK, value, bigNonce, gasPrice, 9000000)
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
	ethClient, err := comn.GetClient(comn.HOST)
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
	ethClient, err := comn.GetClient(comn.HOST)
	if err != nil {
		return 0, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	auth := new(bind.CallOpts)
	auth.From = op.OpAddr

	// get contract instance from address
	cashInstance, err := cash.NewCash(op.ContractAddr, ethClient)
	if err != nil {
		return 0, errors.New("newcash failed")
	}

	nonce, err := cashInstance.GetNonce(auth, to)
	if err != nil {
		return 0, errors.New("tx failed")
	}

	return nonce, nil
}

// deposit some money into contract
func (op *Operator) Deposit(value *big.Int) (*types.Transaction, error) {
	ethClient, err := comn.GetClient(comn.HOST)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	auth, err := comn.MakeAuth(op.OpSK, nil, nil, big.NewInt(1000), 9000000)
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

	tx, err := cashInstance.Deposit(auth)
	if err != nil {
		return nil, errors.New("tx failed")
	}

	return tx, nil
}

// apply to buy a check
type Apply struct {
	Value *big.Int       // 购买的支票金额
	Token common.Address // 购买的货币类型
	From  common.Address // 支票的支付方地址
	To    common.Address // 支票的接收方地址
	Date  time.Time      // 购买日期
	Name  string         // 购买人姓名
	Tel   string         // 购买人联系方式
	Sig   string         // 运营商的签名
}

// receipt of a check
type Receipt struct {
	Dt    time.Time      // 购买日期
	Value *big.Int       // 购买金额
	Token common.Address // 货币类型
	Op    common.Address // 运营商地址
	From  common.Address // 付款方地址
	To    common.Address // 收款方地址
	Nonce uint64         // nonce
	Sig   string         // 运营商的签名
}

// sell a check to user, based on the apply, return a receipt
func (op *Operator) Sell(a *Apply) (*Receipt, error) {
	return nil, nil
}

type CheckPool struct {
	// to -> []check
	Pool map[common.Address][]*check.Check
}

// called by operator when a check is generated.
func (p *CheckPool) Store(c *check.Check) error {
	return nil
}

// get a check by receipt, can be called by user with rpc
func (p *CheckPool) Get(r *Receipt) (*check.Check, error) {
	return nil, nil
}
