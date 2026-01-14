package matchers

import "fmt"

type AnyIntMatcher struct {
	min *int
	max *int
}

func (a *AnyIntMatcher) clone() *AnyIntMatcher {
	newMatcher := &AnyIntMatcher{
		min: a.min,
		max: a.max,
	}
	return newMatcher
}

func (a *AnyIntMatcher) LessThan(value int) *AnyIntMatcher {
	newMatcher := a.clone()
	max := value - 1
	newMatcher.max = &max
	return newMatcher
}

func (a *AnyIntMatcher) LessThanOrEqualTo(value int) *AnyIntMatcher {
	newMatcher := a.clone()
	newMatcher.max = &value
	return newMatcher
}

func (a *AnyIntMatcher) GreaterThan(value int) *AnyIntMatcher {
	newMatcher := a.clone()
	min := value + 1
	newMatcher.min = &min
	return newMatcher
}

func (a *AnyIntMatcher) GreaterThanOrEqualTo(value int) *AnyIntMatcher {
	newMatcher := a.clone()
	newMatcher.min = &value
	return newMatcher
}

func (a *AnyIntMatcher) Match(actualValue any) MatchResult {
	actualValueInt, ok := getInt(actualValue)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type int, int8, int16, int32 or int64, but got %T", actualValue),
		}
	}

	if a.min != nil && actualValueInt < *a.min {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected int >= %d, but got %d", *a.min, actualValueInt),
		}
	}

	if a.max != nil && actualValueInt > *a.max {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected int <= %d, but got %d", *a.max, actualValueInt),
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
