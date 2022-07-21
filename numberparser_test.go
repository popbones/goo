package goo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNumber(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v, err := ParseNumber[int]("42")
		assert.NoError(t, err)
		assert.Equal(t, int(42), v)
	})

	t.Run("int8", func(t *testing.T) {
		t.Parallel()
		v, err := ParseNumber[int8]("42")
		assert.NoError(t, err)
		assert.Equal(t, int8(42), v)
	})

	t.Run("int16", func(t *testing.T) {
		t.Parallel()
		v, err := ParseNumber[int16]("42")
		assert.NoError(t, err)
		assert.Equal(t, int16(42), v)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		v, err := ParseNumber[int32]("42")
		assert.NoError(t, err)
		assert.Equal(t, int32(42), v)
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		v, err := ParseNumber[int64]("42")
		assert.NoError(t, err)
		assert.Equal(t, int64(42), v)
	})

	t.Run("uint", func(t *testing.T) {
		t.Parallel()
		v, err := ParseNumber[uint]("42")
		assert.NoError(t, err)
		assert.Equal(t, uint(42), v)
	})

	t.Run("uint8", func(t *testing.T) {
		t.Parallel()
		v, err := ParseNumber[uint8]("42")
		assert.NoError(t, err)
		assert.Equal(t, uint8(42), v)
	})

	t.Run("uint16", func(t *testing.T) {
		t.Parallel()
		v, err := ParseNumber[uint16]("42")
		assert.NoError(t, err)
		assert.Equal(t, uint16(42), v)
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		v, err := ParseNumber[uint32]("42")
		assert.NoError(t, err)
		assert.Equal(t, uint32(42), v)
	})

	t.Run("uint64", func(t *testing.T) {
		t.Parallel()
		v, err := ParseNumber[uint64]("42")
		assert.NoError(t, err)
		assert.Equal(t, uint64(42), v)
	})

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		v, err := ParseNumber[float32]("42.0")
		assert.NoError(t, err)
		assert.Equal(t, float32(42.0), v)
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		v, err := ParseNumber[float64]("42.0")
		assert.NoError(t, err)
		assert.Equal(t, float64(42.0), v)
	})
}
