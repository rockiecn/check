package operator

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/db"
	"github.com/rockiecn/check/internal/odrmgr"
	"github.com/rockiecn/check/internal/utils"
)

// test for serial order
// Process:
// 1.create an order
// 2.marshal the order
// 3.store serialized order into db
// 4.read it from db
// 5.unmarshal it back to a new order
// 6.check if old order identical to new order
func TestSerialOdr(t *testing.T) {
	// generate operator
	opSk, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	op, err := New(opSk)
	if err != nil {
		t.Fatal(err)
	}
	Op, ok := op.(*Operator)
	if !ok {
		t.Fatal("new operator assertion failed")
	}

	// create an order
	odr := &odrmgr.Order{
		ID:    1,
		Token: common.HexToAddress("0xb213d01542d129806d664248a380db8b12059061"),
		Value: utils.String2BigInt("300000000000000000"), // order value: 0.3 eth
		From:  common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
		To:    common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),

		Time:  time.Now().Unix(),
		Name:  "jack",
		Tel:   "123123123",
		Email: "asdf@asdf.com",
		State: 0,
	}
	if odr == nil {
		t.Fatal("create order failed")
	}

	err = Op.StoreOrder(odr)
	if err != nil {
		t.Fatal(err)
	}

	// read from db
	k := utils.Uint64ToByte(1)
	newBuf, err := db.ReadDB(Op.orderDB, k)
	if err != nil {
		t.Fatal(err)
	}

	// unmarshal order
	newOdr := &odrmgr.Order{}
	err = newOdr.UnMarshal(newBuf)
	if err != nil {
		t.Fatal(err)
	}

	eq, err := odr.Equal(newOdr)
	if err != nil {
		t.Fatal(err)
	}
	if !eq {
		t.Fatal("new order not equal order")
	}
}

// test for serial paycheck
// Process:
// 1.create an paycheck
// 2.marshal the paycheck
// 3.store serialized paycheck into db
// 4.read it from db
// 5.unmarshal it back to a new paycheck
// 6.check if old paycheck identical to new paycheck
func TestSerialChk(t *testing.T) {
	// generate operator
	opSk, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	op, err := New(opSk)
	if err != nil {
		t.Fatal(err)
	}
	Op, ok := op.(*Operator)
	if !ok {
		t.Fatal("new operator assertion failed")
	}

	// create a pay check
	// pchk := &check.Paycheck{
	// 	Check: &check.Check{
	// 		Value:     utils.String2BigInt("100000000000000000000"),
	// 		TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
	// 		Nonce:     6,
	// 		FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
	// 		ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
	// 		OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
	// 		CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
	// 		CheckSig:  utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
	// 	},
	// 	PayValue:    utils.String2BigInt("1000000000000000000"),
	// 	PaycheckSig: utils.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	// }

	// create a check
	chk := &check.Check{
		Value:     utils.String2BigInt("100000000000000000000"),
		TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
		Nonce:     6,
		FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
		ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
		OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
		CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
		CheckSig:  utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
	}

	Op.StoreChk(1, chk)

	// read from db
	newBuf, err := db.ReadDB(Op.chkDB, chk.ToKey())
	if err != nil {
		t.Fatal(err)
	}

	// unmarshal chk
	newchk := &check.Check{}
	err = newchk.UnMarshal(newBuf)
	if err != nil {
		t.Fatal(err)
	}

	eq, err := chk.Equal(newchk)
	if err != nil {
		t.Fatal(err)
	}
	if !eq {
		t.Fatal("newchk not equal chk")
	}
}
