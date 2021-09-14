package user

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rockiecn/check/check"
	"github.com/rockiecn/check/utils"
)

type User struct {
	UserSK   string
	UserAddr string

	//
	History map[string]*check.PayCheck // keyHash -> key, paycheck, key: "operator:xxx, provider:xxx, nonce:xxx"
}

type IUser interface {
	VerifyCheck(check *check.Check, opAddr common.Address) (bool, error)
	GeneratePayCheck(check *check.Check, payValue *big.Int) (*check.PayCheck, error)
	// Sign paycheck by user's sk
	Sign(paycheck *check.PayCheck, skByte []byte) ([]byte, error)
	RecordPayCheck(check *check.PayCheck) error
}

func NewUser(sk string) (*User, error) {
	user := new(User)
	user.UserSK = sk
	user.UserAddr = utils.KeyToAddr(sk)

	return user, nil
}

// verify signature of a check, check should be signed by an operator
func (user *User) VerifyCheck(check *check.Check) (bool, error) {

	hash := utils.CheckHash(check)

	// signature to public key
	pubKeyECDSA, err := crypto.SigToPub(hash, check.CheckSig)
	if err != nil {
		log.Println("SigToPub err:", err)
		return false, err
	}

	// pub key to common.address
	recAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	ok := recAddr == common.HexToAddress(check.OperatorAddr)

	return ok, nil
}

// Sign paycheck by user's sk
func (user *User) Sign(paycheck *check.PayCheck) ([]byte, error) {

	hash := utils.PayCheckHash(paycheck)

	//
	priKeyECDSA, err := crypto.HexToECDSA(user.UserSK)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// sign to bytes
	sigByte, err := crypto.Sign(hash, priKeyECDSA)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return sigByte, nil
}

// generate paycheck based on check, sig of paycheck is updated
func (user *User) GeneratePayCheck(chk *check.Check, payValue *big.Int) (*check.PayCheck, error) {
	pchk := new(check.PayCheck)
	pchk.Check = chk
	pchk.PayValue = payValue

	sig, err := user.Sign(pchk)
	if err != nil {
		fmt.Println("user sign paycheck error:", err)
		return nil, err
	}
	pchk.PayCheckSig = sig

	return pchk, nil
}
