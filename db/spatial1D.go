package spatial

import (
	"sync"

	"github.com/noypi/kv"
	. "github.com/noypi/spatial/common"
)

type Spatial1D struct {
	store     kv.KVStore
	extinfo   kv.KVStore
	extbatch  kv.KVBatch
	extwrtr   kv.KVWriter
	xyzOffset int
	dbpath    string

	syncExtBatch   sync.Mutex
	extBatchCnt    uint
	extBatchMaxCnt uint
}

func New1D(opts ...Options) (o *Spatial1D, err error) {
	mOpts, err := ParseOpts(opts)
	if nil != err {
		return nil, err
	}

	o = new(Spatial1D)
	if err = o.initStore(mOpts); nil != err {
		return
	}
	if err = o.initExtInfoStore(mOpts); nil != err {
		return
	}
	if o.dbpath, _ = mOpts[cOptKVDir].(string); 0 == len(o.dbpath) {
		o.dbpath = "."
	}

	return
}

func (this *Spatial1D) initStore(mOpts map[int]interface{}) error {
	kvname, _ := mOpts[cOptKVName]

	var kvopts []kv.Option
	if v, has := mOpts[cOptKVOptions]; has && nil != v {
		kvopts = v.([]kv.Option)
	}

	store, err := kv.New(kvname, kvopts...)
	if nil != err {
		return err
	}

	this.store = store
	return nil
}

func (this *Spatial1D) initExtInfoStore(mOpts map[int]interface{}) error {

	if b, has := mOpts[cOptEnableExtInfo]; !has || !b.(bool) {
		return nil
	}

	kvname, _ := mOpts[cOptKVName]
	var kvopts []kv.Option

	if fpath, has := mOpts[cOptExtInfoFilePath]; has {
		kvopts = []kv.Option{kv.OptFilePath{fpath.(string)}}
	}

	store, err := kv.New(kvname, kvopts...)
	if nil != err {
		return err
	}

	this.extinfo = store
	this.extBatchMaxCnt = 1000
	return nil
}

func (this Spatial1D) DbPath() string {
	return this.dbpath
}

func CompareFunc1D(a, b interface{}) int {
	a1 := a.(*Range)
	b1 := b.(*Range)
	return a1.Compare(*b1)
}

func CompareReverseFunc1D(b, a interface{}) int {
	a1 := a.(*Range)
	b1 := b.(*Range)
	return a1.Compare(*b1)
}

func (this *Spatial1D) Close() {
	this.FlushExt()
	if nil != this.extwrtr {
		this.extbatch.Close()
		this.extbatch = nil
		this.extwrtr.Close()
		this.extwrtr = nil
	}
	this.store.Close()
	if nil != this.extinfo {
		this.extinfo.Close()
		this.extinfo = nil
	}
	this.store = nil
}
