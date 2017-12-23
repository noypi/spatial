package spatial

import (
	"github.com/noypi/mapk"
	. "github.com/noypi/spatial/common"
)

type Spatial1D struct {
	m, mrev mapk.IMap
}

func New1D() *Spatial1D {
	o := new(Spatial1D)
	o.m = mapk.MapGTreap(CompareFunc1D)
	o.mrev = mapk.MapGTreap(CompareReverseFunc1D)
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
