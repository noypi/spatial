package spatial_test

import (
	"testing"

	. "github.com/noypi/spatial/common"
	. "github.com/noypi/spatial/db"
	assertpkg "github.com/stretchr/testify/assert"
)

func TestWithinRange(t *testing.T) {
	assert := assertpkg.New(t)

	o, err := New1D()
	assert.Nil(err)
	err = o.Set([]byte{1}, Range{5, 10}, "5-10")
	assert.Nil(err)
	err = o.Set([]byte{2}, Range{6, 10}, "6-10")
	assert.Nil(err)
	err = o.Set([]byte{3}, Range{3, 6}, "3-6")
	assert.Nil(err)

	//
	e := o.WithinRange(5, 11)
	v, has := e.Next()
	assert.Equal("5-10", v.Value())
	assert.True(has)
	v, has = e.Next()
	assert.Equal("6-10", v.Value())
	assert.True(has)
	_, has = e.Next()
	assert.False(has)

	e.Close()

	//
	e = o.WithinRange(5, 0)
	v, has = e.Next()
	assert.True(has)
	assert.Equal("5-10", v.Value())
	v, has = e.Next()
	assert.True(has)
	assert.Equal("6-10", v.Value())
	_, has = e.Next()
	assert.False(has)

	//
	e = o.WithinRange(0, 6)
	v, has = e.Next()
	assert.True(has)
	assert.Equal("3-6", v.Value())
	_, has = e.Next()
	assert.False(has)

	e.Close()
}

func aTestContainsRange(t *testing.T) {
	assert := assertpkg.New(t)

	o, err := New1D()
	assert.Nil(err)
	err = o.Set([]byte{1}, Range{5, 10}, "5-10")
	assert.Nil(err)
	err = o.Set([]byte{2}, Range{6, 10}, "6-10")
	assert.Nil(err)
	err = o.Set([]byte{3}, Range{3, 8}, "3-8")
	assert.Nil(err)

	//
	e := o.ContainsRange(5, 7)
	v, has := e.Next()
	assert.Equal("3-8", v.Value())
	assert.True(has)
	v, has = e.Next()
	assert.Equal("5-10", v.Value())
	assert.True(has)
	_, has = e.Next()
	assert.False(has)

	e.Close()

	//
	e = o.ContainsRange(5, 10)
	v, has = e.Next()
	assert.Equal("5-10", v.Value())
	assert.True(has)
	_, has = e.Next()
	assert.False(has)

	e.Close()

	//
	e = o.ContainsRange(4, 7)
	v, has = e.Next()
	assert.Equal("3-8", v.Value())
	assert.True(has)
	_, has = e.Next()
	assert.False(has)

	e.Close()
}

func TestContains(t *testing.T) {
	assert := assertpkg.New(t)

	o, err := New1D()
	assert.Nil(err)
	assert.Nil(o.Set([]byte{1}, Range{5, 10}, "5-10"))
	assert.Nil(o.Set([]byte{2}, Range{6, 10}, "6-10"))
	assert.Nil(o.Set([]byte{3}, Range{3, 8}, "3-8"))

	//
	e := o.Contains(5)
	v, has := e.Next()
	assert.Equal("3-8", v.Value())
	assert.True(has)
	v, has = e.Next()
	assert.Equal("5-10", v.Value())
	assert.True(has)
	_, has = e.Next()
	assert.False(has)

	e.Close()

	//
	e = o.Contains(4)
	v, has = e.Next()
	assert.Equal("3-8", v.Value())
	assert.True(has)
	_, has = e.Next()
	assert.False(has)

	//
	e = o.Contains(3)
	v, has = e.Next()
	assert.Equal("3-8", v.Value())
	assert.True(has)
	_, has = e.Next()
	assert.False(has)

	//
	e = o.Contains(10)
	v, has = e.Next()
	assert.Equal("5-10", v.Value())
	assert.True(has)
	v, has = e.Next()
	assert.Equal("6-10", v.Value())
	assert.True(has)
	_, has = e.Next()
	assert.False(has)

	e.Close()
}
