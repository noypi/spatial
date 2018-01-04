package osm

import (
	"log"
	"sync/atomic"

	"github.com/golang/geo/s2"
	"github.com/thomersch/gosmparse"
)

func (this *OsmParser) ReadRelation(r gosmparse.Relation) {
	if this.tCurrent != gosmparse.RelationType {
		this.syncAdd.Lock()
		if 0 < this.tmpKvBatchCnt {
			this.execBatch()
		}
		this.syncAdd.Unlock()
	}
	this.tCurrent = gosmparse.RelationType

	cnt := atomic.AddUint64(&this.RelationCnt, 1)
	if cnt < this.SkipRelations {
		return
	}
	if (cnt % 500000) == 0 {
		log.Println("relation cnt=", cnt)
	}

	latlngs := this.getS2Latlngs(r)
	if 0 == len(latlngs) {
		atomic.AddUint64(&this.MissedRelations, 1)
		return
	}

	if err := this.indexTags(r); nil != err {
		log.Println("ReadRelation err:", err)
	}

	id := idFromInt(r.ID)
	err := this.AddRegion(
		id,
		nil,
		latlngs,
	)
	if nil != err {
		log.Println("ReadRelation err:", err)
	}

	err = this.SetExtInfo(uint8(Relation), id, r.Tags)
	if nil != err {
		log.Println("ReadRelation err:", err)
	}
}

func (this *OsmParser) getS2Latlngs(r gosmparse.Relation) []s2.LatLng {
	var pts []s2.LatLng
	for _, rm := range r.Members {
		if gosmparse.NodeType == rm.Type {
			latlng, err := this.getLatlngFromTmpNode(rm.ID)
			if nil != err {
				log.Println("getS2Latlngs() node, err:", err)
				continue
			}
			pts = append(pts, latlng)

		} else if gosmparse.WayType == rm.Type {
			latlngs, err := this.getLatlngFromTmpWay(rm.ID)
			if nil != err {
				log.Println("getS2Latlngs() way, err:", err)
				continue
			}
			pts = append(pts, latlngs...)

		} else {
			continue
		}
	}
	return pts
}
