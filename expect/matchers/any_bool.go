package matchers

import "fmt"

type AnyBoolMatcher struct{}

func (a *AnyBoolMatcher) Match(actualValue any) MatchResult {
	_, ok := actualValue.(bool)
	if !ok {
		return MatchResult{Matches: false, Message: fmt.Sprintf("Expected type bool, but got %T", actualValue)}
	}
	return MatchResult{Matches: true}
}
