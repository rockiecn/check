package serial

import (
	"errors"
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/rockiecn/check/internal/order"
	"github.com/rockiecn/check/internal/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

// serialize an order with cbor
func MarshOdr(odr *order.Order) ([]byte, error) {

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
func UnMarshOdr(buf []byte) (*order.Order, error) {
	if buf == nil {
		return nil, errors.New("nil buf")
	}

	odr := new(order.Order)
	err := cbor.Unmarshal(buf, odr)
	if err != nil {
		fmt.Println("error:", err)
	}
	return odr, nil
}

// db operation
func WriteDB(oid uint64, buf []byte) error {
	db, err := leveldb.OpenFile("./order.db", nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return err
	}
	defer db.Close()

	// uint64 to []byte
	k := utils.Uint64ToByte(oid)

	err = db.Put(k, buf, nil)
	if err != nil {
		return err
	}
	return nil
}

// db operation
func ReadDB(oid uint64) ([]byte, error) {
	db, err := leveldb.OpenFile("./order.db", nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return nil, err
	}
	defer db.Close()

	// uint64 to []byte
	k := utils.Uint64ToByte(oid)

	buf, err := db.Get(k, nil)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
