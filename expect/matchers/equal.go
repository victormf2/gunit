package matchers

import (
	"fmt"
	"reflect"
)

type EqualMatcher struct {
	Expected any
}

func (e *EqualMatcher) Match(value any) MatchResult {
	if !reflect.DeepEqual(e.Expected, value) {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected %v, but got %v", e.Expected, value),
		}
	}
	return MatchResult{Matches: true}
}
