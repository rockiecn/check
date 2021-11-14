package pb

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestPb(t *testing.T) {
	pbData := new(Check)
	pbData.Value = 123
	pbData.TokenAddress = "0xb213d01542d129806d664248a380db8b12059061"
	pbData.Nonce = 0
	pbData.From = "0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"
	pbData.To = "0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"
	pbData.OperatorAddress = "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"
	pbData.ContractAddress = "0x9e0153496067c20943724b79515472195a7aedaa"
	pbData.CheckSig = []byte("8001d4e51c97d68fe933f33c02e6383c21e62bf323334daca9fe9e6dc5da1106691d60e3c747c229e0b228af99a74dccc457d6e1215e683ba38da340a374d3f000")

	fmt.Println("data before marshal:", pbData)
	marshaled, err := proto.Marshal(pbData)
	if err != nil {
		t.Fatal("marshal failed")
	}

	pbData2 := new(Check)
	err = proto.Unmarshal(marshaled, pbData2)
	if err != nil {
		t.Fatal("unmarshal failed")
	}
	fmt.Println("data after marshal:", pbData2)
}
