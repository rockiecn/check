package serial

import (
	"errors"
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/odrmgr"
	"github.com/rockiecn/check/internal/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

// serialize an order with cbor
func MarshOdr(odr *odrmgr.Order) ([]byte, error) {

	if odr == nil {
		return nil, errors.New("nil order")
	}

	b, err := cbor.Marshal(*odr)
	if err != nil {
		fmt.Println("error:", err)
	}
	return b, nil
}

// decode a buf into order
func UnMarshOdr(buf []byte) (*odrmgr.Order, error) {
	if buf == nil {
		return nil, errors.New("nil buf")
	}

	odr := new(odrmgr.Order)
	err := cbor.Unmarshal(buf, odr)
	if err != nil {
		fmt.Println("error:", err)
	}
	return odr, nil
}

// serialize an order with cbor
func MarshPchk(pchk *check.Paycheck) ([]byte, error) {

	if pchk == nil {
		return nil, errors.New("nil pchk")
	}

	b, err := cbor.Marshal(*pchk)
	if err != nil {
		fmt.Println("error:", err)
	}
	return b, nil
}

// decode a buf into order
func UnMarshPchk(buf []byte) (*check.Paycheck, error) {
	if buf == nil {
		return nil, errors.New("nil buf")
	}

	pchk := new(check.Paycheck)
	err := cbor.Unmarshal(buf, pchk)
	if err != nil {
		fmt.Println("error:", err)
	}
	return pchk, nil
}

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

// get key from paycheck for db operation
// key = to address + nonce
func GetKey(pchk *check.Paycheck) []byte {
	var key []byte
	key = append(key, pchk.ToAddr.Bytes()...)
	key = append(key, utils.Uint64ToByte(pchk.Nonce)...)
	return key
}
