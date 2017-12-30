package geo

import (
	"github.com/golang/geo/s2"
)

type RangeGeo struct {
	Min, Max s2.LatLng
}

func (r RangeGeo) MinID() uint64 {
	return uint64(s2.CellIDFromLatLng(r.Min))
}

func (r RangeGeo) MaxID() uint64 {
	return uint64(s2.CellIDFromLatLng(r.Max))
}
