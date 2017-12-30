package spatial_test

import (
	"testing"

	"github.com/noypi/spatial/db"
	assertpkg "github.com/stretchr/testify/assert"
)

func TestExtInfo(t *testing.T) {
	assert := assertpkg.New(t)

	fnTest := func(db HasSetExtInfo, err error) {
		assert.Nil(err)
		testExtInfo_GetSet(db, assert)
	}

	fnTest(spatial.New1D())
	fnTest(spatial.New2D())
	fnTest(spatial.New3D())

}

type HasSetExtInfo interface {
	SetExtInfo(uint8, []byte, interface{}) error
	GetExtInfo(uint8, []byte) (interface{}, error)
}

func testExtInfo_GetSet(db HasSetExtInfo, assert *assertpkg.Assertions) {
	assert.Nil(db.SetExtInfo(123, []byte{3, 4}, "some value"))

	v, err := db.GetExtInfo(123, []byte{3, 4})
	assert.Nil(err)
	assert.Equal("some value", v)
}
