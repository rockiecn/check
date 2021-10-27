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
	"github.com/rockiecn/check/pb"
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
	DeployContract(value *big.Int) (*types.Transaction, common.Address, error)
	// query current balance of contract
	QueryBalance() (*big.Int, error)
	// query contract nonce of a provider
	QueryNonce(to common.Address) (uint64, error)
	// give money to contract
	Deposit(value *big.Int) (*types.Transaction, error)

	GenCheck(o *order.Order) (*check.Check, error)
	Mail(o *order.Order) error
	Aggregate(wrap *pb.SerializeData) (*check.BatchCheck, error)
}

// new operator, a contract is deployed.
func New(sk string, token string) (IOperator, *types.Transaction, error) {
	op := &Operator{
		OpSK:   sk,
		OpAddr: comn.KeyToAddr(sk),
		Nonces: make(map[common.Address]uint64),
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

// generate a new check with order
func (op *Operator) GenCheck(o *order.Order) (*check.Check, error) {

	var newNonce uint64

	checks := op.Pool.Data[o.To]
	newNonce = uint64(len(checks))

	chk := &check.Check{
		Value:        o.Value,
		TokenAddr:    o.Token,
		FromAddr:     o.From,
		ToAddr:       o.To,
		Nonce:        newNonce,
		OpAddr:       op.OpAddr,
		ContractAddr: op.ContractAddr,
	}

	// signed by operator
	err := chk.Sign(op.OpSK)
	if err != nil {
		return nil, err
	}

	// update nonce
	op.Nonces[o.To] = newNonce

	return chk, nil
}

func (op *Operator) Mail(o *order.Order) error {
	//先使用订单在支票池调用GetCheck方法找到指定支票，然后将支票发送到订单中的邮箱地址。
	return nil
}

// for serialize
type SerialData struct {
	Data []check.Paycheck
}

// mutli paycheck to a bacthCheck
func (op *Operator) Aggregate(wrap *pb.SerializeData) (*check.BatchCheck, error) {

	// no data
	if len(wrap.Data) == 0 {
		return nil, errors.New("no paycheck in data")
	}

	// initialize
	minNonce := wrap.Data[0].Check.Nonce
	maxNonce := wrap.Data[0].Check.Nonce
	totalPayvalue := new(big.Int)
	toAddr := common.HexToAddress(wrap.Data[0].Check.To)
	for _, v := range wrap.Data {
		// contruct paycheck from pb data
		pc := &check.Paycheck{}

		pc.Check.Value = comn.String2BigInt(v.Check.Value)
		pc.Check.TokenAddr = common.HexToAddress(v.Check.Token)
		pc.Check.Nonce = v.Check.Nonce
		pc.Check.FromAddr = common.HexToAddress(v.Check.From)
		pc.Check.ToAddr = common.HexToAddress(v.Check.To)
		pc.Check.OpAddr = common.HexToAddress(v.Check.Op)
		pc.Check.ContractAddr = common.HexToAddress(v.Check.Contract)
		pc.Check.CheckSig = v.Check.ChkSig
		pc.PayValue = comn.String2BigInt(v.Payvalue)
		pc.PaycheckSig = v.PayCheckSig

		v1, _ := pc.Check.Verify()
		if !v1 {
			return nil, errors.New("check sig verify failed")
		}
		// verify both signs
		v2, _ := pc.Verify()
		if !v2 {
			return nil, errors.New("paycheck sig verify failed")
		}

		// verify payvalue
		if pc.PayValue.Cmp(pc.Check.Value) > 0 {
			return nil, errors.New("payvalue exceed value")
		}

		// to address must be same
		if common.HexToAddress(v.Check.To) != toAddr {
			return nil, errors.New("to address not identical")
		}

		// accumulate payvalue
		totalPayvalue = totalPayvalue.Add(totalPayvalue, pc.PayValue)

		// update minNonce, maxNonce
		if pc.Check.Nonce < minNonce {
			minNonce = v.Check.Nonce
		}
		if pc.Check.Nonce > maxNonce {
			maxNonce = v.Check.Nonce
		}
	}

	// construct batch check
	batch := &check.BatchCheck{}
	batch.OpAddr = op.OpAddr
	batch.ToAddr = toAddr
	batch.BatchValue = totalPayvalue
	batch.MinNonce = minNonce
	batch.MaxNonce = maxNonce
	// sign
	err := batch.Sign(op.OpSK)
	if err != nil {
		return nil, errors.New("batch sign failed")
	}

	return batch, nil
}

type CheckPool struct {
	// to -> []check
	Data map[common.Address][]*check.Check
}

// 如果nonce越界，则先使用nil填充池，直到nonce当前位置，然后把支票放置到nonce指定位置
// 如果nonce没越界，并且此nonce位置已存在支票，报错返回
// 如果nonce没越界，并且此nonce位置不存在支票，则将check直接放到nonce指定的位置
func (p *CheckPool) Store(chk *check.Check) error {
	// get slice
	s := p.Data[chk.ToAddr]

	// if nonce is out of boundary, extend pool and put check into right position
	if chk.Nonce >= uint64(len(s)) {
		// padding nils
		pad := chk.Nonce - uint64(len(s))
		for i := uint64(0); i < pad; i++ {
			s = append(s, nil)
		}
		// right position to append check
		if chk.Nonce == uint64(len(s)) {
			s = append(s, chk)
			p.Data[chk.ToAddr] = s
			return nil
		} else {
			return errors.New("bad check nonce")
		}
	}

	// check in boundary, put check into nonce position
	s[chk.Nonce] = chk

	return nil
}

// get a check according to order
func (p *CheckPool) GetCheck(o *order.Order) (*check.Check, error) {

	// nonce out of boundary
	if o.Nonce >= uint64(len(p.Data[o.To])) {
		return nil, errors.New("nonce exceed check storage boundary")
	}

	if p.Data[o.To][o.Nonce] == nil {
		return nil, errors.New("no check at order's nonce")
	} else {
		return p.Data[o.To][o.Nonce], nil
	}
}
