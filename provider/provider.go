package provider

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rockiecn/check/internal/cash"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/store"
	"github.com/rockiecn/check/internal/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

type Provider struct {
	ProviderSK   string
	ProviderAddr common.Address
	Host         string

	Pool      map[uint64]*check.Paycheck
	BatchPool map[uint64]*check.BatchCheck

	// nonce -> paycheck
	PcStorer store.Storer // storer for paycheck
	// nonce -> batch
	BtStorer store.Storer // storer for batch
}

type IProvider interface {
	Verify(pchk *check.Paycheck, dataValue *big.Int) (bool, error)
	Put(pchk *check.Paycheck) error
	GetNextPayable() (*check.Paycheck, error)
	Withdraw(pc *check.Paycheck) (tx *types.Transaction, err error)
	QueryBalance() (*big.Int, error)
}

// db file name
var pchkDBfile = "pchk.db"
var batchDBfile = "batch.db"

// create a provider out of sk
func New(sk string) (IProvider, error) {
	addr, err := utils.SkToAddr(sk)
	if err != nil {
		return nil, err
	}
	pro := &Provider{
		ProviderSK:   sk,
		ProviderAddr: addr,
		Host:         "http://localhost:8545",
		Pool:         make(map[uint64]*check.Paycheck),
	}

	pchkDB := &store.Store{}
	pchkDB.DB, err = leveldb.OpenFile(pchkDBfile, nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return nil, err
	}

	batchDB := &store.Store{}
	batchDB.DB, err = leveldb.OpenFile(batchDBfile, nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return nil, err
	}

	// init db
	pro.PcStorer = pchkDB
	pro.BtStorer = batchDB

	return pro, nil
}

// verify paycheck before store paycheck into pool
func (pro *Provider) Verify(pchk *check.Paycheck, dataValue *big.Int) (bool, error) {

	// value should no less than payvalue
	if pchk.Check.Value.Cmp(pchk.PayValue) < 0 {
		return false, errors.New("value less than payvalue")
	}

	// check nonce shuould larger than contract nonce
	contractNonce, err := utils.GetCtNonce(pchk.Check.CtrAddr, pro.ProviderAddr)
	if err != nil {
		return false, err
	}
	if pchk.Check.Nonce < contractNonce {
		return false, errors.New("nonce should not less than contract nonce")
	}

	// to address must be provider
	if pchk.Check.ToAddr != pro.ProviderAddr {
		return false, errors.New("to address must be provider")
	}

	// get paycheck in pool
	old := pro.Pool[pchk.Check.Nonce]
	// verify payvalue
	if old == nil {
		if pchk.PayValue.Cmp(dataValue) == 0 {
			return true, nil
		} else {
			return false, errors.New("payAmount not equal dataValue 1")
		}
	} else {
		payAmount := new(big.Int).Sub(pchk.PayValue, old.PayValue)
		if payAmount.Cmp(dataValue) == 0 {
			return true, nil
		} else {
			return false, errors.New("payAmount not equal dataValue 2")
		}
	}
}

// put a paycheck into pool
func (pro *Provider) Put(pc *check.Paycheck) error {
	if pc == nil {
		return errors.New("paycheck nil")
	}

	// put into pool
	pro.Pool[pc.Check.Nonce] = pc

	// write db
	err := pro.Store(pc)
	if err != nil {
		return err
	}

	return nil
}

// put a batch paycheck into pool
func (pro *Provider) PutBatch(bchk *check.BatchCheck) error {
	if bchk == nil {
		return errors.New("batch paycheck nil")
	}

	pro.BatchPool[bchk.MinNonce] = bchk

	return nil
}

// get the next payable paycheck in db
func (pro *Provider) GetNextPayable() (*check.Paycheck, error) {

	// get db
	db := pro.PcStorer.(*store.Store).DB

	if db == nil {
		return nil, errors.New("nil check db")
	}

	var (
		theOne   = (*check.Paycheck)(nil)
		max      = ^uint64(0)
		minNonce = max
	)

	// read data from db
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		// get nonce
		n := utils.ByteToUint64(k)

		// Deserialize paycheck
		pchk := &check.Paycheck{}
		err := pchk.DeSerialize(v)
		if err != nil {
			return nil, err
		}

		// get current nonce in contract
		ctrNonce, err := utils.GetCtNonce(pchk.CtrAddr, pro.ProviderAddr)
		if err != nil {
			return nil, err
		}

		// nonce too old, skip it
		if n < ctrNonce {
			continue
		}

		// get and update minNonce
		if n < minNonce {
			minNonce = n
			theOne = pchk
		}
	}

	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, err
	}

	// not found
	if theOne == nil {
		return nil, errors.New("payable paycheck not found in db")
	}

	// put paycheck into pool
	pro.Put(theOne)

	return theOne, nil
}

