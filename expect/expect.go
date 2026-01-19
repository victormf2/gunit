package expect

import (
	"github.com/victormf2/gunit/expect/matchers"
	"github.com/victormf2/gunit/gunit"
)

type Expector interface {
	ToEqual(t gunit.TestingT, expected any)
	ToHaveBeenCalled(t gunit.TestingT, callMatchers ...matchers.CallMatcher)
	ToMatch(t gunit.TestingT, expected any)
	ToPanic(t gunit.TestingT)
}

func It(value any) Expector {
	return &expector{
		value: value,
	}
}

type expector struct {
	value any
}
