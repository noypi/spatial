package geo

import (
	"github.com/noypi/spatial/db"
)

func (this *SpatialGeo) AsDebug() *spatial.Debugger {
	return this.db.AsDebug()
}
