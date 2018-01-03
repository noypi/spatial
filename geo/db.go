package geo

import (
	"fmt"

	"github.com/golang/geo/s2"
	. "github.com/noypi/spatial/common"
)

func (this *SpatialGeo) AddRegion(id []byte, value interface{}, locs []s2.LatLng) error {

	if 0 == len(locs) {
		return fmt.Errorf("AddRegion err: invalid params, zero length locs.")
	}

	var min, max int

	pts := make([]s2.Point, len(locs))
	ids := make([]s2.CellID, len(locs))
	for i, latlng := range locs {
		pts[i] = s2.PointFromLatLng(latlng)
		ids[i] = s2.CellIDFromLatLng(latlng)
		if ids[i] < ids[min] {
			min = i
		} else if ids[i] > ids[max] {
			max = i
		}
	}

	reg := &Region{Pts: pts, loop: s2.LoopFromPoints(pts), V: value}
	return this.db.Set(id, Range{uint64(ids[min]), uint64(ids[max])}, reg)
}

func (this *SpatialGeo) Set(id []byte, r RangeGeo, v *Region) error {
	return this.db.Set(id, Range{r.MinID(), r.MaxID()}, v)
}

func (this *SpatialGeo) Contains(lat, lng float64) EnumGeo {
	cellid := toCellID(lat, lng)
	enum := this.db.Contains(uint64(cellid))
	return &_EnumGeo{Enum: enum}
}

func (this *SpatialGeo) ContainsRange(r RangeGeo) EnumGeo {
	enum := this.db.ContainsRange(r.MinID(), r.MaxID())
	return &_EnumGeo{Enum: enum}
}

func (this *SpatialGeo) Around(lat, lng float64, d Distance) EnumGeo {
	rc := s2.RegionCoverer{MaxLevel: 1, MaxCells: 20}
	latlng := s2.LatLngFromDegrees(lat, lng)
	p := s2.PointFromLatLng(latlng)
	r := s2.Region(CapAroundPoint(p, d))
	covering := rc.Covering(r)

	var min, max uint64
	for _, c := range covering {
		a := uint64(c.RangeMin())
		b := uint64(c.RangeMax())
		if a < min {
			min = a
		}
		if b > max {
			max = b
		}
	}

	enum := this.db.WithinRange(min, max)

	return &_EnumWithinRegion{reg: r, Enum: enum}
}

func (this *SpatialGeo) Within(item *RegionItem) EnumGeo {
	reg := item.Region
	rect := reg.Loop().RectBound()
	rangeGeo := RangeGeo{Min: rect.Lo(), Max: rect.Hi()}

	enum := this.db.WithinRange(rangeGeo.MinID(), rangeGeo.MaxID())
	return &_EnumWithinLoop{LoopBound: reg.Loop(), Enum: enum}
}
