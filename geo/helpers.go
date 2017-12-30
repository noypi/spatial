package geo

import (
	"encoding/gob"

	"github.com/golang/geo/s2"
)

func RegisterType(v interface{}) {
	gob.Register(v)
}

func LatLng(lat, lng float64) s2.LatLng {
	return s2.LatLngFromDegrees(lat, lng)
}

func toCellID(lat, lng float64) s2.CellID {
	latlng := s2.LatLngFromDegrees(lat, lng)
	return s2.CellIDFromLatLng(latlng)
}
