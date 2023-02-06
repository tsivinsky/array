package array

func Some[T any](s []T, f func(item T, i int) bool) bool {
	result := false

	for i, item := range s {
		result = f(item, i)
	}

	return result
}
