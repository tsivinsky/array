package array

func Filter[T any](s []T, f func(item T, i int) bool) []T {
	var result []T

	for i, item := range s {
		x := f(item, i)
		if x {
			result = append(result, item)
		}
	}

	return result
}
