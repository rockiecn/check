package recorder

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
)

func Test(t *testing.T) {

	sig1, _ := hex.DecodeString("18dbfea51d279adb0ae16dff88b0dd341cd06838970779cfe056a2f7d359f4c431fb61168e261aed805b285e523e0092b5d46a2c51e6b8b093c5a151abf24b0100")
	sig2, _ := hex.DecodeString("142673fc986aa6d456fba6df9ce82cfdc92cf9af7ea169c41787cf59fcaba9b2721b0b72a121dadd072b1f33743cafb21920e0bea2fe6b435b22ffb715c49d6001")
	sig3, _ := hex.DecodeString("18dbfea51d279adb0ae16dff88b0dd341cd06838970779cfe056a2f7d359f4c431fb61168e261aed805b285e523e0092b5d46a2c51e6b8b093c5a151abf24b0100")
	big100, _ := new(big.Int).SetString("100", 10)
	big200, _ := new(big.Int).SetString("200", 10)
	big300, _ := new(big.Int).SetString("300", 10)
	big1, _ := new(big.Int).SetString("1", 10)
	big2, _ := new(big.Int).SetString("2", 10)
	big3, _ := new(big.Int).SetString("3", 10)

	var paychecks = []check.Paycheck{
		{
			Check: check.Check{
				Value:        big100,
				TokenAddr:    common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				Nonce:        big1,
				FromAddr:     common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				ToAddr:       common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				OpAddr:       common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				ContractAddr: common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				CheckSig:     sig1,
			},
			PayValue:    big1,
			PaycheckSig: sig1,
		},
		{
			Check: check.Check{
				Value:        big200,
				TokenAddr:    common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				Nonce:        big2,
				FromAddr:     common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ToAddr:       common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				OpAddr:       common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				ContractAddr: common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				CheckSig:     sig2,
			},
			PayValue:    big2,
			PaycheckSig: sig2,
		},
		{
			Check: check.Check{
				Value:        big300,
				TokenAddr:    common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				Nonce:        big3,
				FromAddr:     common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				ToAddr:       common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				OpAddr:       common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				ContractAddr: common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
				CheckSig:     sig3,
			},
			PayValue:    big3,
			PaycheckSig: sig3,
		},
	}

	rc := NewPRecorder()

	// write payckecks into recorder
	for _, pc := range paychecks {
		rc.Record(&pc)
	}

	// show all paychecks
	rc.List()
}
