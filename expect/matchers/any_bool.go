package matchers

import "fmt"

type AnyBoolMatcher struct{}

func (a *AnyBoolMatcher) Match(value any) MatchResult {
	_, ok := value.(bool)
	if !ok {
		return MatchResult{Matches: false, Message: fmt.Sprintf("Expected type bool, but got %T", value)}
	}
	return MatchResult{Matches: true}
}
