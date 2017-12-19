package spatial

import (
	"github.com/noypi/mapk"
)

type Spatial1D struct {
	m mapk.IMap
}

func New1D() *Spatial1D {
	o := new(Spatial1D)
	o.m = mapk.MapGTreap(CompareFunc1D)
	return o
}

func CompareFunc1D(a, b interface{}) int {
	a1 := a.(*Range)
	b1 := b.(*Range)
	return a1.Compare(*b1)
}
