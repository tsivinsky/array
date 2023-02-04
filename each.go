package array

func Each[T any](s []T, f func(item T, i int)) {
	for i, item := range s {
		f(item, i)
	}
}
