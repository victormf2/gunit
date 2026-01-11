package matchers

import "fmt"

type AnyFloatMatcher struct {
	Min          *float64
	Max          *float64
	CloseToValue *float64
	Tolerance    *float64
}

func (a *AnyFloatMatcher) LessThan(value float64) *AnyFloatMatcher {
	*a.Max = value - 1
	return a
}

func (a *AnyFloatMatcher) LessThanOrEqualTo(value float64) *AnyFloatMatcher {
	*a.Max = value
	return a
}

func (a *AnyFloatMatcher) GreaterThan(value float64) *AnyFloatMatcher {
	*a.Min = value + 1
	return a
}

func (a *AnyFloatMatcher) GreaterThanOrEqualTo(value float64) *AnyFloatMatcher {
	*a.Min = value
	return a
}

func (a *AnyFloatMatcher) CloseTo(value float64, tolerance float64) *AnyFloatMatcher {
	*a.CloseToValue = value
	*a.Tolerance = tolerance
	return a
}

func (a *AnyFloatMatcher) Match(value any) MatchResult {
	floatValue, ok := value.(float64)
	if !ok {
		floatValue32, ok := value.(float32)
		if ok {
			floatValue = float64(floatValue32)
		} else {
			return MatchResult{
				Matches: false,
				Message: fmt.Sprintf("Expected type float64, but got %T", value),
			}
		}
	}
	if a.CloseToValue != nil && a.Tolerance != nil {
		lowerBound := *a.CloseToValue - *a.Tolerance
		upperBound := *a.CloseToValue + *a.Tolerance
		if floatValue < lowerBound || floatValue > upperBound {
			return MatchResult{
				Matches: false,
				Message: fmt.Sprintf("Expected float64 close to %f ± %f, but got %f", *a.CloseToValue, *a.Tolerance, floatValue),
			}
		}
	}
	if a.Min != nil && floatValue < *a.Min {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected float64 >= %f, but got %f", *a.Min, floatValue),
		}
	}
	if a.Max != nil && floatValue > *a.Max {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected float64 <= %f, but got %f", *a.Max, floatValue),
		}
	}
	return MatchResult{Matches: true}
}
