package spatial

import (
	"sync"

	. "github.com/noypi/spatial/common"
)

type _Enum struct {
	syncCh sync.Mutex
	closed bool
	ch     chan Item
}

func (this *_Enum) Next() (v Item, has bool) {
	v, has = <-this.ch
	return
}

func (this *_Enum) Close() {
	this.syncCh.Lock()
	if !this.closed {
		close(this.ch)
		this.closed = true
	}
	this.syncCh.Unlock()
}
