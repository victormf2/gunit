package expect

import (
	"github.com/victormf2/gunit/expect/matchers"
)

type Matcher = matchers.Matcher

func Any() matchers.AnyMatcher {
	return matchers.NewAnyMatcher()
}

func AnyOf[T any]() matchers.AnyOfMatcher[T] {
	return matchers.NewAnyOfMatcher[T]()
}

func Bool() matchers.BoolMatcher {
	return matchers.NewBoolMatcher()
}

func Equal(expected any) Matcher {
	return matchers.NewEqualMatcher(expected)
}

func Float() matchers.FloatMatcher {
	return matchers.NewFloatMatcher()
}

func Int() matchers.IntMatcher {
	return matchers.NewIntMatcher()
}

func Map() matchers.MapMatcher {
	return matchers.NewMapMatcher()
}

func Matching(expected any) Matcher {
	return matchers.NewGeneralMatcher(expected)
}

func Slice() matchers.SliceMatcher {
	return matchers.NewSliceMatcher()
}

func String() matchers.StringMatcher {
	return matchers.NewStringMatcher()
}

func Struct() matchers.StructMatcher {
	return matchers.NewStructMatcher()
}

func Uint() matchers.UintMatcher {
	return matchers.NewUintMatcher()
}
