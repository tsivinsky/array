package array_test

import (
	"testing"

	"github.com/tsivinsky/array"
)

var input = []int{1, 2, 3, 4}

func TestAtFirst(t *testing.T) {
	expected := 1
	got := array.At(input, 0)

	if got != expected {
		t.Errorf("array.At doesn't work. Expected: %v, got: %v", expected, got)
	}
}

func TestAtLast(t *testing.T) {
	expected := 4
	got := array.At(input, -1)

	if got != expected {
		t.Errorf("array.At doesn't work. Expected: %v, got: %v", expected, got)
	}
}
