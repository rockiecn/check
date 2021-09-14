package user

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/rockiecn/check/check"
)

func TestVerifyCheck(t *testing.T) {

	type Input struct {
		value string
		toekn string
		nonce string
		from  string
		to    string
		op    string
		con   string
		sig   string
	}

	var tests = []struct {
		input Input
		want  bool
	}{
		{
			Input{
				"100000000000000000000",
				"b213d01542d129806d664248a380db8b12059061",
				"5",
				"Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
				"4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
				"5B38Da6a701c568545dCfcB03FcB875f56beddC4",
				"aE036c65C649172b43ef7156b009c6221B596B8b",
				"66cec089a3e9d86cc98f829fcf6ed74b6f8bd8537f9ee4eee4c7d8f51fd3fbcf3408429ce1d84a9d107d2e8f1c9730b463b05de5b8f7f221ae5095c8ec58234501",
			},
			true,
		},
		{
			Input{
				"200100000000000000000",
				"b213d01542d129806d664248a380db8b12059061",
				"5",
				"Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
				"4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
				"5B38Da6a701c568545dCfcB03FcB875f56beddC4",
				"aE036c65C649172b43ef7156b009c6221B596B8b",
				"66cec089a3e9d86cc98f829fcf6ed74b6f8bd8537f9ee4eee4c7d8f51fd3fbcf3408429ce1d84a9d107d2e8f1c9730b463b05de5b8f7f221ae5095c8ec58234501",
			},
			false,
		},
	}

	user := new(User)

	for _, test := range tests {

		check := new(check.Check)
		bigValue := new(big.Int)
		bigValue.SetString(test.input.value, 10)
		check.Value = bigValue

		check.TokenAddr = test.input.toekn

		bigNonce := new(big.Int)
		bigNonce.SetString(test.input.nonce, 10)
		check.Nonce = bigNonce

		check.From = test.input.from
		check.To = test.input.to
		check.OperatorAddr = test.input.op
		check.ContractAddr = test.input.con
		sigByte, err := hex.DecodeString(test.input.sig)
		check.CheckSig = sigByte

		got, err := user.VerifyCheck(check)
		if err != nil {
			fmt.Println("verify check failed:", err)
			return
		}
		if got != test.want {
			t.Errorf("verifycheck, want:%v got:%v", test.want, got)
		}
	}
}

func TestSign(t *testing.T) {

	type PC struct {
		Value        string
		TokenAddr    string
		Nonce        string
		From         string
		To           string
		OperatorAddr string
		ContractAddr string
		CheckSig     string
		PayValue     string
		PayCheckSig  string
	}

	var tests = []struct {
		input PC
		want  bool
	}{
		{
			PC{
				"100000000000000000000",
				"b213d01542d129806d664248a380db8b12059061",
				"5",
				"Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
				"4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
				"5B38Da6a701c568545dCfcB03FcB875f56beddC4",
				"aE036c65C649172b43ef7156b009c6221B596B8b",
				"66cec089a3e9d86cc98f829fcf6ed74b6f8bd8537f9ee4eee4c7d8f51fd3fbcf3408429ce1d84a9d107d2e8f1c9730b463b05de5b8f7f221ae5095c8ec58234501",
				"0",
				"c21324bab3b3e75006318dd8fff44b03078b5b162d9b20b27f3811f29de888404abf9416648bc153905d644c292b7ca44b240eedacad8b4370654dc32e91946800",
			},
			true,
		},

		{
			PC{
				"200000000000000000000",
				"b213d01542d129806d664248a380db8b12059061",
				"5",
				"Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2",
				"4B20993Bc481177ec7E8f571ceCaE8A9e22C02db",
				"5B38Da6a701c568545dCfcB03FcB875f56beddC4",
				"aE036c65C649172b43ef7156b009c6221B596B8b",
				"66cec089a3e9d86cc98f829fcf6ed74b6f8bd8537f9ee4eee4c7d8f51fd3fbcf3408429ce1d84a9d107d2e8f1c9730b463b05de5b8f7f221ae5095c8ec58234501",
				"0",
				"c21324bab3b3e75006318dd8fff44b03078b5b162d9b20b27f3811f29de888404abf9416648bc153905d644c292b7ca44b240eedacad8b4370654dc32e91946800",
			},
			false,
		},
	}

	user, err := NewUser("7e5bfb82febc4c2c8529167104271ceec190eafdca277314912eaabdb67c6e5f")
	if err != nil {
		fmt.Println("new user failed:", err)
	}

	for _, test := range tests {

		bigValue, _ := new(big.Int).SetString(test.input.Value, 10)
		bigNonce, _ := new(big.Int).SetString(test.input.Nonce, 10)
		bigPV, _ := new(big.Int).SetString(test.input.PayValue, 10)

		chk := new(check.Check)

		chk.Value = bigValue
		chk.TokenAddr = test.input.TokenAddr
		chk.Nonce = bigNonce
		chk.From = test.input.From
		chk.To = test.input.To
		chk.OperatorAddr = test.input.OperatorAddr
		chk.ContractAddr = test.input.ContractAddr
		sigByte, _ := hex.DecodeString(test.input.CheckSig)
		chk.CheckSig = sigByte

		pchk := new(check.PayCheck)
		pchk.Check = chk
		pchk.PayValue = bigPV

		sig, err := user.Sign(pchk)
		if err != nil {
			fmt.Println("sign paycheck failed:", err)
			return
		}

		sigByte, _ = hex.DecodeString(test.input.PayCheckSig)
		got := bytes.Equal(sig, sigByte)

		if got != test.want {
			t.Errorf("sign, want:%v got:%v", test.want, got)
		}
	}
}
