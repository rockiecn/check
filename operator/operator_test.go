package operator

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
