package array_test

import (
	"testing"

	"github.com/tsivinsky/array"
)

func TestFindIndex(t *testing.T) {
	input := []int{1, 2, 5, 6}

	expected := 2

	got := array.FindIndex(input, func(item, i int) bool {
		return item == 5
	})

	if got != expected {
		t.Errorf("array.findIndex doesn't work. Expected: %v, got: %v", expected, got)
	}
}
