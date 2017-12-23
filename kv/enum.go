package spatial

import (
	"sync"

	. "github.com/noypi/spatial/common"
)

type _Enum struct {
	syncCh    sync.Mutex
	closed    bool
	ch        chan Item
	kvcommits []func()
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
		if 0 < len(this.kvcommits) {
			for _, c := range this.kvcommits {
				c()
			}
		}
		this.kvcommits = nil
	}
	this.syncCh.Unlock()
}

func (this *_Enum) addtocommit(c func()) {
	this.syncCh.Lock()
	this.kvcommits = append(this.kvcommits, c)
	this.syncCh.Unlock()
}
