package goo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSlice(t *testing.T) {
	t.Parallel()

	s := []int{1, 2, 3, 4, 5}
	expect := []string{"1", "2", "3", "4", "5"}

	actual := MapSlice[int, string](s, func(i int) string {
		return fmt.Sprintf("%d", i)
	})

	assert.Equal(t, expect, actual)
}

func TestReduceSlice(t *testing.T) {
	t.Parallel()

	s := []int{1, 2, 3, 4, 5}
	expect := 20

	actual := ReduceSlice[int, int](s, 5, func(v int, i int) int {
		return v + i
	})

	assert.Equal(t, expect, actual)
}
