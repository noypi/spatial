package geo

import (
	"math"

	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
)

type Distance float64

const (
	EarthInchesCircumference = 1577727360
	EarthMetersCircumference = 40074274.944
	EarthKmCircumference     = 40074.274944

	TwoPi = (2 * math.Pi)

	EarthInchRadius  = EarthInchesCircumference / TwoPi
	EarthMeterRadius = EarthMetersCircumference / TwoPi
	EarthKmRadius    = EarthKmCircumference / TwoPi

	Meters Distance = 1.0
	Inches Distance = 39.3701
	Km     Distance = 1000.0
)

func AngleFromDistance(d Distance) s1.Angle {
	return s1.Angle(d / EarthMeterRadius)
}

func CapAroundPoint(p s2.Point, d Distance) s2.Cap {
	angle := AngleFromDistance(d)
	chordangle := s1.ChordAngleFromAngle(angle)
	return s2.CapFromCenterChordAngle(p, chordangle)
}

func (d Distance) Km() float64 {
	return float64(d / Km)
}

func (d Distance) Meters() float64 {
	return float64(d / Meters)
}

func (d Distance) Inches() float64 {
	return float64(d / Inches)
}
