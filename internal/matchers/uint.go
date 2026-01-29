package matchers

import "fmt"

type UintMatcher interface {
	Matcher
	LessThan(value uint) UintMatcher
	LessThanOrEqualTo(value uint) UintMatcher
	GreaterThan(value uint) UintMatcher
	GreaterThanOrEqualTo(value uint) UintMatcher
}

func NewUintMatcher() UintMatcher {
	return &uintMatcher{}
}

type uintMatcher struct {
	min *uint
	max *uint
}

func (a *uintMatcher) clone() *uintMatcher {
	newMatcher := &uintMatcher{
		min: a.min,
		max: a.max,
	}
	return newMatcher
}

func (a *uintMatcher) LessThan(value uint) UintMatcher {
	newMatcher := a.clone()
	max := value - 1
	newMatcher.max = &max
	return newMatcher
}

func (a *uintMatcher) LessThanOrEqualTo(value uint) UintMatcher {
	newMatcher := a.clone()
	newMatcher.max = &value
	return newMatcher
}

func (a *uintMatcher) GreaterThan(value uint) UintMatcher {
	newMatcher := a.clone()
	min := value + 1
	newMatcher.min = &min
	return newMatcher
}

func (a *uintMatcher) GreaterThanOrEqualTo(value uint) UintMatcher {
	newMatcher := a.clone()
	newMatcher.min = &value
	return newMatcher
}

func (a *uintMatcher) Match(value any) MatchResult {
	uintValue, ok := getUint(value)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type uint, uint8, uint16, uint32 or uint64, but got %T", value),
		}
	}
	if a.min != nil && uintValue < *a.min {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected uint >= %d, but got %d", *a.min, uintValue),
		}
	}
	if a.max != nil && uintValue > *a.max {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected uint <= %d, but got %d", *a.max, uintValue),
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
