package db

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

// db operation
// key for order: oid
// key for paycheck: provider address + nonce
func WriteDB(dbfile string, key []byte, buf []byte) error {
	db, err := leveldb.OpenFile(dbfile, nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return err
	}
	defer db.Close()

	err = db.Put(key, buf, nil)
	if err != nil {
		return err
	}
	return nil
}

// db operation
func ReadDB(dbfile string, key []byte) ([]byte, error) {
	db, err := leveldb.OpenFile(dbfile, nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return nil, err
	}
	defer db.Close()

	buf, err := db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
