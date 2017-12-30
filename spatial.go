package spatial

import (
	. "github.com/noypi/spatial/common"
	skv "github.com/noypi/spatial/flatv1"
)

type Dimension int

type Type int

const (
	KVType Type = iota
	AnyType
)

func New1D(t Type) Spatial1D {
	if t == KVType {
		return skv.New1D()
	}

	panic("not implemented")
	return nil
}

func New2D(t Type) Spatial2D {
	if t == KVType {
		return skv.New2D()
	}

	panic("not implemented")
	return nil
}

func New3D(t Type) Spatial3D {
	if t == KVType {
		return skv.New3D()
	}

	panic("not implemented")
	return nil
}

type Spatial1D interface {
	AddRange(r Range, v interface{}) error
	Contains(x float64) Enum
	ContainsRange(min, max float64) Enum
	WithinRange(min, max float64) Enum
}

type Spatial2D interface {
	AddRange(x, y Range, v interface{}) error
	Contains(x, y float64) Enum
	ContainsRange(x, y Range) Enum
	WithinRange(x, y Range) Enum
}

type Spatial3D interface {
	AddRange(x, y, z Range, v interface{}) error
	Contains(x, y, z float64) Enum
	ContainsRange(x, y, z Range) Enum
	WithinRange(x, y, z Range) Enum
}
