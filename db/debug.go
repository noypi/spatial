package spatial

import (
	"github.com/noypi/kv"
)

type Debugger struct {
	*Spatial1D
}

func (this *Spatial1D) AsDebug() *Debugger {
	return &Debugger{this}
}

func (this *Debugger) Store() kv.KVStore {
	return this.store
}

func (this *Debugger) DumpRange(begin, end []byte, cb func(k, v []byte)) {
	this.dumpRange(this.store, begin, end, cb)
}

func (this *Debugger) DumpRangeExtInfo(begin, end []byte, cb func(k, v []byte)) {
	this.dumpRange(this.extinfo, begin, end, cb)
}

func (this *Debugger) dumpRange(store kv.KVStore, begin, end []byte, cb func(k, v []byte)) {
	rdr, _ := store.Reader()
	iter := rdr.RangeIterator(begin, end)
	for iter.Valid() {
		cb(iter.Key(), iter.Value())
		iter.Next()
	}
}
