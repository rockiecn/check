package check

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/utils"
)

func TestCheckSign(t *testing.T) {

	type test struct {
		chk  Check
		want bool
	}

	var tests = []test{
		{
			chk: Check{
				CheckInfo: CheckInfo{
					Value:     utils.String2BigInt("100000000000000000000"),
					TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
					Nonce:     6,
					FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
					ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
					OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
					CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
				},
				CheckSig: utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
			},
			want: true,
		},
		{
			chk: Check{
				CheckInfo: CheckInfo{
					Value:     utils.String2BigInt("100000000000000000000"),
					TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
					Nonce:     7,
					FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
					ToAddr:    common.HexToAddress("3320993Bc481177ec7E8f571ceCaE8A9e22C02db"),
					OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
					CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
				},
				CheckSig: utils.String2Byte("17d8a014f938995220e861feb51befc2a8bfb8430d91e26ff152b35e2027385b5745c0a89deb5212b811afc9bf4887c53b15f76dade32e48d1e361f682fb208000"),
			},
			want: true,
		},
		{
			chk: Check{
				CheckInfo: CheckInfo{
					Value:     utils.String2BigInt("100000000000000000000"),
					TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
					Nonce:     8,
					FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
					ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
					OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
					CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
				},
				CheckSig: utils.String2Byte("584cca0e6eed3558bd07e1ab40206ecc83dc005ccad16ea9d97586726ec43aeb486a6599e1c77b345ce73c4f7f4c26e78230b752a3f3e42de62c9da261f5923e00"),
			},
			want: true,
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
		input Paycheck
		want  bool
	}

	var tests = []test{
		{
			input: Paycheck{
				Check: &Check{
					CheckInfo: CheckInfo{
						Value:     utils.String2BigInt("100000000000000000000"),
						TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
						Nonce:     6,
						FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
						ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
						OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
						CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
					},
					CheckSig: utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
				},
				PayValue:    utils.String2BigInt("1000000000000000000"),
				PaycheckSig: utils.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
			},
			want: true,
		},
		{
			input: Paycheck{
				Check: &Check{
					CheckInfo: CheckInfo{
						Value:     utils.String2BigInt("100000000000000000000"),
						TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
						Nonce:     7,
						FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
						ToAddr:    common.HexToAddress("3320993Bc481177ec7E8f571ceCaE8A9e22C02db"),
						OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
						CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
					},
					CheckSig: utils.String2Byte("17d8a014f938995220e861feb51befc2a8bfb8430d91e26ff152b35e2027385b5745c0a89deb5212b811afc9bf4887c53b15f76dade32e48d1e361f682fb208000"),
				},
				PayValue:    utils.String2BigInt("2000000000000000000"),
				PaycheckSig: utils.String2Byte("626d51362677e8757c3dd6b2b1821c80c18cb581073cced1159bca336fd2cb2d05ea51060ab9ad1184bb7c75bfd8ed22bddbc8b2571f3fc7d8b1bd001282299200"),
			},
			want: true,
		},
		{
			input: Paycheck{
				Check: &Check{
					CheckInfo: CheckInfo{
						Value:     utils.String2BigInt("100000000000000000000"),
						TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
						Nonce:     8,
						FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
						ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
						OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
						CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
					},
					CheckSig: utils.String2Byte("584cca0e6eed3558bd07e1ab40206ecc83dc005ccad16ea9d97586726ec43aeb486a6599e1c77b345ce73c4f7f4c26e78230b752a3f3e42de62c9da261f5923e00"),
				},
				PayValue:    utils.String2BigInt("3000000000000000000"),
				PaycheckSig: utils.String2Byte("c75f9b4f960be6bb48719a01da40370afd00e905a554d54e6adc1fe41ab09ccc149701daf4f1118316997f4b9b748b7ff690c1e9dd653f3e63cb0a3fffa625b300"),
			},
			want: true,
		},
	}

	// user's sk
	sk := "7e5bfb82febc4c2c8529167104271ceec190eafdca277314912eaabdb67c6e5f"

	for _, test := range tests {

		newpchk := new(Paycheck)
		*newpchk = test.input
		newpchk.Sign(sk)

		got := bytes.Equal(newpchk.PaycheckSig, test.input.PaycheckSig)

		if got != test.want {
			t.Errorf("want: %v, got: %v", test.want, got)
		}
	}
}

// a tool for new paycheck case
func TestPaycheckTool(t *testing.T) {

	input := &Paycheck{
		Check: &Check{
			CheckInfo: CheckInfo{
				Value:     utils.String2BigInt("100000000000000000000"),
				TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
				Nonce:     6,
				FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			},
		},
		PayValue: utils.String2BigInt("1000000000000000000"),
	}

	// operator sk
	opSK := "503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb"
	// user's sk
	usrSK := "7e5bfb82febc4c2c8529167104271ceec190eafdca277314912eaabdb67c6e5f"

	// sign check
	input.Check.Sign(opSK)
	// sign paycheck
	input.Sign(usrSK)

	fmt.Printf("checksig: %x\n", input.Check.CheckSig)
	fmt.Printf("paychecksig: %x\n", input.PaycheckSig)
}

var chk0 = &Check{
	CheckInfo: CheckInfo{
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

func TestKey(t *testing.T) {
	k := utils.ToKey(chk0.ToAddr, chk0.Nonce)
	addr, n := utils.FromKey(k)
	fmt.Printf("addr:%x\n, nonce:%d\n", addr, n)
}

/*
func TestCheckVerify(t *testing.T) {

	type test struct {
		chk  Check
		want bool
	}

	var tests = []test{
		{
			chk: Check{
				Value:        internal.String2BigInt("100000000000000000000"),
				TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
				Nonce:        6,
				FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				CtrAddr: common.HexToAddress("406AB5033423Dcb6391Ac9eEEad73294FA82Cfbc"),
				CheckSig:     internal.String2Byte("18dbfea51d279adb0ae16dff88b0dd341cd06838970779cfe056a2f7d359f4c431fb61168e261aed805b285e523e0092b5d46a2c51e6b8b093c5a151abf24b0100"),
			},
			want: true,
		},
		{
			chk: Check{
				Value:        internal.String2BigInt("100000000000000000000"),
				TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
				Nonce:        19,
				FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				CtrAddr: common.HexToAddress("406AB5033423Dcb6391Ac9eEEad73294FA82Cfbc"),
				CheckSig:     internal.String2Byte("142673fc986aa6d456fba6df9ce82cfdc92cf9af7ea169c41787cf59fcaba9b2721b0b72a121dadd072b1f33743cafb21920e0bea2fe6b435b22ffb715c49d6001"),
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

/*
func TestPaycheckVerify(t *testing.T) {

	type test struct {
		pchk Paycheck
		want bool
	}

	var tests = []test{
		{
			pchk: Paycheck{
				Check: Check{
					Value:        internal.String2BigInt("100000000000000000000"),
					TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
					Nonce:        19,
					FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
					ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
					OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
					CtrAddr: common.HexToAddress("406AB5033423Dcb6391Ac9eEEad73294FA82Cfbc"),
					CheckSig:     internal.String2Byte("18dbfea51d279adb0ae16dff88b0dd341cd06838970779cfe056a2f7d359f4c431fb61168e261aed805b285e523e0092b5d46a2c51e6b8b093c5a151abf24b0100"),
				},
				PayValue:    internal.String2BigInt("1000000000000000000"),
				PaycheckSig: internal.String2Byte("142673fc986aa6d456fba6df9ce82cfdc92cf9af7ea169c41787cf59fcaba9b2721b0b72a121dadd072b1f33743cafb21920e0bea2fe6b435b22ffb715c49d6001"),
			},
			want: true,
		},
		{
			pchk: Paycheck{
				Check: Check{
					Value:        internal.String2BigInt("200000000000000000000"),
					TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
					Nonce:        19,
					FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
					ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
					OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
					CtrAddr: common.HexToAddress("406AB5033423Dcb6391Ac9eEEad73294FA82Cfbc"),
					CheckSig:     internal.String2Byte("18dbfea51d279adb0ae16dff88b0dd341cd06838970779cfe056a2f7d359f4c431fb61168e261aed805b285e523e0092b5d46a2c51e6b8b093c5a151abf24b0100"),
				},
				PayValue:    internal.String2BigInt("1000000000000000000"),
				PaycheckSig: internal.String2Byte("56fba6df9ce82cfdc92cf56fba6df9ce82cfdc92cf56fba6df9ce82cfdc92cf56fba6df9ce82cfdc92cf56fba6df9ce82cfdc92cf56fba6df9ce82cfdc92cf56fb"),
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
*/
