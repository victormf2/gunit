package matchers

import (
	"context"
	"fmt"
	"reflect"
)

type GeneralMatcher interface {
	Matcher
}

func NewGeneralMatcher(expected any) GeneralMatcher {
	return &generalMatcher{
		expected:      expected,
		expectedValue: reflect.ValueOf(expected),
	}
}

type generalMatcher struct {
	expected      any
	expectedValue reflect.Value
}

func (g *generalMatcher) Match(actualValueInterface any) MatchResult {
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

	actualCtx, ok := actualValueInterface.(context.Context)
	if ok {
		expectedCtx, ok := g.expected.(context.Context)
		if ok {
			return NewContextMatcher(expectedCtx).Match(actualCtx)
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
		return (&structMatcher{}).matching(g.expected).Match(actualValueInterface)
	case reflect.Map:
		return (&mapMatcher{}).matching(g.expected).Match(actualValueInterface)
	case reflect.Array, reflect.Slice:
		return (&sliceMatcher{}).matching(g.expected).Match(actualValueInterface)
	default:
		return (&equalMatcher{expected: g.expected}).Match(actualValueInterface)
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
