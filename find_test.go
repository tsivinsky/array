package array_test

import (
	"testing"

	"github.com/tsivinsky/array"
)

func TestFind(t *testing.T) {
	input := []int{1, 2, 4, 5}

	expected := 2

	got := array.Find(input, func(item, i int) bool {
		return item == 2
	})

	if got != expected {
		t.Errorf("array.Find doesn't work. Expected: %v, got: %v", expected, got)
	}
}
