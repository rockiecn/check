package operator

import (
	"testing"
)

// import (
// 	"bytes"
// 	"encoding/hex"
// 	"fmt"
// 	"math/big"
// 	"testing"

// 	"github.com/rockiecn/check/check"
// 	"github.com/rockiecn/check/utils"
// )

// func TestKeyToAddr(t *testing.T) {
// 	var tests = []struct {
// 		input string
// 		want  string
// 	}{
// 		{"503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb",
// 			"0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"},
// 		{"7e5bfb82febc4c2c8529167104271ceec190eafdca277314912eaabdb67c6e5f",
// 			"0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"},
// 		{"cc6d63f85de8fef05446ebdd3c537c72152d0fc437fd7aa62b3019b79bd1fdd4",
// 			"0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"},
// 	}

// 	for _, test := range tests {
// 		if got := utils.KeyToAddr(test.input); got != test.want {
// 			t.Errorf("want:%q got:%q", test.want, got)
// 		}
// 	}
// }

// func TestSign(t *testing.T) {

// 	type Input struct {
// 		Value        string
// 		TokenAddr    string
// 		Nonce        string
// 		From         string
// 		To           string
// 		OperatorAddr string
// 		ContractAddr string
// 		OpSk         string
// 		CheckSig     string
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
// 				"aE036c65C649172b43ef7156b009c6221B596B8b",
// 				"503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb",
// 				"66cec089a3e9d86cc98f829fcf6ed74b6f8bd8537f9ee4eee4c7d8f51fd3fbcf3408429ce1d84a9d107d2e8f1c9730b463b05de5b8f7f221ae5095c8ec58234501",
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
// 				"aE036c65C649172b43ef7156b009c6221B596B8b",
// 				"503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb",
// 				"66cec089a3e9d86cc98f829fcf6ed74b6f8bd8537f9ee4eee4c7d8f51fd3fbcf3408429ce1d84a9d107d2e8f1c9730b463b05de5b8f7f221ae5095c8ec58234501",
// 			},
// 			false,
// 		},
// 	}

// 	op, err := NewOperator(
// 		"503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb",
// 		"0x9e0153496067c20943724b79515472195a7aedaa")
// 	if err != nil {
// 		fmt.Println("new operator failed:", err)
// 	}

// 	for _, test := range tests {
// 		bigValue := big.NewInt(0)
// 		bigValue.SetString(test.input.Value, 0)

// 		bigNonce := big.NewInt(0)
// 		bigNonce.SetString(test.input.Nonce, 0)

// 		chk := new(check.Check)

// 		chk.Value = bigValue
// 		chk.TokenAddr = test.input.TokenAddr
// 		chk.Nonce = bigNonce
// 		chk.From = test.input.From
// 		chk.To = test.input.To
// 		chk.OperatorAddr = test.input.OperatorAddr
// 		chk.ContractAddr = test.input.ContractAddr

// 		// decode string to []byte
// 		sigByte, _ := hex.DecodeString(test.input.CheckSig)
// 		calcSig, _ := op.Sign(chk)
// 		got := bytes.Equal(calcSig, sigByte)
// 		if err != nil {
// 			fmt.Print("sign error:", err)
// 			return
// 		}

// 		if got != test.want {
// 			t.Errorf("want: %v, got: %v", test.want, got)
// 		}
// 	}

// }

// func TestDeploy(t *testing.T) {
// 	op, err := NewOperator(
// 		"503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb",
// 		"0x9e0153496067c20943724b79515472195a7aedaa")
// 	if err != nil {
// 		fmt.Println("new operator failed:", err)
// 	}

// 	txHash, comAddr, err := op.DeployContract()

// 	if err != nil {
// 		t.Errorf("deploy contract failed")
// 	}

// 	_, _ = txHash, comAddr
// }

func Test(t *testing.T) {
	_, _ = New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248a380db8b12059061")

}
