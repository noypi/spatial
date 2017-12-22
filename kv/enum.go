package spatial

import (
	"sync"
)

type Enum struct {
	syncCh sync.Mutex
	closed bool
	ch     chan *_Item
}

func (this *Enum) Next() (v *_Item, has bool) {
	v, has = <-this.ch
	return
}

func (this *Enum) Close() {
	this.syncCh.Lock()
	if !this.closed {
		close(this.ch)
		this.closed = true
	}
	this.syncCh.Unlock()
}
