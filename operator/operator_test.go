package operator

import (
	"errors"
	"testing"

	comn "github.com/rockiecn/check/common"
	"github.com/rockiecn/check/pb"
)

func TestAggregateOK(t *testing.T) {
	input := &pb.SerializeData{}

	data0 := &pb.PayCheck{
		Check: &pb.Check{
			Value:    "100000000000000000000",
			Token:    "b213d01542d129806d664248a380db8b12059061",
			Nonce:    6,
			From:     "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
			To:       "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
			Op:       "5B38Da6a701c568545dCfcB03FcB875f56beddC4",
			Contract: "1c91347f2A44538ce62453BEBd9Aa907C662b4bD",
			ChkSig:   comn.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		Payvalue:    "1000000000000000000",
		PayCheckSig: comn.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
	input.Data = append(input.Data, data0)

	data1 := &pb.PayCheck{
		Check: &pb.Check{
			Value:    "100000000000000000000",
			Token:    "b213d01542d129806d664248a380db8b12059061",
			Nonce:    7,
			From:     "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
			To:       "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
			Op:       "5B38Da6a701c568545dCfcB03FcB875f56beddC4",
			Contract: "1c91347f2A44538ce62453BEBd9Aa907C662b4bD",
			ChkSig:   comn.String2Byte("08e76a9bce17997ddfec0926c89b6473798dae9ac047f5214082c094d1ae2939238206d5236c321cd4a8fab42133db38ba54d342a9ffb76b48cf0467fceebbdf01"),
		},
		Payvalue:    "2000000000000000000",
		PayCheckSig: comn.String2Byte("9727d4415f25a4f59badded914e14fc074cc775df15faf8e711981d2b8b97702210f8ebf8ee4956de5ad776a5312026b68ac32f769f0d500846697104e53b72000"),
	}
	input.Data = append(input.Data, data1)

	data2 := &pb.PayCheck{
		Check: &pb.Check{
			Value:    "100000000000000000000",
			Token:    "b213d01542d129806d664248a380db8b12059061",
			Nonce:    8,
			From:     "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
			To:       "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
			Op:       "5B38Da6a701c568545dCfcB03FcB875f56beddC4",
			Contract: "1c91347f2A44538ce62453BEBd9Aa907C662b4bD",
			ChkSig:   comn.String2Byte("584cca0e6eed3558bd07e1ab40206ecc83dc005ccad16ea9d97586726ec43aeb486a6599e1c77b345ce73c4f7f4c26e78230b752a3f3e42de62c9da261f5923e00"),
		},
		Payvalue:    "3000000000000000000",
		PayCheckSig: comn.String2Byte("c75f9b4f960be6bb48719a01da40370afd00e905a554d54e6adc1fe41ab09ccc149701daf4f1118316997f4b9b748b7ff690c1e9dd653f3e63cb0a3fffa625b300"),
	}
	input.Data = append(input.Data, data2)

	// construct operator
	op, _, _ := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	batch, err := op.Aggregate(input)
	if err != nil {
		t.Error(err)
	}

	// batch payvalue
	if batch.BatchValue.Cmp(comn.String2BigInt("6000000000000000000")) != 0 {
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
	input := &pb.SerializeData{}
	op, _, _ := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	_, got := op.Aggregate(input)
	if got == nil || got.Error() != "no paycheck in data" {
		t.Error("case 'no paycheck data' not detected")
	}
}

// test check verify
func TestAggregateCheckVerify(t *testing.T) {
	input := &pb.SerializeData{}
	data0 := &pb.PayCheck{
		Check: &pb.Check{
			Value:    "100000000000000000000",
			Token:    "b213d01542d129806d664248a380db8b12059061",
			Nonce:    6,
			From:     "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
			To:       "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
			Op:       "5B38Da6a701c568545dCfcB03FcB875f56beddC4",
			Contract: "1c91347f2A44538ce62453BEBd9Aa907C662b4bD",
			// changed for test
			ChkSig: comn.String2Byte("444f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		Payvalue:    "1000000000000000000",
		PayCheckSig: comn.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
	input.Data = append(input.Data, data0)

	// construct operator
	op, _, _ := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	_, got := op.Aggregate(input)

	if got == nil || got.Error() != "check sig verify failed" {
		t.Error("case 'check sig verify failed' not detected")
	}
}

// test paycheck verify
func TestAggregatePayCheckVerify(t *testing.T) {
	input := &pb.SerializeData{}
	data0 := &pb.PayCheck{
		Check: &pb.Check{
			Value:    "100000000000000000000",
			Token:    "b213d01542d129806d664248a380db8b12059061",
			Nonce:    6,
			From:     "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
			To:       "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
			Op:       "5B38Da6a701c568545dCfcB03FcB875f56beddC4",
			Contract: "1c91347f2A44538ce62453BEBd9Aa907C662b4bD",
			ChkSig:   comn.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		Payvalue: "1000000000000000000",
		// changed for test
		PayCheckSig: comn.String2Byte("227d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
	input.Data = append(input.Data, data0)

	// construct operator
	op, _, _ := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	_, got := op.Aggregate(input)
	if got == nil || got.Error() != "paycheck sig verify failed" {
		t.Error("case 'paycheck sig verify failed' not detected")
	}
}

// test payvalue verify
func TestAggregatePayValue(t *testing.T) {
	input := &pb.SerializeData{}
	data0 := &pb.PayCheck{
		Check: &pb.Check{
			Value:    "100000000000000000000",
			Token:    "b213d01542d129806d664248a380db8b12059061",
			Nonce:    6,
			From:     "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
			To:       "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
			Op:       "5B38Da6a701c568545dCfcB03FcB875f56beddC4",
			Contract: "1c91347f2A44538ce62453BEBd9Aa907C662b4bD",
			ChkSig:   comn.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		Payvalue:    "200000000000000000000", // larger than value
		PayCheckSig: comn.String2Byte("b5b60bc8b85fde998cfe6cf821447e770170780a5e2f7d7b588254284f0e0e3d2040102ddc6abc1b98af5417fa196f445aecf6b85154f45be0b1c3b05bb9cf2800"),
	}
	input.Data = append(input.Data, data0)

	// construct operator
	op, _, _ := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	_, got := op.Aggregate(input)
	if got == nil || got.Error() != "payvalue exceed value" {
		t.Error("case 'payvalue exceed value' not detected")
	}
}

func TestAggregateToAddressIdentical(t *testing.T) {
	input := &pb.SerializeData{}
	data0 := &pb.PayCheck{
		Check: &pb.Check{
			Value:    "100000000000000000000",
			Token:    "b213d01542d129806d664248a380db8b12059061",
			Nonce:    6,
			From:     "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
			To:       "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
			Op:       "5B38Da6a701c568545dCfcB03FcB875f56beddC4",
			Contract: "1c91347f2A44538ce62453BEBd9Aa907C662b4bD",
			ChkSig:   comn.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		Payvalue:    "1000000000000000000",
		PayCheckSig: comn.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
	input.Data = append(input.Data, data0)

	data1 := &pb.PayCheck{
		Check: &pb.Check{
			Value:    "100000000000000000000",
			Token:    "b213d01542d129806d664248a380db8b12059061",
			Nonce:    7,
			From:     "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
			To:       "3320993Bc481177ec7E8f571ceCaE8A9e22C02db", // changed for test
			Op:       "5B38Da6a701c568545dCfcB03FcB875f56beddC4",
			Contract: "1c91347f2A44538ce62453BEBd9Aa907C662b4bD",
			ChkSig:   comn.String2Byte("17d8a014f938995220e861feb51befc2a8bfb8430d91e26ff152b35e2027385b5745c0a89deb5212b811afc9bf4887c53b15f76dade32e48d1e361f682fb208000"),
		},
		Payvalue:    "2000000000000000000",
		PayCheckSig: comn.String2Byte("626d51362677e8757c3dd6b2b1821c80c18cb581073cced1159bca336fd2cb2d05ea51060ab9ad1184bb7c75bfd8ed22bddbc8b2571f3fc7d8b1bd001282299200"),
	}
	input.Data = append(input.Data, data1)

	// construct operator
	op, _, _ := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248A380Db8B12059061")
	_, got := op.Aggregate(input)
	if got == nil || got.Error() != "to address not identical" {
		t.Error("case 'to address not identical' not detected")
	}
}

// test op.gencheck, then usr.genpaycheck
// func TestGenCheck(t *testing.T) {

// 	op, err := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248a380db8b12059061")
// 	if err != nil {
// 		fmt.Println("new operator failed:", err)
// 	}

// 	// input check params to gen check
// 	chk, _ := op.GenCheck(comn.String2BigInt("10000000000000000000"), // value: 10 eth
// 		common.HexToAddress("0xb213d01542d129806d664248A380Db8B12059061"),
// 		common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
// 		common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
// 	)

// 	fmt.Println("value:", chk.Value)
// 	fmt.Println("TokenAddr:", chk.TokenAddr)
// 	fmt.Println("Nonce:", chk.Nonce)
// 	fmt.Println("FromAddr:", chk.FromAddr)
// 	fmt.Println("ToAddr:", chk.ToAddr)
// 	fmt.Println("OpAddr:", chk.OpAddr)
// 	fmt.Println("ContractAddr:", chk.ContractAddr)
// 	fmt.Printf("CheckSig:%x\n", chk.CheckSig)

// 	// use user's sk
// 	usr, err := user.New("7e5bfb82febc4c2c8529167104271ceec190eafdca277314912eaabdb67c6e5f")
// 	if err != nil {
// 		fmt.Println("new user failed:", err)
// 	}

// 	// gen paycheck
// 	pchk, _ := usr.GenPaycheck(chk, comn.String2BigInt("1000000000000000000"))
// 	fmt.Println("----------------------------")
// 	fmt.Println("PayValue:", pchk.PayValue)
// 	fmt.Printf("PaycheckSig:%x\n", pchk.PaycheckSig)

// }

/*
func TestAll(t *testing.T) {

	// new operator
	t.Log("create an operator")
	op, tx, err := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248a380db8b12059061")
	if err != nil {
		t.Log("new operator failed:", err)
	}
	// new user
	t.Log("create an user")
	usr, err := user.New("7e5bfb82febc4c2c8529167104271ceec190eafdca277314912eaabdb67c6e5f")
	if err != nil {
		t.Log("new user failed:", err)
	}
	// new provider
	t.Log("create an provider")
	pro, err := provider.New("cc6d63f85de8fef05446ebdd3c537c72152d0fc437fd7aa62b3019b79bd1fdd4")
	if err != nil {
		t.Log("new provider failed")
		return
	}

	// generate a check, then paycheck
	t.Log("generating check and paycheck..")
	chk, _ := op.GenCheck(comn.String2BigInt("10000000000000000000"), // value: 10 eth
		common.HexToAddress("0xb213d01542d129806d664248A380Db8B12059061"),
		common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
		common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
	)

	fmt.Println("now operator is recording check into it's recorder..")
	// record check into operator's recorder
	op.(*Operator).Recorder.Record(chk)

	// paycheck
	pchk, _ := usr.GenPaycheck(chk, comn.String2BigInt("1000000000000000000"))

	t.Log("check info:")
	t.Log("---------")
	t.Log("value:", chk.Value)
	t.Log("TokenAddr:", chk.TokenAddr)
	t.Log("Nonce:", chk.Nonce)
	t.Log("FromAddr:", chk.FromAddr)
	t.Log("ToAddr:", chk.ToAddr)
	t.Log("OpAddr:", chk.OpAddr)
	t.Log("ContractAddr:", chk.ContractAddr)
	t.Logf("CheckSig:%x\n", chk.CheckSig)
	t.Log("-------")
	t.Log("PayValue:", pchk.PayValue)
	t.Logf("PaycheckSig:%x\n", pchk.PaycheckSig)

	t.Log("user is storing check..")
	//
	if ok, err := usr.VerifyCheck(chk); !ok {
		t.Error("check verify failed:", err)
		return
	} else {
		// record check into user's recorder
		usr.(*user.User).ChkRecorder.Record(chk)
	}

	t.Log("provider is storing paycheck..")
	// record check into provider's recorder
	pro.(*provider.Provider).Store(pchk)

	// connect chain
	t.Log("connecting to chain..")
	ethClient, err := comn.GetClient(pro.(*provider.Provider).Host)
	if err != nil {
		t.Log("get client failed")
		return
	}
	defer ethClient.Close()

	t.Log("now wait contract deploying.")
	// wait contract deployed, txReceipt is checked
	for {
		txReceipt, _ := ethClient.TransactionReceipt(context.Background(), tx.Hash())
		// receipt ok
		if txReceipt != nil {
			break
		}
		t.Log("5 seconds passed..")
		time.Sleep(time.Duration(5) * time.Second)
	}

	t.Log("contract is deployed, now continue.")

	// get old balance of provider
	bal, err := ethClient.BalanceAt(context.Background(), common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"), nil)
	if err != nil {
		t.Log("get balance failed")
	}
	t.Log("balance of provider before withdraw:", bal.String())

	t.Log("now provider is withdrawing..")
	// call withdraw
	tx, _ = pro.WithDraw(pchk)
	// deploy contract, wait for mining.
	for {
		txReceipt, _ := ethClient.TransactionReceipt(context.Background(), tx.Hash())
		// receipt ok
		if txReceipt != nil {
			break
		}
		t.Log("5 seconds passed..")
		time.Sleep(time.Duration(5) * time.Second)
	}

	t.Log("provider withdraw completed, now check the balance again.")
	newbal, err := ethClient.BalanceAt(context.Background(), common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"), nil)
	if err != nil {
		t.Log("get balance failed")
	}
	t.Log("balance of provider after withdraw:", newbal.String())

	// require: newBalance > oldBalance
	if newbal.Cmp(bal) < 0 {
		t.Errorf("new balance should larger than old balance")
	}

	plusGas := new(big.Int)
	plusGas = plusGas.Add(bal, comn.String2BigInt("9000000"))
	total := new(big.Int)
	total = total.Add(plusGas, pchk.PayValue)
	// require: new < old+paycheck.payvalue + gaslimit(9000000)
	if newbal.Cmp(total) > 0 {
		t.Errorf("new balance should smaller than old balance + payvalue + gaslimit")
	}
}
*/
