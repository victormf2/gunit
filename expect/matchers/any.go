package matchers

type AnyMatcher struct{}

func (a *AnyMatcher) Match(value any) MatchResult {
	return MatchResult{Matches: true}
}
