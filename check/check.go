package check

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Check struct {
	Value        *big.Int
	TokenAddr    common.Address
	Nonce        *big.Int
	FromAddr     common.Address
	ToAddr       common.Address
	OpAddr       common.Address
	ContractAddr common.Address
	CheckSig     []byte
}

// Sign check
func (chk *Check) Sign(sk string) error {

	hash := chk.Hash()

	//
	priKeyECDSA, err := crypto.HexToECDSA(sk)
	if err != nil {
		log.Print(err)
		return err
	}

	// sign to bytes
	sigByte, err := crypto.Sign(hash, priKeyECDSA)
	if err != nil {
		log.Print(err)
		return err
	}

	chk.CheckSig = sigByte

	return nil
}

// verify signature of a check
func (chk *Check) Verify() (bool, error) {

	hash := chk.Hash()

	// signature to public key
	pubKeyECDSA, err := crypto.SigToPub(hash, chk.CheckSig)
	if err != nil {
		log.Println("SigToPub err:", err)
		return false, err
	}

	// pub key to common.address
	recAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	ok := recAddr == chk.OpAddr

	return ok, nil
}

// calc hash of check, used to sign check and verify
func (chk *Check) Hash() []byte {

	valuePad32 := common.LeftPadBytes(chk.Value.Bytes(), 32)
	noncePad32 := common.LeftPadBytes(chk.Nonce.Bytes(), 32)

	// tokenBytes, _ := hex.DecodeString(Check.TokenAddr)
	// fromBytes, _ := hex.DecodeString(Check.From)
	// toBytes, _ := hex.DecodeString(Check.To)
	// operatorBytes, _ := hex.DecodeString(Check.OperatorAddr)
	// contractBytes, _ := hex.DecodeString(Check.ContractAddr)
	tokenBytes := chk.TokenAddr.Bytes()
	fromBytes := chk.FromAddr.Bytes()
	toBytes := chk.ToAddr.Bytes()
	operatorBytes := chk.OpAddr.Bytes()
	contractBytes := chk.ContractAddr.Bytes()

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

// Paycheck is an auto generated low-level Go binding around an user-defined struct.
type Paycheck struct {
	Check       Check
	PayValue    *big.Int
	PaycheckSig []byte
}

// Sign paycheck by user's sk
func (pchk *Paycheck) Sign(sk string) error {

	hash := pchk.Hash()

	//
	priKeyECDSA, err := crypto.HexToECDSA(sk)
	if err != nil {
		log.Print(err)
		return err
	}

	// sign to bytes
	sigByte, err := crypto.Sign(hash, priKeyECDSA)
	if err != nil {
		log.Print(err)
		return err
	}

	pchk.PaycheckSig = sigByte

	return nil
}

// verify signature of paycheck
func (pchk *Paycheck) Verify() (bool, error) {
	hash := pchk.Hash()

	// signature to public key
	pubKeyECDSA, err := crypto.SigToPub(hash, pchk.PaycheckSig)
	if err != nil {
		log.Println("SigToPub err:", err)
		return false, err
	}

	// pub key to common.address
	singerAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	ok := singerAddr == pchk.Check.FromAddr

	return ok, nil
}

// calc hash of check, used to sign check and verify
func (pchk *Paycheck) Hash() []byte {

	valuePad32 := common.LeftPadBytes(pchk.Check.Value.Bytes(), 32)
	noncePad32 := common.LeftPadBytes(pchk.Check.Nonce.Bytes(), 32)
	payvaluePad32 := common.LeftPadBytes(pchk.PayValue.Bytes(), 32)

	tokenBytes := pchk.Check.TokenAddr.Bytes()
	fromBytes := pchk.Check.FromAddr.Bytes()
	toBytes := pchk.Check.ToAddr.Bytes()
	operatorBytes := pchk.Check.OpAddr.Bytes()
	contractBytes := pchk.Check.ContractAddr.Bytes()

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
