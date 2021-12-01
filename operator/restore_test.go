package operator

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/utils"
)

var (
	odr0 = &Order{
		ID:    0,
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

	odr1 = &Order{
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

	odr2 = &Order{
		ID:    2,
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

	chk0 = &check.Check{
		CheckInfo: check.CheckInfo{
			Value:     utils.String2BigInt("100000000000000000000"),
			TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
			Nonce:     0,
			FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
		},
		CheckSig: utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
	}

	chk1 = &check.Check{
		CheckInfo: check.CheckInfo{
			Value:     utils.String2BigInt("100000000000000000000"),
			TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
			Nonce:     1,
			FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
		},
		CheckSig: utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
	}

	chk2 = &check.Check{
		CheckInfo: check.CheckInfo{
			Value:     utils.String2BigInt("100000000000000000000"),
			TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
			Nonce:     2,
			FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
		},
		CheckSig: utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
	}
)

func TestStore(t *testing.T) {
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

	// id->order
	err = Op.StoreOrder(odr0)
	if err != nil {
		t.Fatal(err)
	}
	err = Op.StoreOrder(odr1)
	if err != nil {
		t.Fatal(err)
	}
	err = Op.StoreOrder(odr2)
	if err != nil {
		t.Fatal(err)
	}

	// id->chk
	err = Op.StoreChk(0, chk0)
	if err != nil {
		t.Fatal(err)
	}
	err = Op.StoreChk(1, chk1)
	if err != nil {
		t.Fatal(err)
	}
	err = Op.StoreChk(2, chk2)
	if err != nil {
		t.Fatal(err)
	}
}

/*
func TestRestoreOp(t *testing.T) {
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

	// restore order and check into pool
	err = Op.RestoreChk()
	if err != nil {
		t.Fatal(err)
	}
	err = Op.RestoreOrder()
	if err != nil {
		t.Fatal(err)
	}

	// show data in pool
	Op.ShowChkPool()
	Op.ShowOdrPool()
}
*/
