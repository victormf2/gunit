package matchers

import (
	"fmt"

	"github.com/victormf2/gunit/internal/expect"
)

func NewBoolMatcher() expect.BoolMatcher {
	return &boolMatcher{}
}

type boolMatcher struct{}

func (a *boolMatcher) Match(actualValue any) expect.MatchResult {
	_, ok := actualValue.(bool)
	if !ok {
		return expect.DoesNotMatch(fmt.Sprintf("Expected type bool, but got %T", actualValue), nil)
	}
	return expect.Matches()
}
