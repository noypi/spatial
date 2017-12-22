package spatial

import (
	"github.com/noypi/kv"
	"github.com/noypi/kv/gtreap"
)

type Spatial1D struct {
	store kv.KVStore
}

func New1D() *Spatial1D {
	o := new(Spatial1D)
	o.store = gtreap.GetDefault()
	return o
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