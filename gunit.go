package gunit

import "github.com/victormf2/gunit/internal"

type TestingT = internal.TestingT

func As[T any](value any) T {
	if value == nil {
		var zero T
		return zero
	}
	return value.(T)
}