// get next payable batch check
func (pro *Provider) GetNextPayableBatch() (*check.BatchCheck, error) {

	st := pro.BtStorer.(*store.Store)

	if st.DB == nil {
		return nil, errors.New("nil check db")
	}

	var (
		theOne   = (*check.BatchCheck)(nil)
		max      = ^uint64(0)
		minNonce = max
	)

	// read data from db
	iter := st.DB.NewIterator(nil, nil)
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		// get nonce
		n := utils.ByteToUint64(k)

		// Deserialize batchcheck
		bchk := &check.BatchCheck{}
		err := bchk.DeSerialize(v)
		if err != nil {
			return nil, err
		}

		// get current nonce in contract
		ctrNonce, err := utils.GetCtNonce(bchk.CtrAddr, pro.ProviderAddr)
		if err != nil {
			return nil, err
		}

		// nonce too old, skip it
		if n < ctrNonce {
			continue
		}

		// get and update minNonce
		if n < minNonce {
			minNonce = n
			theOne = bchk
		}
	}

	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, err
	}

	// not found
	if theOne == nil {
		return nil, errors.New("payable batch check not found in db")
	}

	// put paycheck into pool
	pro.PutBatch(theOne)

	return theOne, nil
}

// CallApplyCheque - send tx to contract to call apply cheque method.
func (pro *Provider) Withdraw(pc *check.Paycheck) (tx *types.Transaction, err error) {

	ethClient, err := utils.GetClient(pro.Host)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	auth, err := utils.MakeAuth(pro.ProviderSK, nil, nil, big.NewInt(1000), 9000000)
	if err != nil {
		return nil, errors.New("make auth failed")
	}

	// get contract instance from address
	cashInstance, err := cash.NewCash(pc.Check.CtrAddr, ethClient)
	if err != nil {
		return nil, errors.New("newcash failed")
	}

	// type convertion, from pc to cashpc for contract
	cashpc := cash.Paycheck{
		Check: cash.Check{
			Value:     pc.Check.Value,
			TokenAddr: pc.Check.TokenAddr,
			Nonce:     pc.Check.Nonce,
			FromAddr:  pc.Check.FromAddr,
			ToAddr:    pc.Check.ToAddr,
			OpAddr:    pc.Check.OpAddr,
			CtrAddr:   pc.Check.CtrAddr,
			CheckSig:  pc.Check.CheckSig,
		},
		PayValue:    pc.PayValue,
		PaycheckSig: pc.PaycheckSig,
	}
	tx, err = cashInstance.Withdraw(auth, cashpc)
	if err != nil {
		return nil, errors.New("tx failed")
	}

	//fmt.Println("Mine a block to complete.")

	return tx, nil
}

// CallApplyCheque - send tx to contract to call apply batch method.
func (pro *Provider) WithdrawBatch(bc *check.BatchCheck) (tx *types.Transaction, err error) {

	// connect
	ethClient, err := utils.GetClient(pro.Host)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	// auth
	auth, err := utils.MakeAuth(pro.ProviderSK, nil, nil, big.NewInt(1000), 9000000)
	if err != nil {
		return nil, errors.New("make auth failed")
	}

	// get contract instance from address
	cashInstance, err := cash.NewCash(bc.CtrAddr, ethClient)
	if err != nil {
		return nil, errors.New("newcash failed")
	}

	// type convertion, from pc to cashbc for contract
	cashbc := cash.BatchCheck{
		OpAddr:     bc.OpAddr,
		ToAddr:     bc.ToAddr,
		CtrAddr:    bc.CtrAddr,
		TokenAddr:  bc.TokenAddr,
		BatchValue: bc.BatchValue,
		MinNonce:   bc.MinNonce,
		MaxNonce:   bc.MaxNonce,
		BatchSig:   bc.BatchSig,
	}

	tx, err = cashInstance.WithdrawBatch(auth, cashbc)
	if err != nil {
		return nil, errors.New("tx failed")
	}

	//fmt.Println("Mine a block to complete.")

	return tx, nil
}

// query provider balance in contract
func (pro *Provider) QueryBalance() (*big.Int, error) {
	ethClient, err := utils.GetClient(utils.HOST)
	if err != nil {
		return nil, errors.New("failed to dial geth")
	}
	defer ethClient.Close()

	balance, err := ethClient.BalanceAt(context.Background(), pro.ProviderAddr, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

// serialize and store a paycheck to db
func (pro *Provider) Store(pchk *check.Paycheck) error {
	// serialize paycheck
	b, err := pchk.Serialize()
	if err != nil {
		return err
	}
	// write db
	err = pro.PcStorer.Put(utils.Uint64ToByte(pchk.Nonce), b)
	if err != nil {
		return err
	}

	return nil
}

// restore a paycheck from db
// key = to + nonce
func (pro *Provider) Restore(to common.Address, n uint64) error {

	k := utils.ToKey(to, n)
	v, err := pro.PcStorer.Get(k)
	if err != nil {
		return err
	}
	// deserialize paycheck
	pc := &check.Paycheck{}
	err = pc.DeSerialize(v)
	if err != nil {
		return err
	}
	// put into pool
	err = pro.Put(pc)
	if err != nil {
		return err
	}

	return nil
}

// store a batch check into db
// key = minNonce
func (pro *Provider) StoreBatch(bc *check.BatchCheck) error {
	// serialize
	b, err := bc.Serialize()
	if err != nil {
		return err
	}
	// write db
	err = pro.BtStorer.Put(utils.Uint64ToByte(bc.MinNonce), b)
	if err != nil {
		return err
	}

	return nil
}

/*
// show pool
func (pro *Provider) ShowPool() {
	for k, v := range pro.Pool {
		fmt.Println("nonce:", k)
		fmt.Println("paycheck info:")
		fmt.Println(*v)
	}
}
*/
