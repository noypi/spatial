package geo

import (
	"github.com/noypi/spatial/db"
)

type SpatialGeo struct {
	db *spatial.Spatial1D
}

func NewGeo(opts ...spatial.Options) (o *SpatialGeo, err error) {
	o = new(SpatialGeo)
	o.db, err = spatial.New1D(opts...)
	return
}

func (this SpatialGeo) DbPath() string {
	return this.db.DbPath()
}
