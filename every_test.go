package array_test

import (
	"testing"

	"github.com/tsivinsky/array"
)

func TestEvery(t *testing.T) {
	input := []int{2, 4, 6}

	expected := true

	got := array.Every(input, func(item, i int) bool {
		return item%2 == 0
	})

	if got != expected {
		t.Errorf("array.Every doesn't work, expected: %v, got: %v", expected, got)
	}
}
