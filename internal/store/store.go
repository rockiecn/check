package store

import (
	"errors"

	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

type Store struct {
	DB *leveldb.DB
}

type Storer interface {
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
	Delete(key []byte) error

	// get next payable paycheck/batchcheck from db
	GetNextPayable() (*check.Paycheck, error)
	GetNextPayableBatch() (*check.BatchCheck, error)

	Clear() error
	Close() error
}

// db operation
// key for order: oid
// key for paycheck: provider address + nonce
func (st *Store) Put(key []byte, buf []byte) error {

	if st.DB == nil {
		return errors.New("nil db")
	}

	err := st.DB.Put(key, buf, nil)
	if err != nil {
		return err
	}
	return nil
}

// db operation
func (st *Store) Get(key []byte) ([]byte, error) {
	if st.DB == nil {
		return nil, errors.New("nil db")
	}

	buf, err := st.DB.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// check if key exists
func (st *Store) Has(key []byte) (bool, error) {
	ok, err := st.DB.Has(key, nil)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errors.New("key not exists")
	}

	return true, nil
}

// delete a key
func (st *Store) Delete(key []byte) error {
	err := st.DB.Delete(key, nil)
	if err != nil {
		return err
	}

	return nil
}

// get the next payable paycheck in db
func (st *Store) GetNextPayable() (*check.Paycheck, error) {

	// get db
	db := st.DB

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
		//ctrNonce, err := utils.GetCtNonce(pchk.CtrAddr, pro.ProviderAddr)c
		ctrNonce, err := utils.GetCtNonce(pchk.CtrAddr, pchk.ToAddr)
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

	return theOne, nil
}

// get next payable batch check
func (st *Store) GetNextPayableBatch() (*check.BatchCheck, error) {

	if st.DB == nil {
		return nil, errors.New("nil batch db")
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
		ctrNonce, err := utils.GetCtNonce(bchk.CtrAddr, bchk.ToAddr)
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

	return theOne, nil
}

// delete all data in db
func (st *Store) Clear() error {

	// read data from db
	iter := st.DB.NewIterator(nil, nil)
	for iter.Next() {
		k := iter.Key()
		err := st.DB.Delete(k, nil)
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

func (st *Store) Close() error {
	err := st.DB.Close()
	if err != nil {
		return err
	}

	return nil
}
