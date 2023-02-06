package array_test

import (
	"testing"

	"github.com/tsivinsky/array"
)

func TestSome(t *testing.T) {
	input := []int{1, 3, 5, 6}

	expected := true

	got := array.Some(input, func(item, i int) bool {
		return item%2 == 0
	})

	if got != expected {
		t.Errorf("array.Some doesn't work. Expected %v, got: %v\n", expected, got)
	}
}
