package matchers

import "fmt"

type AnyUintMatcher struct {
	Min *uint
	Max *uint
}

func (a *AnyUintMatcher) clone() *AnyUintMatcher {
	newMatcher := &AnyUintMatcher{
		Min: a.Min,
		Max: a.Max,
	}
	return newMatcher
}

func (a *AnyUintMatcher) LessThan(value uint) *AnyUintMatcher {
	newMatcher := a.clone()
	max := value - 1
	newMatcher.Max = &max
	return newMatcher
}

func (a *AnyUintMatcher) LessThanOrEqualTo(value uint) *AnyUintMatcher {
	newMatcher := a.clone()
	newMatcher.Max = &value
	return newMatcher
}

func (a *AnyUintMatcher) GreaterThan(value uint) *AnyUintMatcher {
	newMatcher := a.clone()
	min := value + 1
	newMatcher.Min = &min
	return newMatcher
}

func (a *AnyUintMatcher) GreaterThanOrEqualTo(value uint) *AnyUintMatcher {
	newMatcher := a.clone()
	newMatcher.Min = &value
	return newMatcher
}

func (a *AnyUintMatcher) Match(value any) MatchResult {
	uintValue, ok := getUint(value)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type uint, uint8, uint16, uint32 or uint64, but got %T", value),
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

func getUint(value any) (uint, bool) {
	switch v := value.(type) {
	case uint:
		return v, true
	case uint8:
		return uint(v), true
	case uint16:
		return uint(v), true
	case uint32:
		return uint(v), true
	case uint64:
		return uint(v), true
	default:
		return 0, false
	}
}
