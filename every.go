package array

func Every[T any](s []T, f func(item T, i int) bool) bool {
	for i, item := range s {
		result := f(item, i)
		if !result {
			return false
		}
	}

	return true
}
