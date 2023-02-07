package array

func Find[T any](s []T, f func(item T, i int) bool) T {
	var result T

	for i, item := range s {
		x := f(item, i)
		if x {
			result = item
		}
	}

	return result
}
