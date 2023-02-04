package array_test

import (
	"testing"

	"github.com/tsivinsky/array"
)

func TestEach(t *testing.T) {
	nums := []int{1, 2, 3}

	doubles := nums

	array.Each(nums, func(num int, i int) {
		doubles[i] = num * 2
	})

	if doubles[0] != 2 || doubles[1] != 4 || doubles[2] != 6 {
		t.Errorf("array.Each not working properly, passed function didn't double numbers")
	}
}
