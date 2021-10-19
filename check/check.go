package check

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	comn "github.com/rockiecn/check/common"
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

	return nil
}

// verify signature of a check
func (chk *Check) Verify() (bool, error) {

	hash := chk.Serialize()

	// signature to public key
	pubKeyECDSA, err := crypto.SigToPub(hash, chk.CheckSig)
	if err != nil {
		return false, errors.New("SigToPub err")
	}

	// pub key to common.address
	recAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	ok := recAddr == chk.OpAddr

	return ok, nil
}

// calc hash of check, used to sign check and verify
func (chk *Check) Serialize() []byte {

	valuePad32 := common.LeftPadBytes(chk.Value.Bytes(), 32)
	tokenBytes := chk.TokenAddr.Bytes()
	noncePad8 := common.LeftPadBytes(comn.Uint64ToBytes(chk.Nonce), 8)
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
	Check       Check
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

	return nil
}

// verify signature of paycheck
func (pchk *Paycheck) Verify() (bool, error) {
	hash := pchk.Serialize()

	// signature to public key
	pubKeyECDSA, err := crypto.SigToPub(hash, pchk.PaycheckSig)
	if err != nil {
		return false, errors.New("SigToPub err")
	}

	// pub key to common.address
	signerAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	ok := signerAddr == pchk.Check.FromAddr

	return ok, nil
}

// calc hash of check, used to sign check and verify
func (pchk *Paycheck) Serialize() []byte {

	valuePad32 := common.LeftPadBytes(pchk.Check.Value.Bytes(), 32)
	noncePad8 := common.LeftPadBytes(comn.Uint64ToBytes(pchk.Check.Nonce), 8)
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

type BatchCheck struct {
	cheque_batch_to    common.Address // 存储节点号
	cheque_batch_value uint           // 聚合后的支票面额
	min_nonce          uint           // 聚合的nonce最小值
	max_nonce          uint           // 聚合的nonce最大值

	download_cheque_batch_sign []byte // signature of operator
}
