package matchers

import "fmt"

type AnyFloatMatcher struct {
	min           *float64
	max           *float64
	expectedValue *float64
	tolerance     *float64
}

func (a *AnyFloatMatcher) clone() *AnyFloatMatcher {
	newMatcher := &AnyFloatMatcher{
		min:           a.min,
		max:           a.max,
		expectedValue: a.expectedValue,
		tolerance:     a.tolerance,
	}
	return newMatcher
}

func (a *AnyFloatMatcher) LessThan(value float64) *AnyFloatMatcher {
	newMatcher := a.clone()
	max := value - 1
	newMatcher.max = &max
	return newMatcher
}

func (a *AnyFloatMatcher) LessThanOrEqualTo(value float64) *AnyFloatMatcher {
	newMatcher := a.clone()
	newMatcher.max = new(float64)
	*newMatcher.max = value
	return newMatcher
}

func (a *AnyFloatMatcher) GreaterThan(value float64) *AnyFloatMatcher {
	newMatcher := a.clone()
	min := value + 1
	newMatcher.min = &min
	return newMatcher
}

func (a *AnyFloatMatcher) GreaterThanOrEqualTo(value float64) *AnyFloatMatcher {
	newMatcher := a.clone()
	newMatcher.min = &value
	return newMatcher
}

func (a *AnyFloatMatcher) CloseTo(expectedValue float64, tolerance float64) *AnyFloatMatcher {
	newMatcher := a.clone()
	newMatcher.expectedValue = &expectedValue
	newMatcher.tolerance = &tolerance
	return newMatcher
}

func (a *AnyFloatMatcher) Match(actualValue any) MatchResult {
	actualValueFloat, ok := getFloat(actualValue)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type float32 or float64, but got %T", actualValue),
		}
	}
	if a.expectedValue != nil && a.tolerance != nil {
		lowerBound := *a.expectedValue - *a.tolerance
		upperBound := *a.expectedValue + *a.tolerance
		if actualValueFloat < lowerBound || actualValueFloat > upperBound {
			return MatchResult{
				Matches: false,
				Message: fmt.Sprintf("Expected float64 close to %f ± %f, but got %f", *a.expectedValue, *a.tolerance, actualValueFloat),
			}
		}
	}
	if a.min != nil && actualValueFloat < *a.min {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected float64 >= %f, but got %f", *a.min, actualValueFloat),
		}
	}
	if a.max != nil && actualValueFloat > *a.max {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected float64 <= %f, but got %f", *a.max, actualValueFloat),
		}
	}
	return MatchResult{Matches: true}
}

func getFloat(value any) (float64, bool) {
	switch v := value.(type) {
	case float32:
		return float64(v), true
	case float64:
		return v, true
	default:
		return 0, false
	}
}
