package matchers

import (
	"fmt"
	"reflect"
)

type EqualMatcher interface {
	Matcher
}

func NewEqualMatcher(expected any) EqualMatcher {
	return &equalMatcher{expected: expected}
}

type equalMatcher struct {
	expected any
}

func (e *equalMatcher) Match(value any) MatchResult {
	if !reflect.DeepEqual(e.expected, value) {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected %v, but got %v", e.expected, value),
		}
	}
	return MatchResult{Matches: true}
}
