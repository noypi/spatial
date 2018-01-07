package osm

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve"
)

func (this *OsmParser) indexTags(t tCategory, id int64, tags map[string]string) (err error) {
	if nil == this.index {
		if err = this.openIndex(); nil != err {
			return
		}
	}

	this.syncIndex.Lock()
	defer this.syncIndex.Unlock()

	if nil == this.indexBatch {
		this.indexBatch = this.index.NewBatch()
	}

	return this.indexBatch.Index(toSearchId(t, id), tags)
}

func (this *OsmParser) flushIndex() {
	this.syncIndex.Lock()
	defer this.syncIndex.Unlock()
	if nil != this.index && nil != this.indexBatch {
		this.index.Batch(this.indexBatch)
		this.indexBatch.Reset()
	}
}

func (this *OsmParser) cleanupIndex() {
	this.syncIndex.Lock()
	defer this.syncIndex.Unlock()
	if nil == this.index {
		this.index.Close()
		this.index = nil
		this.indexBatch = nil
	}
}

func toSearchId(t tCategory, id int64) (sid string) {
	return fmt.Sprintf("%c%d", toPrefix(t), id)
}

func (this *Osm) openIndex() (err error) {
	this.syncIndex.Lock()
	defer this.syncIndex.Unlock()
	if nil != this.index {
		return
	}

	fpath := this.myDbPath() + "/index.bleve"
	_, err = os.Stat(fpath)
	if os.IsNotExist(err) {
		indexMap := bleve.NewIndexMapping()
		if this.index, err = bleve.New(fpath, indexMap); nil != err {
			return
		}

	} else {
		if this.index, err = bleve.Open(fpath); nil != err {
			return
		}
	}

	return
}
