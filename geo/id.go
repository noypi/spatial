package geo

import (
	"encoding/binary"

	"github.com/golang/geo/s2"
)

type ID [8]byte

func NewID(latitude, longitude float64) ID {
	id := ID{}
	latlng := s2.LatLngFromDegrees(latitude, longitude)
	cellid := s2.CellIDFromLatLng(latlng)
	binary.BigEndian.PutUint64(id[:], uint64(cellid))
	return id
}
