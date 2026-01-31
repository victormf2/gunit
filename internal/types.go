package internal

type TestingT interface {
	Fatalf(format string, args ...any)
	Fatal(args ...any)
	Helper()
}
