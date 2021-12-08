package test

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/common"
	"github.com/rockiecn/check/internal/utils"
	"github.com/rockiecn/check/operator"
)

// Process:
// 1.init a operator with a contract deployed for it
// 2.init a user
// 3.init a provider
// 4.create 3 orders each has value 0.3 eth
// 5.generate 3 paychecks from each order with payvalue 0.1 eth
// 6.create a batch check with the 3 paychecks
// 7.call withdraw batch to get 0.3 eth
// 8.check the balance of provider, extra 0.3 eth for provider expected
// 9.check the nonce in contract is now maxNonce + 1
func TestBatch(t *testing.T) {

	fmt.Println("-> Init Operator")
	op, err := common.InitOperator("./op/order.db", "./op/check.db")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Init User")
	usr, err := common.InitUser("./usr/pc.db")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Init Provider")
	pro, err := common.InitPro("./pro/pc.db", "./pro/bt.db")
	if err != nil {
		t.Fatal(err)
	}

	op.ChkStorer.Clear()
	op.OdrStorer.Clear()
	usr.PcStorer.Clear()
	pro.BtStorer.Clear()
	pro.PcStorer.Clear()

	// create  order for each provider
	odr0 := &operator.Order{
		ID:    0,
		Token: common.Token,
		Value: utils.String2BigInt("3000000000000000"), // order value: 0.003 eth
		From:  usr.UserAddr,
		To:    pro.ProviderAddr,
		Time:  time.Now().Unix(),
		Name:  "jack",
		Tel:   "123123123",
		Email: "asdf@asdf.com",
		State: 0,
	}
	if odr0 == nil {
		t.Fatal("create order 0 failed")
	}

	// operator store order into pool
	err = op.PutOrder(odr0)
	if err != nil {
		t.Fatal(err)
	}

	// create  order for each provider
	odr1 := &operator.Order{
		ID:    1,
		Token: common.Token,
		Value: utils.String2BigInt("3000000000000000"), // order value: 0.003 eth
		From:  usr.UserAddr,
		To:    pro.ProviderAddr,
		Time:  time.Now().Unix(),
		Name:  "jack",
		Tel:   "123123123",
		Email: "asdf@asdf.com",
		State: 0,
	}
	if odr1 == nil {
		t.Fatal("create order 1 failed")
	}

	// operator store order into pool
	err = op.PutOrder(odr1)
	if err != nil {
		t.Fatal(err)
	}

	// create  order for each provider
	odr2 := &operator.Order{
		ID:    2,
		Token: common.Token,
		Value: utils.String2BigInt("3000000000000000"), // order value: 0.003 eth
		From:  usr.UserAddr,
		To:    pro.ProviderAddr,
		Time:  time.Now().Unix(),
		Name:  "jack",
		Tel:   "123123123",
		Email: "asdf@asdf.com",
		State: 0,
	}
	if odr2 == nil {
		t.Fatal("create order 2 failed")
	}

	// operator store order into pool
	err = op.PutOrder(odr2)
	if err != nil {
		t.Fatal(err)
	}

	// operator create checks from orders
	chk0, err := op.CreateCheck(0)
	if err != nil {
		t.Fatal(err)
	}
	chk1, err := op.CreateCheck(1)
	if err != nil {
		t.Fatal(err)
	}
	chk2, err := op.CreateCheck(2)
	if err != nil {
		t.Fatal(err)
	}

	pchk0 := &check.Paycheck{
		Check:    chk0,
		PayValue: utils.String2BigInt("1000000000000000"), // 0.001 eth
	}
	pchk0.Sign(usr.UserSK) // user sk

	pchk1 := &check.Paycheck{
		Check:    chk1,
		PayValue: utils.String2BigInt("1000000000000000"), // 0.001 eth
	}
	pchk1.Sign(usr.UserSK) // user sk

	pchk2 := &check.Paycheck{
		Check:    chk2,
		PayValue: utils.String2BigInt("1000000000000000"), // 0.001 eth
	}
	pchk2.Sign(usr.UserSK) // user sk

	// aggregate
	pchks := []*check.Paycheck{pchk0, pchk1, pchk2}
	bc, err := op.Aggregate(pchks)
	if err != nil {
		t.Fatal(err)
	}

	b1, err := pro.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("balance before withdraw:", b1)

	// withdraw batch
	fmt.Println("-> withdraw batch check")
	tx, err := pro.WithdrawBatch(bc)
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	gasUsed, err := utils.GetGasUsed(tx)
	if err != nil {
		t.Fatal(err)
	}

	b2, err := pro.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("balance after withdraw:", b2)

	// need add used gas for withdraw tx
	delta := new(big.Int).Sub(b2, b1)
	delta.Add(delta, gasUsed)

	// 0.003 eth
	want := utils.String2BigInt("3000000000000000")

	// check result
	if delta.Cmp(want) != 0 {
		t.Fatal("withdrawed money not equal total payvalue")
	} else {
		fmt.Println("OK- withdrawed money equal total payvalue")
	}

	// check if maxNonce correct
	if bc.MaxNonce != 2 {
		t.Fatal("max nonce in batch check error:", bc.MaxNonce)
	} else {
		fmt.Println("OK- max nonce is right:", bc.MaxNonce)
	}

	// check contract nonce with max nonce
	ctrNonce, err := op.GetNonce(pro.ProviderAddr)
	if err != nil {
		t.Fatal(err)
	}
	if bc.MaxNonce+1 != ctrNonce {
		t.Fatal("nonce in contract must be maxNonce + 1")
	} else {
		fmt.Println("OK- nonce in contract is equal to maxNonce + 1")
	}

}

