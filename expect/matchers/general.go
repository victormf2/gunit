package matchers

import (
	"fmt"
	"reflect"
)

type GeneralMatcher struct {
	expected      any
	expectedValue reflect.Value
}

func (g *GeneralMatcher) clone() *GeneralMatcher {
	newMatcher := &GeneralMatcher{
		expectedValue: g.expectedValue,
	}
	return newMatcher
}

func (g *GeneralMatcher) Matching(expected any) *GeneralMatcher {
	newMatcher := g.clone()
	newMatcher.expected = expected
	newMatcher.expectedValue = reflect.ValueOf(expected)
	return newMatcher
}

func (g *GeneralMatcher) Match(actualValueInterface any) MatchResult {
	actualValue := reflect.ValueOf(actualValueInterface)
	// Handle nil pointers
	if isNil(actualValue) {
		if isNil(g.expectedValue) {
			return MatchResult{Matches: true} // Both nil, match
		}
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected %v, but got nil", g.expectedValue.Interface()),
		}
	}
	if isNil(g.expectedValue) {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected nil, but got %v", actualValue.Interface()),
		}
	}

	// Dereference pointers
	if actualValue.Kind() == reflect.Ptr {
		actualValue = actualValue.Elem()
	}
	if g.expectedValue.Kind() == reflect.Ptr {
		g.expectedValue = g.expectedValue.Elem()
	}

	if !compatibleKinds(actualValue, g.expectedValue) {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Type mismatch: expected %s, got %s", g.expectedValue.Kind(), actualValue.Kind()),
		}
	}

	switch actualValue.Kind() {
	case reflect.Struct:
		return (&AnyStructMatcher{}).matching(g.expected).Match(actualValueInterface)
	case reflect.Map:
		return (&AnyMapMatcher{}).matching(g.expected).Match(actualValueInterface)
	case reflect.Array, reflect.Slice:
		return (&AnySliceMatcher{}).matching(g.expected).Match(actualValueInterface)
	default:
		return (&EqualMatcher{Expected: g.expected}).Match(actualValueInterface)
	}
}

func isNil(value reflect.Value) bool {
	if !value.IsValid() {
		return true
	}
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		if value.IsNil() {
			return true
		}
	}
	return false
}

func compatibleKinds(actualValue reflect.Value, expectedValue reflect.Value) bool {
	if actualValue.Kind() == expectedValue.Kind() {
		return true
	}
	if actualValue.Kind() == reflect.Array && expectedValue.Kind() == reflect.Slice {
		return true
	}
	if actualValue.Kind() == reflect.Slice && expectedValue.Kind() == reflect.Array {
		return true
	}
	return false
}
