// package provider

// import (
// 	"encoding/hex"
// 	"fmt"
// 	"math/big"
// 	"testing"

// 	"github.com/rockiecn/check/check"
// 	"github.com/rockiecn/check/operator"
// 	"github.com/rockiecn/check/user"
// )

// func TestVerifyPayCheck(t *testing.T) {

// 	type Input struct {
// 		value       string
// 		toekn       string
// 		nonce       string
// 		from        string
// 		to          string
// 		op          string
// 		con         string
// 		checksig    string
// 		payvalue    string
// 		payckecksig string
// 	}

// 	var tests = []struct {
// 		input Input
// 		want  bool
// 	}{
// 		{
// 			Input{
// 				"100000000000000000000",
// 				"b213d01542d129806d664248a380db8b12059061",
// 				"5",
// 				"Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
// 				"4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
// 				"5B38Da6a701c568545dCfcB03FcB875f56beddC4",
// 				"cD6a42782d230D7c13A74ddec5dD140e55499Df9",
// 				"688ebe9157c8a338cac3c5505440b9a0f84b7a87ca9e30188422d87b39dd2e2f452f7b953e4129b41b967e7818441e0346ff096b63b916e2aee45591797ab1e700",
// 				"0",
// 				"6f0fd75a9d7150f8189324373f659cc441f36768396407131a2c5ca5ed57b3ed02e474289050efc698531560c1b0914e52ca49b9c98e610a3de86c7e6e77c48100",
// 			},
// 			true,
// 		},
// 		{
// 			Input{
// 				"200000000000000000000",
// 				"b213d01542d129806d664248a380db8b12059061",
// 				"5",
// 				"Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
// 				"4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
// 				"5B38Da6a701c568545dCfcB03FcB875f56beddC4",
// 				"cD6a42782d230D7c13A74ddec5dD140e55499Df9",
// 				"688ebe9157c8a338cac3c5505440b9a0f84b7a87ca9e30188422d87b39dd2e2f452f7b953e4129b41b967e7818441e0346ff096b63b916e2aee45591797ab1e700",
// 				"0",
// 				"6f0fd75a9d7150f8189324373f659cc441f36768396407131a2c5ca5ed57b3ed02e474289050efc698531560c1b0914e52ca49b9c98e610a3de86c7e6e77c48100",
// 			},
// 			false,
// 		},
// 	}

// 	pro, err := NewProvider("cc6d63f85de8fef05446ebdd3c537c72152d0fc437fd7aa62b3019b79bd1fdd4")
// 	if err != nil {
// 		fmt.Println("new provider failed:", err)
// 	}

// 	for _, test := range tests {

// 		chk := new(check.Check)
// 		bigValue := new(big.Int)
// 		bigValue.SetString(test.input.value, 10)
// 		chk.Value = bigValue

// 		chk.TokenAddr = test.input.toekn

// 		bigNonce := new(big.Int)
// 		bigNonce.SetString(test.input.nonce, 10)
// 		chk.Nonce = bigNonce

// 		chk.From = test.input.from
// 		chk.To = test.input.to
// 		chk.OperatorAddr = test.input.op
// 		chk.ContractAddr = test.input.con
// 		sigByte, _ := hex.DecodeString(test.input.checksig)
// 		chk.CheckSig = sigByte

// 		paycheck := new(check.PayCheck)
// 		paycheck.Check = chk
// 		bigPayValue := new(big.Int)
// 		bigPayValue.SetString(test.input.payvalue, 10)
// 		paycheck.PayValue = bigPayValue
// 		sigByte, _ = hex.DecodeString(test.input.payckecksig)
// 		paycheck.PayCheckSig = sigByte

// 		got, err := pro.VerifyPayCheck(paycheck)
// 		if err != nil {
// 			fmt.Println("verify check failed:", err)
// 			return
// 		}
// 		if got != test.want {
// 			t.Errorf("verifycheck, want:%v got:%v", test.want, got)
// 		}
// 	}
// }

// func TestCallContract(t *testing.T) {

// 	op, _ := operator.NewOperator(
// 		"503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb",
// 		"b213d01542d129806d664248A380Db8B12059061")
// 	pro, _ := NewProvider(
// 		"cc6d63f85de8fef05446ebdd3c537c72152d0fc437fd7aa62b3019b79bd1fdd4")

// 	// construct paycheck
// 	bigValue, _ := new(big.Int).SetString("100000000000000000000", 10)
// 	bigNonce, _ := new(big.Int).SetString("19", 10)

// 	chk := &check.Check{
// 		Value:        bigValue,
// 		TokenAddr:    "b213d01542d129806d664248a380db8b12059061",
// 		Nonce:        bigNonce,
// 		From:         "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
// 		To:           "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
// 		OperatorAddr: "5B38Da6a701c568545dCfcB03FcB875f56beddC4",
// 		ContractAddr: "0498B7c793D7432Cd9dB27fb02fc9cfdBAfA1Fd3",
// 		CheckSig:     []byte{},
// 	}

// 	chkSig, _ := op.Sign(chk)
// 	chk.CheckSig = chkSig

// 	bigPayValue, _ := new(big.Int).SetString("1000000000000000000", 10)

// 	pchk := &check.PayCheck{
// 		Check:       chk,
// 		PayValue:    bigPayValue,
// 		PayCheckSig: []byte{},
// 	}
// 	user, _ := user.NewUser("7e5bfb82febc4c2c8529167104271ceec190eafdca277314912eaabdb67c6e5f")
// 	pchkSig, _ := user.Sign(pchk)
// 	pchk.PayCheckSig = pchkSig

// 	// call contract with this paycheck
// 	pro.CallContract(pchk)
// }
