package recorder

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
	comn "github.com/rockiecn/check/common"
)

func TestRecord(t *testing.T) {

	// all cases
	type oneCase struct {
		data interface{}
		want error
	}

	cases := []oneCase{
		// *check.Check case
		{
			data: &check.Check{
				Value:        comn.String2BigInt("10000000000000000000"),
				TokenAddr:    common.HexToAddress("0xb213d01542d129806d664248A380Db8B12059061"),
				Nonce:        0,
				FromAddr:     common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ToAddr:       common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				OpAddr:       common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				ContractAddr: common.HexToAddress("0xb31BA5cDC07A2EaFAF77c95294fd4aE27D04E9CA"),
				CheckSig:     comn.String2Byte("456391994ed07bd03809b03a7afd9cdd4554c4a6c382289f7e4ea6c749afa7014937817cd32c0c8f983375dbae9959756e82f855538375d0bb99813a0770323500"),
			},
			want: nil,
		},
		// *check.Paycheck case
		{
			data: &check.Paycheck{
				Check: check.Check{
					Value:        comn.String2BigInt("10000000000000000000"),
					TokenAddr:    common.HexToAddress("0xb213d01542d129806d664248A380Db8B12059061"),
					Nonce:        0,
					FromAddr:     common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
					ToAddr:       common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
					OpAddr:       common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
					ContractAddr: common.HexToAddress("0xb31BA5cDC07A2EaFAF77c95294fd4aE27D04E9CA"),
					CheckSig:     comn.String2Byte("456391994ed07bd03809b03a7afd9cdd4554c4a6c382289f7e4ea6c749afa7014937817cd32c0c8f983375dbae9959756e82f855538375d0bb99813a0770323500"),
				},
				PayValue:    comn.String2BigInt("1000000000000000000"),
				PaycheckSig: comn.String2Byte("2a8488c7a44fefe771b9fcc1c0ab801f95611906f12e58c4fe2e2ac9406dc97c1601c2ae6f328c4a18df90c1c33170872cab29e1ef970644b163d117a9819ac100"),
			},
			want: nil,
		},
		// check.Paycheck case
		{
			data: check.Paycheck{
				Check: check.Check{
					Value:        comn.String2BigInt("10000000000000000000"),
					TokenAddr:    common.HexToAddress("0xb213d01542d129806d664248A380Db8B12059061"),
					Nonce:        0,
					FromAddr:     common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
					ToAddr:       common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
					OpAddr:       common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
					ContractAddr: common.HexToAddress("0xb31BA5cDC07A2EaFAF77c95294fd4aE27D04E9CA"),
					CheckSig:     comn.String2Byte("456391994ed07bd03809b03a7afd9cdd4554c4a6c382289f7e4ea6c749afa7014937817cd32c0c8f983375dbae9959756e82f855538375d0bb99813a0770323500"),
				},
				PayValue:    comn.String2BigInt("1000000000000000000"),
				PaycheckSig: comn.String2Byte("2a8488c7a44fefe771b9fcc1c0ab801f95611906f12e58c4fe2e2ac9406dc97c1601c2ae6f328c4a18df90c1c33170872cab29e1ef970644b163d117a9819ac100"),
			},
			want: errors.New("type of value must be Check or Paycheck"),
		},
	}

	rec := New()

	// test all cases
	for _, c := range cases {
		got := rec.Record(c.data)
		if got == nil && c.want == nil {
			//t.Log("want err nil, got err nil, test OK")
			continue
		}

		if got == nil && c.want != nil {
			t.Errorf("want: %v, got: %v", c.want.Error(), nil)
		}

		if got != nil && c.want == nil {
			t.Errorf("want: %v, got: %v", nil, got.Error())
		}

		if got != nil || c.want != nil {
			if got.Error() != c.want.Error() {
				t.Errorf("want: %v, got: %v", c.want.Error(), got.Error())
			}
		}
	}
}

func TestIsValid(t *testing.T) {

	type test struct {
		input interface{}
		ret   bool
		e     error
	}

	// check/payckeck
	tests := []test{
		{
			input: &check.Check{
				Nonce:  1,
				ToAddr: common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa"),
			},
			ret: false,
			e:   nil,
		},
		{
			input: &check.Check{
				Nonce:  3,
				ToAddr: common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa"),
			},
			ret: true,
			e:   nil,
		},
		{
			input: &check.Paycheck{
				Check: check.Check{
					Nonce:  4,
					ToAddr: common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa"),
				},
				PayValue: big.NewInt(10),
			},
			ret: false,
			e:   nil,
		},
		{
			input: &check.Paycheck{
				Check: check.Check{
					Nonce:  4,
					ToAddr: common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa"),
				},
				PayValue: big.NewInt(30),
			},
			ret: true,
			e:   nil,
		},
		{
			input: check.Paycheck{
				Check: check.Check{
					Nonce:  4,
					ToAddr: common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa"),
				},
				PayValue: big.NewInt(30),
			},
			ret: false,
			e:   errors.New("type must be check/paycheck"),
		},
	}

	// data to be put into Entrys
	records := []interface{}{

		&check.Check{
			Nonce:  1,
			ToAddr: common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa"),
		},

		&check.Check{
			Nonce:  2,
			ToAddr: common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa"),
		},
		&check.Check{
			Nonce:  1,
			ToAddr: common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
		},
		&check.Paycheck{
			Check: check.Check{
				Nonce:  4,
				ToAddr: common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa"),
			},
			PayValue: big.NewInt(20),
		},
	}

	// new Record
	rec := New()

	// prepair for testing, put all testing data into Entrys
	for _, r := range records {

		err := rec.Record(r)
		if err != nil {
			t.Fatal("got error when Recording")
		}
	}

	// begin test
	for _, tst := range tests {

		gotRet, gotErr := rec.IsValid(tst.input)

		if gotErr != nil {
			if tst.e == nil {
				t.Errorf("want error: nil, got error: %v", gotErr)
			} else {
				if tst.e.Error() != gotErr.Error() || tst.ret != gotRet {
					t.Errorf("want: %v,%v, got: %v,%v", tst.ret, tst.e.Error(), gotRet, gotErr.Error())
				}
			}
		} else {
			if tst.e != nil {
				t.Errorf("want err: %v, got err: %v", tst.e, gotErr)
			}
			if gotRet != tst.ret {
				t.Errorf("want: %v, got: %v", tst.ret, gotRet)
			}
		}
	}
}
