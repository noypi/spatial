package osm

import (
	"github.com/blevesearch/bleve"
)

func (this *Osm) Search(q string, from, limit int) (results *bleve.SearchResult, err error) {
	if nil == this.index {
		if err = this.openIndex(); nil != err {
			return
		}
	}

	query := bleve.NewQueryStringQuery(q)
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.From = from
	searchRequest.Size = limit
	results, _ = this.index.Search(searchRequest)

	return
}
