package expect

import (
	"github.com/victormf2/gunit"
)

type Expector interface {
	ToEqual(t gunit.TestingT, expected any)
	ToHaveBeenCalled(t gunit.TestingT, callMatchers ...CallMatcher)
	ToMatch(t gunit.TestingT, expected any)
	ToPanic(t gunit.TestingT)
}

type Matcher interface {
	Match(value any) MatchResult
}
type MatchResult interface {
	Matches() bool
	Message() string
	ProblemsByField() map[any]MatchResult
	String() string
}

type AnyMatcher interface {
	Matcher
}

type AnyOfMatcher[T any] interface {
	Matcher
}

type BoolMatcher interface {
	Matcher
}

type CallMatcher interface {
	Matcher
	WithArgs(args ...any) CallMatcher
	Times(times int) CallMatcher
	AtLeast(times int) CallMatcher
	AtMost(times int) CallMatcher
	Never() CallMatcher
}

type EqualMatcher interface {
	Matcher
}

type FloatMatcher interface {
	Matcher
	LessThan(value float64) FloatMatcher
	LessThanOrEqualTo(value float64) FloatMatcher
	GreaterThan(value float64) FloatMatcher
	GreaterThanOrEqualTo(value float64) FloatMatcher
	CloseTo(expectedValue float64, tolerance float64) FloatMatcher
}

type GeneralMatcher interface {
	Matcher
}

type IntMatcher interface {
	Matcher
	LessThan(value int) IntMatcher
	LessThanOrEqualTo(value int) IntMatcher
	GreaterThan(value int) IntMatcher
	GreaterThanOrEqualTo(value int) IntMatcher
}

type MapMatcher interface {
	Matcher
	WithLength(length int) MapMatcher
	WithMaxLength(max int) MapMatcher
	WithMinLength(min int) MapMatcher
	WithLengthBetween(min int, max int) MapMatcher
	ContainingAny(keyValues ...[]any) MapMatcher
	ContainingAll(keyValues ...[]any) MapMatcher
	ContainingAnyKeys(keys ...any) MapMatcher
	ContainingAllKeys(keys ...any) MapMatcher
	ContainingAnyValues(values ...any) MapMatcher
	ContainingAllValues(values ...any) MapMatcher
}

type SliceMatcher interface {
	Matcher
	WithLength(length int) SliceMatcher
	WithMaxLength(max int) SliceMatcher
	WithMinLength(min int) SliceMatcher
	WithLengthBetween(min int, max int) SliceMatcher
	ContainingAny(values ...any) SliceMatcher
	ContainingAll(values ...any) SliceMatcher
}

type StringMatcher interface {
	Matcher
	WithLength(length int) StringMatcher
	WithMaxLength(max int) StringMatcher
	WithMinLength(min int) StringMatcher
	WithLengthBetween(min int, max int) StringMatcher
	MatchingAny(values ...any) StringMatcher
	MatchingAll(values ...any) StringMatcher
	ContainingAny(values ...any) StringMatcher
	ContainingAll(values ...any) StringMatcher
}

type StructMatcher interface {
	Matcher
}

type UintMatcher interface {
	Matcher
	LessThan(value uint) UintMatcher
	LessThanOrEqualTo(value uint) UintMatcher
	GreaterThan(value uint) UintMatcher
	GreaterThanOrEqualTo(value uint) UintMatcher
}
