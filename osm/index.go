package osm

import (
	"fmt"

	"github.com/thomersch/gosmparse"
)

func (this *OsmParser) indexTags(r gosmparse.Relation) (err error) {
	if nil == this.index {
		if err = this.openIndex(); nil != err {
			return
		}
	}

	return this.index.Index(fmt.Sprintf("%d", r.ID), r.Tags)
}
