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

	cnt := atomic.AddUint64(&this.WayCnt, 1)
	if cnt < this.SkipWays {
		return
	}

	// placed outside of workerpool, because can eat large memory
	latlngs, err := this.getLatlngFromTmpNodeMulti(w.NodeIDs)
	if nil != err {
		atomic.AddUint64(&this.MissedWays, 1)
		return
	}

	this.workerPool.AddWork(func() {
		if err := this.SetExtInfo(uint8(Way), idFromInt(w.ID), &Item{
			LatLngs: latlngs,
		}); nil == err {
			this.indexTags(Way, w.ID, w.Tags)
		}

		this.syncAdd.Lock()
		if int(this.batchsize) < this.indexBatch.Size() {
			this.FlushExt()
			this.flushIndex()

			cnt := atomic.LoadUint64(&this.WayCnt) - uint64(this.workerPool.ActiveWorksCount())
			log.Println("batch way cnt=", cnt)
		}

		this.syncAdd.Unlock()

	})

}
