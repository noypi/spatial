package spatial

import (
	"github.com/noypi/kv"
	. "github.com/noypi/spatial/common"
)

type Spatial1D struct {
	store     kv.KVStore
	extinfo   kv.KVStore
	xyzOffset int
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
	return nil
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
