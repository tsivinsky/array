package array_test

import (
	"testing"

	"github.com/tsivinsky/array"
)

var s = []int{1, 5, 69, 420, 1337}

func TestIncludesTrue(t *testing.T) {
	input := 69

	expected := true

	got := array.Includes(s, input)

	if got != expected {
		t.Errorf("array.Includes doesn't work. Expected: %v, got: %v", expected, got)
	}
}

func TestIncludesFalse(t *testing.T) {
	input := 0

	expected := false

	got := array.Includes(s, input)

	if got != expected {
		t.Errorf("array.Includes doesn't work. Expected: %v, got: %v", expected, got)
	}
}
