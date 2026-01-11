package matchers

import (
	"fmt"
	"reflect"
)

type EqualMatcher struct {
	expected any
}

func (e *EqualMatcher) Match(value any) MatchResult {
	if reflect.DeepEqual(e.expected, value) {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected %v, but got %v", e.expected, value),
		}
	}
	return MatchResult{Matches: true}
}
