package recorder

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
)

func TestIsValid(t *testing.T) {

	type test struct {
		input interface{}
		want  bool
	}

	// check/payckeck
	tests := []test{
		{
			input: &check.Check{
				Nonce:  1,
				ToAddr: common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa"),
			},
			want: false,
		},
		{
			input: &check.Check{
				Nonce:  3,
				ToAddr: common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa"),
			},
			want: true,
		},
		{
			input: &check.Paycheck{
				Check: check.Check{
					Nonce:  4,
					ToAddr: common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa"),
				},
				PayValue: big.NewInt(10),
			},
			want: false,
		},
		{
			input: &check.Paycheck{
				Check: check.Check{
					Nonce:  4,
					ToAddr: common.HexToAddress("0x9e0153496067c20943724b79515472195a7aedaa"),
				},
				PayValue: big.NewInt(30),
			},
			want: true,
		},
	}

	type record interface{}

	// data to be put into Entrys
	records := []record{

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

		refType := reflect.TypeOf((r))

		switch refType.String() {
		case "*check.Check":
			// cast type to check.Check for calling Record()
			v := r.(*check.Check)
			err := rec.Record(v)
			if err != nil {
				fmt.Println(err)
			}
		case "*check.Paycheck":
			// cast type to check.Paycheck for calling Record()
			v := r.(*check.Paycheck)
			err := rec.Record(v)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	// begin test
	for _, tst := range tests {

		refType := reflect.TypeOf(tst.input)

		switch refType.String() {
		case "*check.Check":
			// cast type to *check.Check, for calling Valid()
			v := tst.input.(*check.Check)
			// call valid with a check param
			got, err := rec.IsValid(v)
			if err != nil {
				fmt.Println("call valid error:", err)
				continue
			}
			if got != tst.want {
				t.Errorf("test faild,got:%v, want:%v", got, tst.want)
			}
		case "*check.Paycheck":
			// cast type to *check.Paycheck, for calling Valid()
			v := tst.input.(*check.Paycheck)
			// call valid with a paycheck param
			got, err := rec.IsValid(v)
			if err != nil {
				fmt.Println("call valid error:", err)
				continue
			}
			if got != tst.want {
				t.Errorf("test faild,got:%v, want:%v", got, tst.want)
			}
		default:
			fmt.Println("test data error, data type not check or payckeck")
		}
	}
}
