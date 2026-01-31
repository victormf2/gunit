package matchers

import "github.com/victormf2/gunit/internal/expect"

func NewAnyMatcher() expect.AnyMatcher {
	return &anyMatcher{}
}

type anyMatcher struct{}

func (a *anyMatcher) Match(value any) expect.MatchResult {
	return expect.Matches()
}
