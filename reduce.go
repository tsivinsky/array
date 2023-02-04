package array

func Reduce[T any, M any](s []T, f func(accumulator M, currentValue T) M, initialValue M) M {
	result := initialValue

	for _, item := range s {
		result = f(result, item)
	}

	return result
}
