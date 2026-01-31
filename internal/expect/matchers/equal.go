package matchers

import (
	"fmt"
	"reflect"

	"github.com/victormf2/gunit/internal/expect"
)

func NewEqualMatcher(expected any) expect.EqualMatcher {
	return &equalMatcher{expected: expected}
}

type equalMatcher struct {
	expected any
}

func (e *equalMatcher) Match(value any) expect.MatchResult {
	if !reflect.DeepEqual(e.expected, value) {
		return expect.DoesNotMatch(fmt.Sprintf("Expected %v, but got %v", e.expected, value), nil)
	}
	return expect.Matches()
}
