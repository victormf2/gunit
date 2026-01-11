package matchers

import "fmt"

type AnyIntMatcher struct {
	Min *int
	Max *int
}

func (a *AnyIntMatcher) LessThan(value int) *AnyIntMatcher {
	*a.Max = value - 1
	return a
}

func (a *AnyIntMatcher) LessThanOrEqualTo(value int) *AnyIntMatcher {
	*a.Max = value
	return a
}

func (a *AnyIntMatcher) GreaterThan(value int) *AnyIntMatcher {
	*a.Min = value + 1
	return a
}

func (a *AnyIntMatcher) GreaterThanOrEqualTo(value int) *AnyIntMatcher {
	*a.Min = value
	return a
}

func (a *AnyIntMatcher) Match(value any) MatchResult {
	intValue, ok := value.(int)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type int, but got %T", value),
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
