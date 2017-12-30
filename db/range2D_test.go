package spatial_test

import (
	"testing"

	. "github.com/noypi/spatial/common"
	. "github.com/noypi/spatial/db"
	assertpkg "github.com/stretchr/testify/assert"
)

func TestWithinRange2D(t *testing.T) {
	assert := assertpkg.New(t)

	o := New2D()
	o.Set([]byte{1}, Range{5, 10}, Range{0, 11}, "5-10, 0-11")
	o.Set([]byte{2}, Range{6, 10}, Range{4, 6}, "6-10, 4-6")
	o.Set([]byte{3}, Range{3, 6}, Range{8, 15}, "3-6, 8-15")

	//
	e := o.WithinRange(Range{5, 11}, Range{0, 11})
	v, has := e.Next()
	assert.True(has)
	assert.Equal("5-10, 0-11", v.Value())
	v, has = e.Next()
	assert.True(has)
	assert.Equal("6-10, 4-6", v.Value())
	_, has = e.Next()
	assert.False(has)

	e.Close()

	//
	e = o.WithinRange(Range{5, 0}, Range{0, 0})
	v, has = e.Next()
	assert.True(has)
	assert.Equal("5-10, 0-11", v.Value())
	v, has = e.Next()
	assert.True(has)
	assert.Equal("6-10, 4-6", v.Value())
	_, has = e.Next()
	assert.False(has)

	//
	e = o.WithinRange(Range{0, 6}, Range{7, 15})
	v, has = e.Next()
	assert.True(has)
	assert.Equal("3-6, 8-15", v.Value())
	_, has = e.Next()
	assert.False(has)

	e.Close()
}
