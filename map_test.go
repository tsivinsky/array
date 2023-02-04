package array_test

import (
	"testing"

	"github.com/tsivinsky/array"
)

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3}

	doubles := array.Map(nums, func(item, i int) int {
		return item * 2
	})

	if doubles[0] != 2 || doubles[1] != 4 || doubles[2] != 6 {
		t.Errorf("array.Map not working properly, returned: %v, expected: %v", doubles, []int{2, 4, 6})
	}
}
