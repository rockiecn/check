package serial

import (
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/order"
	"github.com/rockiecn/check/internal/utils"
)

func TestDB(t *testing.T) {
	odr := order.NewOdr(1,
		common.HexToAddress("0xb213d01542d129806d664248a380db8b12059061"),
		common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
		common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
		utils.String2BigInt("300000000000000000"), // order value: 0.3 eth
		time.Now().Unix(),
		"jack",
		"123123123",
		"asdf@asdf.com",
		0,
	)
	if odr == nil {
		t.Fatal("create order failed")
	}

	fmt.Printf("original order:\n%v\n", odr)

	// marshal order
	buf, err := MarshOdr(odr)
	if err != nil {
		t.Fatal(err)
	}

	// put into db
	err = WriteDB("./order.db", 1, buf)
	if err != nil {
		t.Fatal(err)
	}

	// read from db
	newBuf, err := ReadDB("./order.db", 1)
	if err != nil {
		t.Fatal(err)
	}

	// unmarshal order
	newOdr, err := UnMarshOdr(newBuf)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("unmarshaled order read from db:\n%v\n", newOdr)
}
