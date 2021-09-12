package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rockiecn/check/check"
)

// calc hash of check, used to sign check and verify
func CheckHash(Check *check.Check) []byte {

	valuePad32 := common.LeftPadBytes(Check.Value.Bytes(), 32)
	noncePad32 := common.LeftPadBytes(Check.Nonce.Bytes(), 32)

	tokenBytes, _ := hex.DecodeString(Check.TokenAddr)
	fromBytes, _ := hex.DecodeString(Check.From)
	toBytes, _ := hex.DecodeString(Check.To)
	operatorBytes, _ := hex.DecodeString(Check.OperatorAddr)
	contractBytes, _ := hex.DecodeString(Check.ContractAddr)

	// calc hash
	hash := crypto.Keccak256(
		valuePad32,
		tokenBytes,
		noncePad32,
		fromBytes,
		toBytes,
		operatorBytes,
		contractBytes,
	)

	return hash
}

// calc hash of check, used to sign check and verify
func PayCheckHash(PayCheck *check.PayCheck) []byte {

	valuePad32 := common.LeftPadBytes(PayCheck.Check.Value.Bytes(), 32)
	noncePad32 := common.LeftPadBytes(PayCheck.Check.Nonce.Bytes(), 32)
	payvaluePad32 := common.LeftPadBytes(PayCheck.PayValue.Bytes(), 32)

	tokenBytes, _ := hex.DecodeString(PayCheck.Check.TokenAddr)
	fromBytes, _ := hex.DecodeString(PayCheck.Check.From)
	toBytes, _ := hex.DecodeString(PayCheck.Check.To)
	operatorBytes, _ := hex.DecodeString(PayCheck.Check.OperatorAddr)
	contractBytes, _ := hex.DecodeString(PayCheck.Check.ContractAddr)

	// calc hash
	hash := crypto.Keccak256(
		valuePad32,
		tokenBytes,
		noncePad32,
		fromBytes,
		toBytes,
		operatorBytes,
		contractBytes,
		payvaluePad32,
	)

	return hash
}

// get address from private key
func KeyToAddr(sk string) string {
	skECDSA, err := crypto.HexToECDSA(sk)
	if err != nil {
		fmt.Println("hex to ecdsa failed:", err)
	}

	pubKey := skECDSA.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	addr := crypto.PubkeyToAddress(*pubKeyECDSA)

	big.NewInt(1231231231231231231)
	return addr.String()
}
