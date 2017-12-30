package geo

import (
	"github.com/golang/geo/s2"
	. "github.com/noypi/spatial/common"
)

type EnumGeo interface {
	Close()
	Next() (item *RegionItem, has bool)
}

type _EnumGeo struct {
	Enum
}

func (this *_EnumGeo) Next() (item *RegionItem, has bool) {
	v, has := this.Enum.Next()
	if !has {
		return
	}
	reg := v.Value().(*Region)
	item = &RegionItem{Region: reg, Item: v}
	return
}

type _EnumWithinLoop struct {
	LoopBound *s2.Loop
	Enum
}

func (this *_EnumWithinLoop) Next() (item *RegionItem, has bool) {
	for {
		v, has := this.Enum.Next()
		if !has {
			return nil, false
		}
		reg := v.Value().(*Region)

		if this.LoopBound.Contains(reg.Loop) {
			return &RegionItem{Region: reg, Item: v}, true
		}
	}

	return nil, false
}

type _EnumWithinRegion struct {
	reg s2.Region
	Enum
}

func (this *_EnumWithinRegion) Next() (item *RegionItem, has bool) {
OUTER_LOOP:
	for {
		v, has := this.Enum.Next()
		if !has {
			return nil, false
		}
		reg := v.Value().(*Region)
		pts := reg.Loop.Vertices()
		for _, p := range pts {
			if !this.reg.ContainsPoint(p) {
				continue OUTER_LOOP
			}
		}

		return &RegionItem{Region: reg, Item: v}, true
	}

	return nil, false
}
