package array

func At[T any](s []T, i int) T {
	if i < 0 {
		return s[len(s)+i]
	}

	return s[i]
}
