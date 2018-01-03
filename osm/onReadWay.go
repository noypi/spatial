package osm

import (
	"log"
	"sync/atomic"

	"github.com/thomersch/gosmparse"
)

func (this *OsmParser) ReadWay(w gosmparse.Way) {
	if this.tCurrent != gosmparse.WayType {
		this.syncAdd.Lock()
		if 0 < this.tmpKvBatchCnt {
			this.execBatch()
		}
		this.syncAdd.Unlock()
	}
	this.tCurrent = gosmparse.WayType

	cnt := atomic.AddUint64(&this.nWayCnt, 1)
	if cnt < this.SkipWays {
		return
	}
	if (cnt % 500000) == 0 {
		log.Println("way cnt=", cnt)
	}

	latlngs, err := this.getLatlngFromTmpNodeMulti(w.NodeIDs)
	if nil != err {
		log.Println("ReadWay err:", err)
		return
	}
	this.addTmpItemLatLngs('w', w.ID, latlngs)
}
