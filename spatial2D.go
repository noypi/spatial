package spatial

import (
	"github.com/noypi/mapk"
)

type Spatial2D struct {
	m mapk.IMap
}

func New2D() *Spatial1D {
	o := new(Spatial1D)
	o.m = mapk.MapGTreap(CompareFunc1D)
	return o
}

func CompareFunc2D(a, b interface{}) int {
	a1 := a.(*_item2D)
	b1 := b.(*_item2D)
	return a1.Compare(b1)
}

type _item2D struct {
	x, y Range
}

func (a _item2D) Compare(b *_item2D) int {
	cmp := a.x.Compare(b.x)
	if 0 == cmp {
		cmp = a.y.Compare(b.y)
	}
	return cmp
}
