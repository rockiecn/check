package user

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/db"
	"github.com/rockiecn/check/internal/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

// nonce to check
type Paychecks map[uint64]*check.Paycheck

type User struct {
	UserSK   string
	UserAddr common.Address
	Host     string

	// address to paychecks
	Pool map[common.Address]Paychecks

	db *leveldb.DB

	ClosedbFunc func(*leveldb.DB) error
}

type IUser interface {
	Put(pchk *check.Paycheck) error
	Pay(to common.Address, dataValue *big.Int) (*check.Paycheck, error)
}

// user db file name
var userDBfile = "user.db"

// create an user object out of sk
func New(sk string) (IUser, error) {
	addr, err := utils.SkToAddr(sk)
	if err != nil {
		return nil, err
	}
	user := &User{
		UserSK:   sk,
		UserAddr: addr,
		Pool:     make(map[common.Address]Paychecks),
	}

	// open user db
	user.db, err = leveldb.OpenFile(userDBfile, nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return nil, err
	}

	// declare close func
	user.ClosedbFunc = func(db *leveldb.DB) error {
		err := db.Close()
		if err != nil {
			return err
		}
		return nil
	}

	return user, nil
}

// generate a paycheck out of a check
// check here is acquired from operator's order manager by order id
func (user *User) GenPchk(chk *check.Check) (*check.Paycheck, error) {
	if chk == nil {
		return nil, errors.New("check nil")
	}

	pchk := &check.Paycheck{
		Check:    chk,
		PayValue: big.NewInt(0),
	}

	err := pchk.Sign(user.UserSK)
	if err != nil {
		return nil, errors.New("paycheck sign error")
	}

	return pchk, nil
}

// put a paycheck into user's pool
func (user *User) Put(pchk *check.Paycheck) error {

	// put into pool
	if user.Pool[pchk.ToAddr] == nil {
		user.Pool[pchk.ToAddr] = make(Paychecks)
		user.Pool[pchk.ToAddr][pchk.Nonce] = pchk
	} else {
		if user.Pool[pchk.ToAddr][pchk.Nonce] != nil {
			return errors.New("paycheck with same nonce already exist in pool")
		}
		user.Pool[pchk.ToAddr][pchk.Nonce] = pchk
	}

	return nil
}

// serialize and store a paycheck into db
func (user *User) Store(pchk *check.Paycheck) error {
	// serialize paycheck
	b, err := pchk.Serialize()
	if err != nil {
		return err
	}
	// write db
	err = db.WriteDB(user.db, utils.ToKey(pchk.Check.ToAddr, pchk.Check.Nonce), b)
	if err != nil {
		return err
	}

	return nil
}

// restore all paychecks from db
func (user *User) Restore() error {
	if user.db == nil {
		return errors.New("nil db")
	}

	// read data from db
	iter := user.db.NewIterator(nil, nil)
	for iter.Next() {
		//k := iter.Key()
		v := iter.Value()
		pchk := &check.Paycheck{}
		err := pchk.DeSerialize(v)
		if err != nil {
			return err
		}

		// put pchk into memory
		err = user.Put(pchk)
		if err != nil {
			return err
		}
	}

	iter.Release()
	err := iter.Error()
	if err != nil {
		return err
	}

	return nil
}

// Pay- Create a paycheck legal for paying dataValue
// First find the legal paycheck in pool.
// 1. Remain value is enough for paying dataValue.
// 2. Paycheck's nonce no less than nonce in contract
// 3. The one with the minimum nonce in the result
// Then update it with accumulated payvalue and new signature.
func (user *User) Pay(proAddr common.Address, dataValue *big.Int) (*check.Paycheck, error) {

	var (
		theOne   = (*check.Paycheck)(nil)
		minNonce = ^uint64(0)
	)

	// view each paycheck in user pool
	for _, v := range user.Pool[proAddr] {
		// get nonce in contract
		ctNonce, err := utils.GetCtNonce(v.CtrAddr, v.ToAddr)
		if err != nil {
			return nil, err
		}
		// nonce too old
		if v.Nonce < ctNonce {
			continue
		}

		// remain value must no less than dataValue
		remain := new(big.Int).Sub(v.Value, v.PayValue)
		if remain.Cmp(dataValue) < 0 {
			continue
		} else {
			// got one
			if v.Nonce < minNonce {
				minNonce = v.Check.Nonce
				theOne = v
			} else {
				continue
			}
		}
	}

	// usable paycheck not found
	if theOne == nil {
		return nil, errors.New("user: usable paycheck not found")
	} else {
		// a tempor paycheck for sign
		newPchk := new(check.Paycheck)
		*newPchk = *theOne
		newPchk.PayValue = new(big.Int).Add(theOne.PayValue, dataValue)
		// sign
		err := newPchk.Sign(user.UserSK)
		if err != nil {
			return nil, errors.New("sign payckeck failed")
		}

		// update data in pool with new paycheck
		*theOne = *newPchk

		return theOne, nil
	}
}

// show pool
func (user *User) ShowPool() {
	for k, v := range user.Pool {
		fmt.Println("-> provider:", k)
		for k1, v1 := range v {
			fmt.Println("nonce:", k1)
			fmt.Println("paycheck info:")
			fmt.Println(*v1)
		}
	}
}
