package operator

import (
	"fmt"
	"testing"
)

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

// test new operator and deploy contract
func TestNew(t *testing.T) {
	_, err := New("503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb", "0xb213d01542d129806d664248a380db8b12059061")
	if err != nil {
		fmt.Println("new operator failed:", err)
	}
}
