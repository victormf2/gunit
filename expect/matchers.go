package expect

import (
	"github.com/victormf2/gunit/internal/expect/matchers"
)

func Any() AnyMatcher {
	return matchers.NewAnyMatcher()
}

func AnyOf[T any]() AnyOfMatcher[T] {
	return matchers.NewAnyOfMatcher[T]()
}

func Bool() BoolMatcher {
	return matchers.NewBoolMatcher()
}

func Call() CallMatcher {
	return matchers.NewCallMatcher()
}

func Equal(expected any) EqualMatcher {
	return matchers.NewEqualMatcher(expected)
}

func Float() FloatMatcher {
	return matchers.NewFloatMatcher()
}

func Int() IntMatcher {
	return matchers.NewIntMatcher()
}

func Map() MapMatcher {
	return matchers.NewMapMatcher()
}

func Matching(expected any) GeneralMatcher {
	return matchers.NewGeneralMatcher(expected)
}

func Slice() SliceMatcher {
	return matchers.NewSliceMatcher()
}

func String() StringMatcher {
	return matchers.NewStringMatcher()
}

func Struct() StructMatcher {
	return matchers.NewStructMatcher()
}

func Uint() UintMatcher {
	return matchers.NewUintMatcher()
}
