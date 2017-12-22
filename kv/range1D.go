package spatial

import (
	"math"
)

const (
	cPrefixRange        = 0x00
	cPrefixRangeReverse = 0x01
)

func (this *Spatial1D) AddRange(r Range, v interface{}) error {
	var item *_Item
	if o, ok := v.(*_Item); ok {
		item = o
	} else {
		item = NewItem(v)
	}
	vbb, err := GobSerialize(item)
	if nil != err {
		return err
	}

	r.maximizeIfZeroMax()
	w, _ := this.store.Writer()
	batch := w.NewBatch()
	batch.Set(toKey(cPrefixRange, r), vbb)
	batch.Set(toKeyReverse(cPrefixRangeReverse, r), vbb)
	w.ExecuteBatch(batch)

	return nil
}

func (this *Spatial1D) Contains(x float64) *Enum {
	return this.ContainsRange(x, x)
}

func (this *Spatial1D) ContainsRange(min, max float64) *Enum {
	if IsLessOrEqual(max, min) {
		max = min - Epsilonx10
	}

	oEnum := &Enum{ch: make(chan *_Item, 0)}
	go func() {
		rdr, _ := this.store.Reader()
		iter := rdr.RangeIterator(toKeyReverse(cPrefixRangeReverse, Range{min, max}), bbEndKeyRangeReverse)
		for iter.Valid() {
			k, v, _ := iter.Current()
			r := keyToRange(k)
			bValid := IsLessOrEqual(r.Min, min) && IsLessOrEqual(max, r.Max)
			if !bValid {
				oEnum.Close()
				break
			}

			o, err := GobDeserialize(v)
			if nil != err {
				oEnum.ch <- &_Item{Error: err}
			} else {
				oEnum.ch <- o
			}

			iter.Next()
		}

		oEnum.Close()
	}()
	return oEnum
}

func (this *Spatial1D) WithinRange(min, max float64) *Enum {
	oEnum := &Enum{ch: make(chan *_Item, 0)}
	if max <= 0 {
		max = float64(math.MaxFloat64)
	}

	go func() {
		rdr, _ := this.store.Reader()
		iter := rdr.RangeIterator(searchKey(cPrefixRange, min), bbEndKeyRange)
		for iter.Valid() {
			k, v, _ := iter.Current()
			r := keyToRange(k)
			bValid := IsLessOrEqual(r.Max, max)
			if !bValid {
				oEnum.Close()
				break
			}

			o, err := GobDeserialize(v)
			if nil != err {
				oEnum.ch <- &_Item{Error: err}
			} else {
				oEnum.ch <- o
			}

			iter.Next()
		}
		oEnum.Close()
	}()

	return oEnum
}
