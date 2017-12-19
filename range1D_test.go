package spatial_test

import (
	"testing"

	. "github.com/noypi/spatial"
	assertpkg "github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	assert := assertpkg.New(t)

	o := New1D()
	o.AddRange(Range{5, 10}, "5-10")
	o.AddRange(Range{6, 10}, "6-10")
	o.AddRange(Range{3, 6}, "3-6")

	//
	e := o.Get(5, 11)
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
	e = o.Get(5, 0)
	v, has = e.Next()
	assert.True(has)
	assert.Equal("5-10", v)
	v, has = e.Next()
	assert.True(has)
	assert.Equal("6-10", v)
	_, has = e.Next()
	assert.False(has)

	//
	e = o.Get(0, 6)
	v, has = e.Next()
	assert.True(has)
	assert.Equal("3-6", v)
	_, has = e.Next()
	assert.False(has)

	e.Close()
}
