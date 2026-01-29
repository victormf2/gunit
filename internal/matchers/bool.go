package matchers

import "fmt"

type BoolMatcher interface {
	Matcher
}

func NewBoolMatcher() BoolMatcher {
	return &boolMatcher{}
}

type boolMatcher struct{}

func (a *boolMatcher) Match(actualValue any) MatchResult {
	_, ok := actualValue.(bool)
	if !ok {
		return MatchResult{Matches: false, Message: fmt.Sprintf("Expected type bool, but got %T", actualValue)}
	}
	return MatchResult{Matches: true}
}
