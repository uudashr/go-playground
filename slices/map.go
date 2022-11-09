package slices

func Map[T any, O any](slice []T, fn func(T) O) []O {
	result := make([]O, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}
