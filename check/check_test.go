package check

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestCheckSign(t *testing.T) {

	type test struct {
		chk  Check
		want bool
	}

	var tests = []test{
		{
			chk: Check{
				Value:        String2BigInt("100000000000000000000"),
				TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
				Nonce:        19,
				FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				ContractAddr: common.HexToAddress("406AB5033423Dcb6391Ac9eEEad73294FA82Cfbc"),
				CheckSig:     String2Byte("18dbfea51d279adb0ae16dff88b0dd341cd06838970779cfe056a2f7d359f4c431fb61168e261aed805b285e523e0092b5d46a2c51e6b8b093c5a151abf24b0100"),
			},
			want: true,
		},
		{
			chk: Check{
				Value:        String2BigInt("200000000000000000000"),
				TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
				Nonce:        19,
				FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				ContractAddr: common.HexToAddress("406AB5033423Dcb6391Ac9eEEad73294FA82Cfbc"),
				CheckSig:     String2Byte("18dbfea51d279adb0ae16dff88b0dd341cd06838970779cfe056a2f7d359f4c431fb61168e261aed805b285e523e0092b5d46a2c51e6b8b093c5a151abf24b0100"),
			},
			want: false,
		},
	}

	// operator's sk
	sk := "503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb"

	for _, test := range tests {

		newChk := new(Check)
		*newChk = test.chk
		newChk.Sign(sk)

		got := bytes.Equal(newChk.CheckSig, test.chk.CheckSig)

		if got != test.want {
			t.Errorf("want: %v, got: %v", test.want, got)
		}
	}
}

func TestPaycheckSign(t *testing.T) {

	type test struct {
		pchk Paycheck
		want bool
	}

	var tests = []test{
		{
			pchk: Paycheck{
				Check: Check{
					Value:        String2BigInt("100000000000000000000"),
					TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
					Nonce:        19,
					FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
					ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
					OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
					ContractAddr: common.HexToAddress("406AB5033423Dcb6391Ac9eEEad73294FA82Cfbc"),
					CheckSig:     String2Byte("18dbfea51d279adb0ae16dff88b0dd341cd06838970779cfe056a2f7d359f4c431fb61168e261aed805b285e523e0092b5d46a2c51e6b8b093c5a151abf24b0100"),
				},
				PayValue:    String2BigInt("1000000000000000000"),
				PaycheckSig: String2Byte("142673fc986aa6d456fba6df9ce82cfdc92cf9af7ea169c41787cf59fcaba9b2721b0b72a121dadd072b1f33743cafb21920e0bea2fe6b435b22ffb715c49d6001"),
			},
			want: true,
		},
		{
			pchk: Paycheck{
				Check: Check{
					Value:        String2BigInt("200000000000000000000"),
					TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
					Nonce:        19,
					FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
					ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
					OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
					ContractAddr: common.HexToAddress("406AB5033423Dcb6391Ac9eEEad73294FA82Cfbc"),
					CheckSig:     String2Byte("18dbfea51d279adb0ae16dff88b0dd341cd06838970779cfe056a2f7d359f4c431fb61168e261aed805b285e523e0092b5d46a2c51e6b8b093c5a151abf24b0100"),
				},
				PayValue:    String2BigInt("1000000000000000000"),
				PaycheckSig: String2Byte("142673fc986aa6d456fba6df9ce82cfdc92cf9af7ea169c41787cf59fcaba9b2721b0b72a121dadd072b1f33743cafb21920e0bea2fe6b435b22ffb715c49d6001"),
			},
			want: false,
		},
	}

	// user's sk
	sk := "7e5bfb82febc4c2c8529167104271ceec190eafdca277314912eaabdb67c6e5f"

	for _, test := range tests {

		newpchk := new(Paycheck)
		*newpchk = test.pchk
		newpchk.Sign(sk)

		got := bytes.Equal(newpchk.PaycheckSig, test.pchk.PaycheckSig)

		if got != test.want {
			t.Errorf("want: %v, got: %v", test.want, got)
		}
	}
}

