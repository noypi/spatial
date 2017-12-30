package spatial

import (
	"fmt"

	"github.com/noypi/kv"
	. "github.com/noypi/spatial/common"
)

type _Item struct {
	V       interface{}
	Keys    []_ID
	err     error
	store   kv.KVStore
	kvbatch kv.KVBatch
	enum    *_Enum

	currKeyOffset int
}

func NewItem(id []byte, v interface{}, rs ...Range) *_Item {
	o := &_Item{V: v}
	o.Keys = make([]_ID, len(rs))
	for i, r := range rs {
		o.setItem(i, r, id)
	}

	if 0 == len(o.Keys) {
		panic("keys should not be zero.")
	}

	return o
}

func (this *_Item) setItem(keyOffset int, r Range, id []byte) {
	newid := NewID(id)
	newid.SetPrefix(cPrefixRange)
	r.MaximizeIfZeroMax()
	newid.SetLeftUint64(r.Min)
	newid.SetRightUint64(r.Max)
	this.Keys[keyOffset] = newid
}

func (this _Item) Clone() *_Item {
	o := new(_Item)
	o.V = this.V
	o.Keys = this.Keys
	o.store = this.store
	o.enum = this.enum
	if 0 == len(o.Keys) {
		panic("keys should not be zero.")
	}
	return o
}

func (this _Item) Error() error {
	return this.err
}

func (this _Item) Range(n int) (r Range) {
	id := this.Keys[n]
	left := id.LeftUint64()
	right := id.RightUint64()
	switch id.Prefix() {
	case cPrefixRange:
		r.Min = left
		r.Max = right
	case cPrefixRangeReverse:
		r.Min = right
		r.Max = left
	}
	return
}

func (this _Item) Value() interface{} {
	return this.V
}

func (this _Item) ID() string {
	return fmt.Sprintf("%x", this.Keys[this.currKeyOffset])
}

func (this *_Item) Set(v interface{}) error {
	this.ensurebatch()

	this.V = v
	vbb, err := GobSerialize(this)
	if nil != err {
		return err
	}
	for _, id := range this.Keys {
		setItemToBatch(this.kvbatch, id, vbb)
	}
	return nil
}

func (this *_Item) Delete() {
	this.ensurebatch()
	for _, id := range this.Keys {
		deleteItemToBatch(this.kvbatch, id)
	}
}

func (this *_Item) ensurebatch() {
	if nil != this.kvbatch {
		wrtr, err := this.store.Writer()
		if nil != err {
			panic(err)
		}
		this.kvbatch = wrtr.NewBatch()
		this.enum.addtocommit(func() {
			wrtr.ExecuteBatch(this.kvbatch)
		})
	}
}
