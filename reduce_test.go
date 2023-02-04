package array_test

import (
	"testing"

	"github.com/tsivinsky/array"
)

func TestReduce(t *testing.T) {
	nums := []int{1, 2, 3}

	expectedSum := 1 + 2 + 3

	sum := array.Reduce(nums, func(accumulator, currentValue int) int {
		return accumulator + currentValue
	}, 0)

	if sum != expectedSum {
		t.Errorf("array.Reduce is broken, expected: %d, got: %d", expectedSum, sum)
	}
}
