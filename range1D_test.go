package spatial_test

import (
	"testing"

	. "github.com/noypi/spatial"
	assertpkg "github.com/stretchr/testify/assert"
)

func TestWithinRange(t *testing.T) {
	assert := assertpkg.New(t)

	o := New1D()
	o.AddRange(Range{5, 10}, "5-10")
	o.AddRange(Range{6, 10}, "6-10")
	o.AddRange(Range{3, 6}, "3-6")

	//
	e := o.WithinRange(5, 11)
	v, has := e.Next()
	assert.True(has)
	assert.Equal("5-10", v)
	v, has = e.Next()
	assert.True(has)
	assert.Equal("6-10", v)
	_, has = e.Next()
	assert.False(has)

	e.Close()

	//
	e = o.WithinRange(5, 0)
	v, has = e.Next()
	assert.True(has)
	assert.Equal("5-10", v)
	v, has = e.Next()
	assert.True(has)
	assert.Equal("6-10", v)
	_, has = e.Next()
	assert.False(has)

	//
	e = o.WithinRange(0, 6)
	v, has = e.Next()
	assert.True(has)
	assert.Equal("3-6", v)
	_, has = e.Next()
	assert.False(has)

	e.Close()
}

func TestContainsRange(t *testing.T) {
	assert := assertpkg.New(t)

	o := New1D()
	o.AddRange(Range{5, 10}, "5-10")
	o.AddRange(Range{6, 10}, "6-10")
	o.AddRange(Range{3, 8}, "3-8")

	//
	e := o.ContainsRange(5, 7)
	v, has := e.Next()
	assert.Equal("5-10", v)
	assert.True(has)
	v, has = e.Next()
	assert.Equal("3-8", v)
	assert.True(has)
	_, has = e.Next()
	assert.False(has)

	e.Close()

	//
	e = o.ContainsRange(4, 7)
	v, has = e.Next()
	assert.Equal("3-8", v)
	assert.True(has)
	v, has = e.Next()
	_, has = e.Next()
	assert.False(has)

	e.Close()
}
