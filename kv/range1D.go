package spatial

import (
	"math"

	. "github.com/noypi/spatial/common"
)

const (
	cPrefixRange        = 0x00
	cPrefixRangeReverse = 0x01
)

func (this *Spatial1D) AddRange(r Range, v interface{}) error {
	vbb, id, err := serializev(v)
	if nil != err {
		return err
	}

	r.MaximizeIfZeroMax()
	w, _ := this.store.Writer()
	batch := w.NewBatch()
	setItemToBatch(batch, r, id, vbb)
	w.ExecuteBatch(batch)

	return nil
}

func (this *Spatial1D) Contains(x float64) Enum {
	return this.ContainsRange(x, x)
}

func (this *Spatial1D) ContainsRange(min, max float64) Enum {
	if IsLessOrEqual(max, min) {
		max = min - Epsilonx10
	}

	oEnum := &_Enum{ch: make(chan Item, 0)}
	go func() {
		rdr, _ := this.store.Reader()
		iter := rdr.RangeIterator(searchKey(cPrefixRangeReverse, max), bbEndKeyRangeReverse)
		for iter.Valid() {
			k, v, _ := iter.Current()
			r, id := keyToRange(k)
			bValid := IsLessOrEqual(r.Min, min) && IsLessOrEqual(max, r.Max)
			if !bValid {
				oEnum.Close()
				break
			}

			o, err := GobDeserialize(v)
			if nil != err {
				oEnum.ch <- &_Item{err: err}
			} else {
				o.id = id
				o.kvkey = k
				o.enum = oEnum
				oEnum.ch <- o
			}

			iter.Next()
		}

		oEnum.Close()
	}()
	return oEnum
}

func (this *Spatial1D) WithinRange(min, max float64) Enum {
	oEnum := &_Enum{ch: make(chan Item, 0)}
	if max <= 0 {
		max = float64(math.MaxFloat64)
	}

	go func() {
		rdr, _ := this.store.Reader()
		iter := rdr.RangeIterator(searchKey(cPrefixRange, min), bbEndKeyRange)
		for iter.Valid() {
			k, v, _ := iter.Current()
			r, id := keyToRange(k)
			bValid := IsLessOrEqual(r.Max, max)
			if !bValid {
				oEnum.Close()
				break
			}

			o, err := GobDeserialize(v)
			if nil != err {
				oEnum.ch <- &_Item{err: err}
			} else {
				o.id = id
				o.kvkey = k
				o.enum = oEnum
				oEnum.ch <- o
			}

			iter.Next()
		}
		oEnum.Close()
	}()

	return oEnum
}
