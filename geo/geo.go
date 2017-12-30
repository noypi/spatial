package geo

import (
	"github.com/noypi/spatial/db"
)

type SpatialGeo struct {
	db *spatial.Spatial1D
}

func NewGeo() *SpatialGeo {
	o := new(SpatialGeo)
	o.db = spatial.New1D()
	return o
}
