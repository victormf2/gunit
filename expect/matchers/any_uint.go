package matchers

import "fmt"

type AnyUintMatcher struct {
	Min *uint
	Max *uint
}

func (a *AnyUintMatcher) LessThan(value uint) *AnyUintMatcher {
	*a.Max = value - 1
	return a
}

func (a *AnyUintMatcher) LessThanOrEqualTo(value uint) *AnyUintMatcher {
	*a.Max = value
	return a
}

func (a *AnyUintMatcher) GreaterThan(value uint) *AnyUintMatcher {
	*a.Min = value + 1
	return a
}

func (a *AnyUintMatcher) GreaterThanOrEqualTo(value uint) *AnyUintMatcher {
	*a.Min = value
	return a
}

func (a *AnyUintMatcher) Match(value any) MatchResult {
	uintValue, ok := value.(uint)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type uint, but got %T", value),
		}
	}
	if a.Min != nil && uintValue < *a.Min {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected uint >= %d, but got %d", *a.Min, uintValue),
		}
	}
	if a.Max != nil && uintValue > *a.Max {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected uint <= %d, but got %d", *a.Max, uintValue),
		}
	}
	return MatchResult{Matches: true}
}
