package osm

import (
	"github.com/blevesearch/bleve"
)

func (this *Osm) Search(q string) (results *bleve.SearchResult, err error) {
	if nil == this.index {
		if err = this.openIndex(); nil != err {
			return
		}
	}

	query := bleve.NewQueryStringQuery(q)
	searchRequest := bleve.NewSearchRequest(query)
	results, _ = this.index.Search(searchRequest)

	return
}
