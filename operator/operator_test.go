package operator

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/utils"
)

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
			ContractAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
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
			ContractAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
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
			ContractAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			CheckSig:     utils.String2Byte("584cca0e6eed3558bd07e1ab40206ecc83dc005ccad16ea9d97586726ec43aeb486a6599e1c77b345ce73c4f7f4c26e78230b752a3f3e42de62c9da261f5923e00"),
		},
		PayValue:    utils.String2BigInt("3000000000000000000"),
		PaycheckSig: utils.String2Byte("c75f9b4f960be6bb48719a01da40370afd00e905a554d54e6adc1fe41ab09ccc149701daf4f1118316997f4b9b748b7ff690c1e9dd653f3e63cb0a3fffa625b300"),
	}
	input = append(input, data2)

	// construct operator
	op, _, _ := NewOperator("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
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
	op, _, _ := NewOperator("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
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
			ContractAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			// changed for test
			CheckSig: utils.String2Byte("444f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		PayValue:    utils.String2BigInt("1000000000000000000"),
		PaycheckSig: utils.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
	input = append(input, data0)

	// construct operator
	op, _, _ := NewOperator("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
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
			ContractAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			// changed for test
			CheckSig: utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		PayValue: utils.String2BigInt("1000000000000000000"),
		// changed for test
		PaycheckSig: utils.String2Byte("117d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
	input = append(input, data0)

	// construct operator
	op, _, _ := NewOperator("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
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
			ContractAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			CheckSig:     utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		PayValue:    utils.String2BigInt("200000000000000000000"), // larger than value
		PaycheckSig: utils.String2Byte("b5b60bc8b85fde998cfe6cf821447e770170780a5e2f7d7b588254284f0e0e3d2040102ddc6abc1b98af5417fa196f445aecf6b85154f45be0b1c3b05bb9cf2800"),
	}
	input = append(input, data0)

	// construct operator
	op, _, _ := NewOperator("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
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
			ContractAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
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
			ContractAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			CheckSig:     utils.String2Byte("17d8a014f938995220e861feb51befc2a8bfb8430d91e26ff152b35e2027385b5745c0a89deb5212b811afc9bf4887c53b15f76dade32e48d1e361f682fb208000"),
		},
		PayValue:    utils.String2BigInt("2000000000000000000"),
		PaycheckSig: utils.String2Byte("626d51362677e8757c3dd6b2b1821c80c18cb581073cced1159bca336fd2cb2d05ea51060ab9ad1184bb7c75bfd8ed22bddbc8b2571f3fc7d8b1bd001282299200"),
	}
	input = append(input, data1)

	// construct operator
	op, _, _ := NewOperator("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	_, got := op.Aggregate(input)
	if got == nil || got.Error() != "to address not identical" {
		t.Error("case 'to address not identical' not detected")
	}
}

func TestDeploy(t *testing.T) {
	op, _, err := NewOperator(
		"503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb",
		"b213d01542d129806d664248a380db8b12059061")
	if err != nil {
		fmt.Println("new operator failed:", err)
	}

	// 1 eth
	txHash, _, err := op.DeployContract(utils.String2BigInt("1000000000000000000"))

	if err != nil {
		t.Errorf("deploy contract failed")
	}

	// connect to geth
	ethClient, err := utils.GetClient(utils.HOST)
	if err != nil {
		fmt.Println("get client failed")
		return
	}
	defer ethClient.Close()

	// wait contract deployed, txReceipt is checked
	for {
		txReceipt, _ := ethClient.TransactionReceipt(context.Background(), txHash.Hash())
		// receipt ok
		if txReceipt != nil {
			break
		}
		fmt.Println("waiting for miner, 5 seconds..")
		time.Sleep(time.Duration(5) * time.Second)
	}

	fmt.Println("contract deploy complete")
}
