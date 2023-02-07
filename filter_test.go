package array_test

import (
	"testing"

	"github.com/tsivinsky/array"
)

func TestFilter(t *testing.T) {
	input := []int{1, 3, 5, 2, 10, 69, 42}

	expected := []int{2, 10, 42}

	got := array.Filter(input, func(item, i int) bool {
		return item%2 == 0
	})

	if len(got) != len(expected) {
		t.Errorf("array.Filter doesn't work. Expected: %v, got: %v", expected, got)
	}
}
