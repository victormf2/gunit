package testo

func As[T any](value any) T {
	if value == nil {
		var zero T
		return zero
	}
	return value.(T)
}
