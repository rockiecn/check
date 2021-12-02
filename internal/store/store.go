package store

import (
	"errors"

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

func (st *Store) Close() error {
	err := st.DB.Close()
	if err != nil {
		return err
	}

	return nil
}
