package osm

import (
	"github.com/golang/geo/s2"
	"github.com/noypi/spatial/geo"
)

type Item struct {
	LatLngs []s2.LatLng
}

type GeoInfo struct {
	LatLngs []s2.LatLng
	Tags    map[string]string
}

func init() {
	geo.RegisterType(Item{})
}
