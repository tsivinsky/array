package array

func Includes[T comparable](s []T, item T) bool {
	for _, i := range s {
		if i == item {
			return true
		}
	}

	return false
}