/*
var globalOp *Operator

// everything ok
func TestAggregateOK(t *testing.T) {
	input := []*check.Paycheck{}

	data0 := &check.Paycheck{
		Check: &check.Check{
			Value:        utils.String2BigInt("100000000000000000000"),
			TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
			Nonce:        6,
			FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			CtrAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			CheckSig:     utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		PayValue:    utils.String2BigInt("1000000000000000000"),
		PaycheckSig: utils.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
	input = append(input, data0)

	data1 := &check.Paycheck{
		Check: &check.Check{
			Value:        utils.String2BigInt("100000000000000000000"),
			TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
			Nonce:        7,
			FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			CtrAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			CheckSig:     utils.String2Byte("08e76a9bce17997ddfec0926c89b6473798dae9ac047f5214082c094d1ae2939238206d5236c321cd4a8fab42133db38ba54d342a9ffb76b48cf0467fceebbdf01"),
		},
		PayValue:    utils.String2BigInt("2000000000000000000"),
		PaycheckSig: utils.String2Byte("9727d4415f25a4f59badded914e14fc074cc775df15faf8e711981d2b8b97702210f8ebf8ee4956de5ad776a5312026b68ac32f769f0d500846697104e53b72000"),
	}
	input = append(input, data1)

	data2 := &check.Paycheck{
		Check: &check.Check{
			Value:        utils.String2BigInt("100000000000000000000"),
			TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
			Nonce:        8,
			FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			CtrAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			CheckSig:     utils.String2Byte("584cca0e6eed3558bd07e1ab40206ecc83dc005ccad16ea9d97586726ec43aeb486a6599e1c77b345ce73c4f7f4c26e78230b752a3f3e42de62c9da261f5923e00"),
		},
		PayValue:    utils.String2BigInt("3000000000000000000"),
		PaycheckSig: utils.String2Byte("c75f9b4f960be6bb48719a01da40370afd00e905a554d54e6adc1fe41ab09ccc149701daf4f1118316997f4b9b748b7ff690c1e9dd653f3e63cb0a3fffa625b300"),
	}
	input = append(input, data2)

	// construct operator
	op, _ := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	batch, err := op.Aggregate(input)
	if err != nil {
		t.Error(err)
	}

	// batch payvalue
	if batch.BatchValue.Cmp(utils.String2BigInt("6000000000000000000")) != 0 {
		t.Error(errors.New("batch payvalue error"))
	}

	// verify minNonce
	if batch.MinNonce != 6 {
		t.Error(errors.New("minNonce error"))
	}
	// verify maxNonce
	if batch.MaxNonce != 8 {
		t.Error(errors.New("maxNonce error"))
	}

	// verify batch signature
	ok, _ := batch.Verify()
	if !ok {
		t.Error("batch verify failed")
	}
}

// test input no paycheck data
func TestAggregateNoData(t *testing.T) {
	input := []*check.Paycheck{}
	op, _ := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	_, got := op.Aggregate(input)
	if got == nil || got.Error() != "no paycheck in data" {
		t.Error("case 'no paycheck data' not detected")
	}
}

// test check verify
func TestAggregateCheckVerify(t *testing.T) {
	input := []*check.Paycheck{}

	data0 := &check.Paycheck{
		Check: &check.Check{
			Value:        utils.String2BigInt("100000000000000000000"),
			TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
			Nonce:        6,
			FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			CtrAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			// changed for test
			CheckSig: utils.String2Byte("444f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		PayValue:    utils.String2BigInt("1000000000000000000"),
		PaycheckSig: utils.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
	input = append(input, data0)

	// construct operator
	op, _ := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	_, got := op.Aggregate(input)

	if got == nil || got.Error() != "check sig verify failed" {
		t.Error("case 'check sig verify failed' not detected")
	}
}

// test paycheck verify
func TestAggregatePayCheckVerify(t *testing.T) {
	input := []*check.Paycheck{}

	data0 := &check.Paycheck{
		Check: &check.Check{
			Value:        utils.String2BigInt("100000000000000000000"),
			TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
			Nonce:        6,
			FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			CtrAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			// changed for test
			CheckSig: utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		PayValue: utils.String2BigInt("1000000000000000000"),
		// changed for test
		PaycheckSig: utils.String2Byte("117d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
	input = append(input, data0)

	// construct operator
	op, _ := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	_, got := op.Aggregate(input)
	if got == nil || got.Error() != "paycheck sig verify failed" {
		t.Error("case 'paycheck sig verify failed' not detected")
	}
}

// test payvalue verify
func TestAggregatePayValue(t *testing.T) {
	input := []*check.Paycheck{}

	data0 := &check.Paycheck{
		Check: &check.Check{
			Value:        utils.String2BigInt("100000000000000000000"),
			TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
			Nonce:        6,
			FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			CtrAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			CheckSig:     utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		PayValue:    utils.String2BigInt("200000000000000000000"), // larger than value
		PaycheckSig: utils.String2Byte("b5b60bc8b85fde998cfe6cf821447e770170780a5e2f7d7b588254284f0e0e3d2040102ddc6abc1b98af5417fa196f445aecf6b85154f45be0b1c3b05bb9cf2800"),
	}
	input = append(input, data0)

	// construct operator
	op, _ := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	_, got := op.Aggregate(input)
	if got == nil || got.Error() != "payvalue exceed value" {
		t.Error("case 'payvalue exceed value' not detected")
	}
}

func TestAggregateToAddressIdentical(t *testing.T) {
	input := []*check.Paycheck{}

	data0 := &check.Paycheck{
		Check: &check.Check{
			Value:        utils.String2BigInt("100000000000000000000"),
			TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
			Nonce:        6,
			FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			CtrAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			CheckSig:     utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		PayValue:    utils.String2BigInt("1000000000000000000"),
		PaycheckSig: utils.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
	input = append(input, data0)

	data1 := &check.Paycheck{
		Check: &check.Check{
			Value:        utils.String2BigInt("100000000000000000000"),
			TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
			Nonce:        7,
			FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:       common.HexToAddress("3320993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			CtrAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			CheckSig:     utils.String2Byte("17d8a014f938995220e861feb51befc2a8bfb8430d91e26ff152b35e2027385b5745c0a89deb5212b811afc9bf4887c53b15f76dade32e48d1e361f682fb208000"),
		},
		PayValue:    utils.String2BigInt("2000000000000000000"),
		PaycheckSig: utils.String2Byte("626d51362677e8757c3dd6b2b1821c80c18cb581073cced1159bca336fd2cb2d05ea51060ab9ad1184bb7c75bfd8ed22bddbc8b2571f3fc7d8b1bd001282299200"),
	}
	input = append(input, data1)

	// construct operator
	op, _ := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	_, got := op.Aggregate(input)
	if got == nil || got.Error() != "to address not identical" {
		t.Error("case 'to address not identical' not detected")
	}
}

func TestContract(t *testing.T) {
	Deploy(t)
	fmt.Println("balance before deposit")
	QueryBalance(t)
	fmt.Println("begin deposit 1 eth to contract.")
	Deposit(t)
	fmt.Println("balance after deposit")
	QueryBalance(t)
	fmt.Println("nonce before setNonce:")
	GetNonce(t)
	fmt.Println("begin set nonce")
	SetNonce(t)
	fmt.Println("nonce after setNonce:")
	GetNonce(t)
}

func Deploy(t *testing.T) {

	op, err := NewOperator(
		"503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb",
		"b213d01542d129806d664248a380db8b12059061")
	if err != nil {
		t.Error("create operator failed:", err)
	}

	globalOp = op.(*Operator)

	// 1 eth
	txHash, addr, err := op.Deploy(utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Error("deploy contract failed:", err)
	}

	// set contract address into operator
	op.SetCtrAddr(addr)

	fmt.Println("deploying contract, address:", addr.String())

	op.WaitForMiner(txHash)

	fmt.Println("contract deploy complete, contract address:", addr.String())
}

func QueryBalance(t *testing.T) {
	b, _ := globalOp.QueryBalance()
	fmt.Println("balance of contract:", b.String())
}

func Deposit(t *testing.T) {

	// deposit 1 eth into contract
	txHash, err := globalOp.Deposit(utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Error(err)
	}

	globalOp.WaitForMiner(txHash)

}

func GetNonce(t *testing.T) {

	to := common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db")
	nonce, err := globalOp.GetNonce(to)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("nonce:", nonce)

}

func SetNonce(t *testing.T) {

	to := common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db")
	txHash, err := globalOp.SetNonce(to, 3)
	if err != nil {
		t.Error(err)
	}

	globalOp.WaitForMiner(txHash)
}


*/
