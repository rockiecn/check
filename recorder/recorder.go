package recorder

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
)

type Key struct {
	Operator common.Address
	Provider common.Address
	Nonce    uint64
}

// recorder for paycheck
type Recorder struct {
	Entrys map[Key]interface{}
}

func New() *Recorder {

	rec := &Recorder{
		Entrys: make(map[Key]interface{}),
	}
	return rec
}

// put a check/paycheck into Entrys
func (rec *Recorder) Record(entry interface{}) error {

	t := reflect.TypeOf(entry)

	switch t.String() {
	case "*check.Check":
		if c, ok := entry.(*check.Check); ok {
			key := Key{
				Operator: c.OpAddr,
				Provider: c.ToAddr,
				Nonce:    c.Nonce,
			}
			rec.Entrys[key] = c
			return nil
		}
		return errors.New("check type assertion failed")
	case "*check.Paycheck":
		if pc, ok := entry.(*check.Paycheck); ok {
			key := Key{
				Operator: pc.Check.OpAddr,
				Provider: pc.Check.ToAddr,
				Nonce:    pc.Check.Nonce,
			}
			rec.Entrys[key] = pc
			return nil
		}
		return errors.New("paycheck type assertion failed")
	default:
		return errors.New("type of value must be Check or Paycheck")
	}
}

// check if a check/paycheck exist. by key (to, nonce)
func (rec *Recorder) Exist(entry interface{}) (bool, error) {

	t := reflect.TypeOf(entry)

	switch t.String() {
	case "*check.Check":
		c := entry.(*check.Check)
		for k := range rec.Entrys {
			if k.Nonce == c.Nonce && k.Provider == c.ToAddr {
				return true, nil
			}
		}
		return false, nil
	case "*check.Paycheck":
		pc := entry.(*check.Paycheck)
		for k := range rec.Entrys {
			if k.Nonce == pc.Check.Nonce && k.Provider == pc.Check.ToAddr {
				return true, nil
			}
		}
		return false, nil
	default:
		return false, errors.New("invalid type, type must be check/paycheck")
	}
}

// show info of a check/paycheck
func (rec *Recorder) Show(entry interface{}) error {

	t := reflect.TypeOf(entry)

	switch t.String() {
	case "*check.Check":
		if c, ok := entry.(*check.Check); ok {
			fmt.Println("value:", c.Value)
			return nil
		}
		return errors.New("check type assertion failed")
	case "*check.Paycheck":
		if c, ok := entry.(*check.Paycheck); ok {
			fmt.Println("payvalue:", c.PayValue)
			return nil
		}
		return errors.New("paycheck type assertion failed")
	default:
		return errors.New("type of value must be Check or Paycheck")
	}

}

// show all values of Entrys
func (rec *Recorder) ShowAll() error {
	for k, v := range rec.Entrys {
		fmt.Println("provider:", k.Provider)
		rec.Show(v)
	}
	return nil
}
