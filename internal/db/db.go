package db

import (
	"errors"

	"github.com/syndtr/goleveldb/leveldb"
)

// db operation
// key for order: oid
// key for paycheck: provider address + nonce
func WriteDB(db *leveldb.DB, key []byte, buf []byte) error {

	if db == nil {
		return errors.New("nil db")
	}

	err := db.Put(key, buf, nil)
	if err != nil {
		return err
	}
	return nil
}

// db operation
func ReadDB(db *leveldb.DB, key []byte) ([]byte, error) {
	if db == nil {
		return nil, errors.New("nil db")
	}

	buf, err := db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
