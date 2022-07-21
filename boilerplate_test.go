package goo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFallback(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		fallback := 42
		assert.Equal(t, 42, Fallback(0, fallback))
		assert.Equal(t, 31, Fallback(31, fallback))
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		fallback := "foobar"
		assert.Equal(t, "foobar", Fallback("", fallback))
		assert.Equal(t, "baz", Fallback("baz", fallback))
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()
		type dummy struct {
			Foo int
		}
		fallback := dummy{Foo: 42}
		assert.Equal(t, dummy{Foo: 42}, Fallback(dummy{}, fallback))
	})

}

func TestMust(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name  string
		Value any
		Error error
	}{
		{
			Name:  "simple",
			Value: 42,
			Error: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Value, Must(tc.Value, tc.Error))
		})
	}
}
