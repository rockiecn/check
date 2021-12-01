package user

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/utils"
)

var (
	pchk0 = &check.Paycheck{
		Check: &check.Check{
			CheckInfo: check.CheckInfo{
				Value:     utils.String2BigInt("100000000000000000000"),
				TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
				Nonce:     0,
				FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			},
			CheckSig: utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		PayValue:    utils.String2BigInt("1000000000000000000"),
		PaycheckSig: utils.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
	pchk1 = &check.Paycheck{
		Check: &check.Check{
			CheckInfo: check.CheckInfo{
				Value:     utils.String2BigInt("100000000000000000000"),
				TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
				Nonce:     1,
				FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			},
			CheckSig: utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		PayValue:    utils.String2BigInt("1000000000000000000"),
		PaycheckSig: utils.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
	pchk2 = &check.Paycheck{
		Check: &check.Check{
			CheckInfo: check.CheckInfo{
				Value:     utils.String2BigInt("100000000000000000000"),
				TokenAddr: common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
				Nonce:     2,
				FromAddr:  common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ToAddr:    common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				OpAddr:    common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				CtrAddr:   common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			},
			CheckSig: utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		PayValue:    utils.String2BigInt("1000000000000000000"),
		PaycheckSig: utils.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}
)

// store some paychecks into db
func TestStore(t *testing.T) {
	usrSk, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	usr, err := New(usrSk)
	if err != nil {
		t.Fatal(err)
	}
	Usr := usr.(*User)

	Usr.Store(pchk0)
	Usr.Store(pchk1)
	Usr.Store(pchk2)
}

// restore paychecks from db
func TestRestore(t *testing.T) {
	usrSk, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	usr, err := New(usrSk)
	if err != nil {
		t.Fatal(err)
	}
	Usr := usr.(*User)

	// restore and show
	err = Usr.Restore()
	if err != nil {
		t.Fatal(err)
	}
	Usr.ShowPool()
}
