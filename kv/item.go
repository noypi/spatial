package spatial

import (
	"github.com/noypi/kv"
	. "github.com/noypi/spatial/common"
	"github.com/rs/xid"
)

type _Item struct {
	V       interface{}
	Ranges  []Range
	err     error
	id      xid.ID
	store   kv.KVStore
	kvbatch kv.KVBatch
	kvkey   []byte
	enum    *_Enum
}

func (this _Item) Error() error {
	return this.err
}

func (this _Item) Range(n int) Range {
	return this.Ranges[n]
}

func (this _Item) Value() interface{} {
	return this.V
}

func (this _Item) ID() string {
	return this.id.String()
}

func (this *_Item) Set(v interface{}) error {
	this.ensurebatch()

	this.V = v
	vbb, id, err := serializev(v)
	if nil != err {
		return err
	}
	for _, r := range this.Ranges {
		setItemToBatch(this.kvbatch, r, id, vbb)
	}
	return nil
}

func (this *_Item) Delete() {
	this.ensurebatch()
	_, id := keyToRange(this.kvkey)
	for _, r := range this.Ranges {
		deleteItemToBatch(this.kvbatch, r, id)
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
