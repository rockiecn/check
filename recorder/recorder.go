package recorder

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
)

type Key struct {
	Operator common.Address
	Provider common.Address
	Nonce    uint64
}

type Entry map[*Key]interface{}

// recorder for paycheck
type Recorder struct {
	Entrys Entry
}

func New() *Recorder {
	rec := new(Recorder)
	return rec
}

func (rec *Recorder) Record(entry interface{}) error {

	if c, ok := entry.(*check.Check); ok {
		key := new(Key)
		key.Operator = c.OpAddr
		key.Provider = c.ToAddr
		key.Nonce = c.Nonce
		if rec.Entrys == nil {
			rec.Entrys = make(Entry)
		}
		rec.Entrys[key] = c
		return nil
	}

	if pc, ok := entry.(*check.Paycheck); ok {
		key := new(Key)
		key.Operator = pc.Check.OpAddr
		key.Provider = pc.Check.ToAddr
		key.Nonce = pc.Check.Nonce
		if rec.Entrys == nil {
			rec.Entrys = make(Entry)
		}
		rec.Entrys[key] = pc
		return nil
	}

	return errors.New("entry type error, must be check or paycheck")

}
