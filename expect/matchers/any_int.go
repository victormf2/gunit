package matchers

import "fmt"

type AnyIntMatcher struct {
	Min *int
	Max *int
}

func (a *AnyIntMatcher) clone() *AnyIntMatcher {
	newMatcher := &AnyIntMatcher{
		Min: a.Min,
		Max: a.Max,
	}
	return newMatcher
}

func (a *AnyIntMatcher) LessThan(value int) *AnyIntMatcher {
	newMatcher := a.clone()
	max := value - 1
	newMatcher.Max = &max
	return newMatcher
}

func (a *AnyIntMatcher) LessThanOrEqualTo(value int) *AnyIntMatcher {
	newMatcher := a.clone()
	newMatcher.Max = &value
	return newMatcher
}

func (a *AnyIntMatcher) GreaterThan(value int) *AnyIntMatcher {
	newMatcher := a.clone()
	min := value + 1
	newMatcher.Min = &min
	return newMatcher
}

func (a *AnyIntMatcher) GreaterThanOrEqualTo(value int) *AnyIntMatcher {
	newMatcher := a.clone()
	newMatcher.Min = &value
	return newMatcher
}

func (a *AnyIntMatcher) Match(value any) MatchResult {
	intValue, ok := getInt(value)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type int, int8, int16, int32 or int64, but got %T", value),
		}
	}

	if a.Min != nil && intValue < *a.Min {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected int >= %d, but got %d", *a.Min, intValue),
		}
	}

	if a.Max != nil && intValue > *a.Max {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected int <= %d, but got %d", *a.Max, intValue),
		}
	}

	return MatchResult{Matches: true}
}

func getInt(value any) (int, bool) {
	switch v := value.(type) {
	case int:
		return v, true
	case int8:
		return int(v), true
	case int16:
		return int(v), true
	case int32:
		return int(v), true
	case int64:
		return int(v), true
	default:
		return 0, false
	}
}
