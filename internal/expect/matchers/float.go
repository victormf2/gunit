package matchers

import (
	"fmt"

	"github.com/victormf2/gunit/internal/expect"
)

func NewFloatMatcher() expect.FloatMatcher {
	return &floatMatcher{}
}

type floatMatcher struct {
	min           *float64
	max           *float64
	expectedValue *float64
	tolerance     *float64
}

func (a *floatMatcher) clone() *floatMatcher {
	newMatcher := &floatMatcher{
		min:           a.min,
		max:           a.max,
		expectedValue: a.expectedValue,
		tolerance:     a.tolerance,
	}
	return newMatcher
}

func (a *floatMatcher) LessThan(value float64) expect.FloatMatcher {
	newMatcher := a.clone()
	max := value - 1
	newMatcher.max = &max
	return newMatcher
}

func (a *floatMatcher) LessThanOrEqualTo(value float64) expect.FloatMatcher {
	newMatcher := a.clone()
	newMatcher.max = new(float64)
	*newMatcher.max = value
	return newMatcher
}

func (a *floatMatcher) GreaterThan(value float64) expect.FloatMatcher {
	newMatcher := a.clone()
	min := value + 1
	newMatcher.min = &min
	return newMatcher
}

func (a *floatMatcher) GreaterThanOrEqualTo(value float64) expect.FloatMatcher {
	newMatcher := a.clone()
	newMatcher.min = &value
	return newMatcher
}

func (a *floatMatcher) CloseTo(expectedValue float64, tolerance float64) expect.FloatMatcher {
	newMatcher := a.clone()
	newMatcher.expectedValue = &expectedValue
	newMatcher.tolerance = &tolerance
	return newMatcher
}

func (a *floatMatcher) Match(actualValue any) expect.MatchResult {
	actualValueFloat, ok := getFloat(actualValue)
	if !ok {
		return expect.DoesNotMatch(fmt.Sprintf("Expected type float32 or float64, but got %T", actualValue), nil)
	}
	if a.expectedValue != nil && a.tolerance != nil {
		lowerBound := *a.expectedValue - *a.tolerance
		upperBound := *a.expectedValue + *a.tolerance
		if actualValueFloat < lowerBound || actualValueFloat > upperBound {
			return expect.DoesNotMatch(fmt.Sprintf("Expected float64 close to %f ± %f, but got %f", *a.expectedValue, *a.tolerance, actualValueFloat), nil)
		}
	}
	if a.min != nil && actualValueFloat < *a.min {
		return expect.DoesNotMatch(fmt.Sprintf("Expected float64 >= %f, but got %f", *a.min, actualValueFloat), nil)
	}
	if a.max != nil && actualValueFloat > *a.max {
		return expect.DoesNotMatch(fmt.Sprintf("Expected float64 <= %f, but got %f", *a.max, actualValueFloat), nil)
	}
	return expect.Matches()
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
