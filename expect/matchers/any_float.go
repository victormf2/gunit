package matchers

import "fmt"

type AnyFloatMatcher struct {
	Min          *float64
	Max          *float64
	CloseToValue *float64
	Tolerance    *float64
}

func (a *AnyFloatMatcher) clone() *AnyFloatMatcher {
	newMatcher := &AnyFloatMatcher{
		Min:          a.Min,
		Max:          a.Max,
		CloseToValue: a.CloseToValue,
		Tolerance:    a.Tolerance,
	}
	return newMatcher
}

func (a *AnyFloatMatcher) LessThan(value float64) *AnyFloatMatcher {
	newMatcher := a.clone()
	max := value - 1
	newMatcher.Max = &max
	return newMatcher
}

func (a *AnyFloatMatcher) LessThanOrEqualTo(value float64) *AnyFloatMatcher {
	newMatcher := a.clone()
	newMatcher.Max = new(float64)
	*newMatcher.Max = value
	return newMatcher
}

func (a *AnyFloatMatcher) GreaterThan(value float64) *AnyFloatMatcher {
	newMatcher := a.clone()
	min := value + 1
	newMatcher.Min = &min
	return newMatcher
}

func (a *AnyFloatMatcher) GreaterThanOrEqualTo(value float64) *AnyFloatMatcher {
	newMatcher := a.clone()
	newMatcher.Min = &value
	return newMatcher
}

func (a *AnyFloatMatcher) CloseTo(value float64, tolerance float64) *AnyFloatMatcher {
	newMatcher := a.clone()
	newMatcher.CloseToValue = &value
	newMatcher.Tolerance = &tolerance
	return newMatcher
}

func (a *AnyFloatMatcher) Match(value any) MatchResult {
	floatValue, ok := getFloat(value)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type float32 or float64, but got %T", value),
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
