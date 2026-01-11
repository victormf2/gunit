package gunit

type TestingT interface {
	Fatalf(format string, args ...any)
	Fatal(args ...any)
	Helper()
}

func As[T any](value any) T {
	if value == nil {
		var zero T
		return zero
	}
	return value.(T)
}
