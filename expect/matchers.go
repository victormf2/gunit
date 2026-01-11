package expect

import (
	"github.com/victormf2/gunit/expect/matchers"
)

type Matcher = matchers.Matcher

func Any() Matcher {
	return &matchers.AnyMatcher{}
}

func AnyOf[T any]() Matcher {
	return &matchers.AnyOfMatcher[T]{}
}

func AnyInt() Matcher {
	return &matchers.AnyIntMatcher{}
}

func AnyUint() Matcher {
	return &matchers.AnyUintMatcher{}
}

func AnyFloat() Matcher {
	return &matchers.AnyFloatMatcher{}
}

func AnyString() Matcher {
	return &matchers.AnyStringMatcher{}
}

func AnyBool() Matcher {
	return &matchers.AnyBoolMatcher{}
}

func AnySlice() Matcher {
	return &matchers.AnySliceMatcher{}
}

func AnyMap() Matcher {
	return &matchers.AnyMapMatcher{}
}

func Equal(expected any) Matcher {
	return &matchers.EqualMatcher{Expected: expected}
}
