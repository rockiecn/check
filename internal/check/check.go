package check

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rockiecn/check/internal/utils"
)

type Check struct {
	Value        *big.Int
	TokenAddr    common.Address
	Nonce        uint64
	FromAddr     common.Address
	ToAddr       common.Address
	OpAddr       common.Address
	ContractAddr common.Address
	CheckSig     []byte
}

type ICheck interface {
	Sign(sk string) error
	Verify() (bool, error)
	Serialize() []byte
	GetNonce() (uint64, error)
}

// Sign check
func (chk *Check) Sign(sk string) error {

	hash := chk.Serialize()

	//fmt.Printf("sign hash:%x\n", hash)
	//
	priKeyECDSA, err := crypto.HexToECDSA(sk)
	if err != nil {
		return errors.New("hex to ecdsa failed when sign check")
	}

	// sign to bytes
	sigByte, err := crypto.Sign(hash, priKeyECDSA)
	if err != nil {
		return errors.New("sign failed when sign check")
	}

	chk.CheckSig = sigByte

	//fmt.Printf("sign chk sig:%x\n", sigByte)

	return nil
}

// verify signature of a check
func (chk *Check) Verify() bool {

	hash := chk.Serialize()

	// signature to public key
	pubKeyECDSA, err := crypto.SigToPub(hash, chk.CheckSig)
	if err != nil {
		return false
	}

	// pub key to common.address
	recAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	ok := recAddr == chk.OpAddr

	return ok
}

// calc hash of check, used to sign check and verify
func (chk *Check) Serialize() []byte {

	valuePad32 := common.LeftPadBytes(chk.Value.Bytes(), 32)
	tokenBytes := chk.TokenAddr.Bytes()
	noncePad8 := common.LeftPadBytes(utils.Uint64ToByte(chk.Nonce), 8)
	fromBytes := chk.FromAddr.Bytes()
	toBytes := chk.ToAddr.Bytes()
	operatorBytes := chk.OpAddr.Bytes()
	contractBytes := chk.ContractAddr.Bytes()

	// calc hash
	hash := crypto.Keccak256(
		valuePad32,
		tokenBytes,
		noncePad8,
		fromBytes,
		toBytes,
		operatorBytes,
		contractBytes,
	)

	return hash
}

// Paycheck is an auto generated low-level Go binding around an user-defined struct.
type Paycheck struct {
	Check       *Check
	PayValue    *big.Int
	PaycheckSig []byte
}

// Sign paycheck by user's sk
func (pchk *Paycheck) Sign(sk string) error {

	hash := pchk.Serialize()

	//
	priKeyECDSA, err := crypto.HexToECDSA(sk)
	if err != nil {
		return errors.New("hex to ecdsa failed when sign paycheck")
	}

	// sign to bytes
	sigByte, err := crypto.Sign(hash, priKeyECDSA)
	if err != nil {
		return errors.New("sign paycheck error")
	}

	pchk.PaycheckSig = sigByte

	//fmt.Printf("sign pc sig:%x\n", sigByte)

	return nil
}

// verify signature of paycheck
func (pchk *Paycheck) Verify() bool {
	hash := pchk.Serialize()

	// signature to public key
	pubKeyECDSA, err := crypto.SigToPub(hash, pchk.PaycheckSig)
	if err != nil {
		return false
	}

	// pub key to common.address
	signerAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	ok := signerAddr == pchk.Check.FromAddr

	return ok
}

// calc hash of check, used to sign check and verify
func (pchk *Paycheck) Serialize() []byte {

	valuePad32 := common.LeftPadBytes(pchk.Check.Value.Bytes(), 32)
	noncePad8 := common.LeftPadBytes(utils.Uint64ToByte(pchk.Check.Nonce), 8)
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
		noncePad8,
		fromBytes,
		toBytes,
		operatorBytes,
		contractBytes,
		payvaluePad32,
	)

	return hash
}

// equal
func (pchk *Paycheck) Equal(p2 *Paycheck) (bool, error) {
	if pchk.Check.Value.Cmp(p2.Check.Value) != 0 {
		return false, errors.New("value not equal")
	}
	if pchk.Check.TokenAddr != p2.Check.TokenAddr {
		return false, errors.New("token not equal")
	}
	if pchk.Check.Nonce != p2.Check.Nonce {
		return false, errors.New("nonce not equal")
	}
	if pchk.Check.FromAddr != p2.Check.FromAddr {
		return false, errors.New("from not equal")
	}
	if pchk.Check.ToAddr != p2.Check.ToAddr {
		return false, errors.New("to not equal")
	}
	if pchk.Check.OpAddr != p2.Check.OpAddr {
		return false, errors.New("op not equal")
	}
	if pchk.Check.ContractAddr != p2.Check.ContractAddr {
		return false, errors.New("contrAddr not equal")
	}
	if !bytes.Equal(pchk.Check.CheckSig, p2.Check.CheckSig) {
		return false, errors.New("check sig not equal")
	}
	if pchk.PayValue.String() != p2.PayValue.String() {
		return false, errors.New("pay value not equal")
	}
	if !bytes.Equal(pchk.PaycheckSig, p2.PaycheckSig) {
		return false, errors.New("paycheck sig not equal")
	}

	return true, nil
}

type BatchCheck struct {
	OpAddr     common.Address // operator address
	ToAddr     common.Address // 存储节点号
	BatchValue *big.Int       // 聚合后的支票面额
	MinNonce   uint64         // 聚合的nonce最小值
	MaxNonce   uint64         // 聚合的nonce最大值

	BatchSig []byte // signature of operator
}

// Sign BatchCheck by opertator's sk
func (bc *BatchCheck) Sign(sk string) error {

	hash := bc.Serialize()

	//
	priKeyECDSA, err := crypto.HexToECDSA(sk)
	if err != nil {
		return errors.New("hex to ecdsa failed when sign paycheck")
	}

	// sign to bytes
	sigByte, err := crypto.Sign(hash, priKeyECDSA)
	if err != nil {
		return errors.New("sign paycheck error")
	}

	bc.BatchSig = sigByte

	return nil
}

// verify signature of paycheck
func (bc *BatchCheck) Verify() (bool, error) {
	hash := bc.Serialize()

	// signature to public key
	pubKeyECDSA, err := crypto.SigToPub(hash, bc.BatchSig)
	if err != nil {
		return false, errors.New("SigToPub err")
	}

	// pub key to common.address
	signerAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	ok := signerAddr == bc.OpAddr

	return ok, nil
}

// calc hash of BatchCheck, used to sign and verify
func (bc *BatchCheck) Serialize() []byte {

	opBytes := bc.OpAddr.Bytes()
	toBytes := bc.ToAddr.Bytes()
	valuePad32 := common.LeftPadBytes(bc.BatchValue.Bytes(), 32)
	minPad8 := common.LeftPadBytes(utils.Uint64ToByte(bc.MinNonce), 8)
	maxPad8 := common.LeftPadBytes(utils.Uint64ToByte(bc.MaxNonce), 8)

	// calc hash
	hash := crypto.Keccak256(
		opBytes,
		toBytes,
		valuePad32,
		minPad8,
		maxPad8,
	)

	return hash
}
