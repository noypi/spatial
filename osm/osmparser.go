package osm

import (
	"fmt"
	"os"
	"sync"

	"github.com/noypi/kv"
	"github.com/noypi/kv/leveldb"
	"github.com/thomersch/gosmparse"
)

type OsmParser struct {
	*Osm

	syncAdd       sync.Mutex
	NodeCnt       uint64
	WayCnt        uint64
	RelationCnt   uint64
	tCurrent      gosmparse.MemberType
	tmpKv         kv.KVStore
	tmpKvWrtr     kv.KVWriter
	tmpKvBatch    kv.KVBatch
	tmpKvBatchCnt uint32
	dbpath        string

	batchsize uint32

	SkipNodes     uint64
	SkipWays      uint64
	SkipRelations uint64

	PreserveTmp bool

	MissedWays      uint64
	MissedRelations uint64
}

func (this OsmParser) DbPath() string {
	return this.dbpath
}

func (this *OsmParser) LoadTempKV() (err error) {
	this.PreserveTmp = true
	return this.useTempKV()
}

func (this *OsmParser) useTempKV() (err error) {
	if nil != this.tmpKv {
		return
	}

	this.tmpKv, err = leveldb.GetDefault(this.dbpath + "/_tmpdb")
	if nil != err {
		err = fmt.Errorf("%v, while: useTempKV GetDefault.", err)
		return
	}
	this.tmpKvWrtr, err = this.tmpKv.Writer()
	if nil != err {
		return
	}
	this.tmpKvBatch = this.tmpKvWrtr.NewBatch()
	return
}

func (this *OsmParser) Close() {
	this.SpatialGeo.Close()
	this.Cleanup()
}

func (this *OsmParser) Cleanup() {
	if nil != this.tmpKvWrtr {
		this.tmpKvWrtr.ExecuteBatch(this.tmpKvBatch)
		this.tmpKvBatch.Close()
		this.tmpKvBatch = nil
		this.tmpKvBatchCnt = 0
		this.tmpKvWrtr.Close()
		this.tmpKvWrtr = nil
		this.tmpKv.Close()
	}

	if !this.PreserveTmp {
		os.RemoveAll(this.dbpath + "/_tmpdb")
	}

}

func (this OsmParser) MissedWaysPercent() float64 {
	return float64(this.MissedWays) / float64(this.WayCnt)
}

func (this OsmParser) MissedRelationsPercent() float64 {
	return float64(this.MissedRelations) / float64(this.RelationCnt)
}

func (this *OsmParser) CountWaysInTmpKV() int {
	return this.countInTmpKV([]byte{'w'})
}

func (this *OsmParser) CountNodesInTmpKV() int {
	return this.countInTmpKV([]byte{'n'})
}

func (this *OsmParser) countInTmpKV(prefix []byte) int {
	rdr, err := this.tmpKv.Reader()
	if nil != err {
		panic("countInTmpKV err=" + err.Error())
	}
	iter := rdr.PrefixIterator(prefix)
	defer iter.Close()
	return iter.Count()
}
