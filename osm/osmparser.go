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
	nNodeCnt      uint64
	nWayCnt       uint64
	nRelationCnt  uint64
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

	if this.PreserveTmp {
		os.RemoveAll(this.dbpath + "/_tmpdb")
	}

}
