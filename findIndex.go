package array

func FindIndex[T any](s []T, f func(item T, i int) bool) int {
	for i, item := range s {
		x := f(item, i)
		if x {
			return i
		}
	}

	return -1
}