func TestCheckVerify(t *testing.T) {

	type test struct {
		chk  Check
		want bool
	}

	var tests = []test{
		{
			chk: Check{
				Value:        String2BigInt("100000000000000000000"),
				TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
				Nonce:        19,
				FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				ContractAddr: common.HexToAddress("406AB5033423Dcb6391Ac9eEEad73294FA82Cfbc"),
				CheckSig:     String2Byte("18dbfea51d279adb0ae16dff88b0dd341cd06838970779cfe056a2f7d359f4c431fb61168e261aed805b285e523e0092b5d46a2c51e6b8b093c5a151abf24b0100"),
			},
			want: true,
		},
		{
			chk: Check{
				Value:        String2BigInt("100000000000000000000"),
				TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
				Nonce:        19,
				FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				ContractAddr: common.HexToAddress("406AB5033423Dcb6391Ac9eEEad73294FA82Cfbc"),
				CheckSig:     String2Byte("142673fc986aa6d456fba6df9ce82cfdc92cf9af7ea169c41787cf59fcaba9b2721b0b72a121dadd072b1f33743cafb21920e0bea2fe6b435b22ffb715c49d6001"),
			},
			want: false,
		},
	}

	for _, test := range tests {

		got, _ := test.chk.Verify()

		if got != test.want {
			t.Errorf("want: %v, got: %v", test.want, got)
		}
	}
}

func TestPaycheckVerify(t *testing.T) {

	type test struct {
		pchk Paycheck
		want bool
	}

	var tests = []test{
		{
			pchk: Paycheck{
				Check: Check{
					Value:        String2BigInt("100000000000000000000"),
					TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
					Nonce:        19,
					FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
					ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
					OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
					ContractAddr: common.HexToAddress("406AB5033423Dcb6391Ac9eEEad73294FA82Cfbc"),
					CheckSig:     String2Byte("18dbfea51d279adb0ae16dff88b0dd341cd06838970779cfe056a2f7d359f4c431fb61168e261aed805b285e523e0092b5d46a2c51e6b8b093c5a151abf24b0100"),
				},
				PayValue:    String2BigInt("1000000000000000000"),
				PaycheckSig: String2Byte("142673fc986aa6d456fba6df9ce82cfdc92cf9af7ea169c41787cf59fcaba9b2721b0b72a121dadd072b1f33743cafb21920e0bea2fe6b435b22ffb715c49d6001"),
			},
			want: true,
		},
		{
			pchk: Paycheck{
				Check: Check{
					Value:        String2BigInt("200000000000000000000"),
					TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
					Nonce:        19,
					FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
					ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
					OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
					ContractAddr: common.HexToAddress("406AB5033423Dcb6391Ac9eEEad73294FA82Cfbc"),
					CheckSig:     String2Byte("18dbfea51d279adb0ae16dff88b0dd341cd06838970779cfe056a2f7d359f4c431fb61168e261aed805b285e523e0092b5d46a2c51e6b8b093c5a151abf24b0100"),
				},
				PayValue:    String2BigInt("1000000000000000000"),
				PaycheckSig: String2Byte("56fba6df9ce82cfdc92cf56fba6df9ce82cfdc92cf56fba6df9ce82cfdc92cf56fba6df9ce82cfdc92cf56fba6df9ce82cfdc92cf56fba6df9ce82cfdc92cf56fb"),
			},
			want: false,
		},
	}

	for _, test := range tests {

		got, _ := test.pchk.Verify()

		if got != test.want {
			t.Errorf("want: %v, got: %v", test.want, got)
		}
	}
}

func String2BigInt(str string) *big.Int {
	v, _ := big.NewInt(0).SetString(str, 10)
	return v
}

func String2Byte(str string) []byte {
	v, _ := hex.DecodeString(str)
	return v
}
