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
	comn "github.com/rockiecn/check/common"
	"github.com/rockiecn/check/order"
)

type Operator struct {
	OpSK         string
	OpAddr       common.Address
	ContractAddr common.Address
	// to -> nonce
	Nonces map[common.Address]uint64

	Pool CheckPool
}

type IOperator interface {
	GenCheck(o *order.Order) (*check.Check, error)
	DeployContract(value *big.Int) (*types.Transaction, common.Address, error)

	// query current balance of contract
	QueryBalance() (*big.Int, error)
	// query contract nonce of a provider
	QueryNonce(to common.Address) (uint64, error)
	// give money to contract
	Deposit(value *big.Int) (*types.Transaction, error)
}

// new operator, a contract is deployed.
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

// value: money to new contract
func (op *Operator) DeployContract(value *big.Int) (tx *types.Transaction, contractAddr common.Address, err error) {

	// connect to node
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
	nonce, err := comn.QueryNonce(op.OpAddr, op.ContractAddr, to)
	return nonce, err
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

	// call contract
	tx, err := cashInstance.Deposit(auth)
	if err != nil {
		return nil, errors.New("tx failed")
	}

	return tx, nil
}

// generate a check with apply
func (op *Operator) GenCheck(o *order.Order) (*check.Check, error) {
	chk := &check.Check{
		Value:        o.Value,
		TokenAddr:    o.Token,
		FromAddr:     o.From,
		ToAddr:       o.To,
		Nonce:        op.Nonces[o.To] + 1,
		OpAddr:       op.OpAddr,
		ContractAddr: op.ContractAddr,
	}

	// signed by operator
	err := chk.Sign(op.OpSK)
	if err != nil {
		return nil, err
	}

	// update nonce to latest nonce
	op.Nonces[o.To] = chk.Nonce

	return chk, nil
}

func (op *Operator) SendCheck(o *order.Order) error {
	//先使用订单在支票池调用GetCheck方法找到指定支票，然后将支票发送到订单中的邮箱地址。
	return nil
}

func (op *Operator) Aggregate(data []byte) (batch *check.BatchCheck, sigBatch []byte, err error) {
	/*
		先将序列化的数据反序列化成paycheck数组。

		然后验证每一张paycheck的签名（operator和user），以及paycheck的payvalue值是否不大于value值。

		找到这批paycheck的minNonce和maxNonce并计算出总金额。

		然后使用节点地址，支票累计总金额，minNonce，maxNonce生成聚合支票，并对聚合支票生成签名sig。

		返回聚合支票batch和签名sigBatch。
	*/
	return nil, nil, nil
}

type CheckPool struct {
	// to -> []check
	Data map[common.Address][]*check.Check
}

// called when a new check is generated.
func (p *CheckPool) Store(c *check.Check) error {
	s := p.Data[c.ToAddr]

	// new nonce must be max
	if len(s) > 0 && c.Nonce <= s[len(s)-1].Nonce {
		return errors.New("nonce not max")
	}

	// ok to append
	p.Data[c.ToAddr] = append(p.Data[c.ToAddr], c)

	return nil
}

// get a check according to order
func (p *CheckPool) GetCheck(o *order.Order) (*check.Check, error) {

	return nil, nil
}
