package array

func Map[T any, M any](s []T, f func(item T, i int) M) []M {
	var result []M

	for i, item := range s {
		result = append(result, f(item, i))
	}

	return result
}
