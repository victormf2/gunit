package matchers

type AnyMatcher interface {
	Matcher
}

func NewAnyMatcher() AnyMatcher {
	return &anyMatcher{}
}

type anyMatcher struct{}

func (a *anyMatcher) Match(value any) MatchResult {
	return MatchResult{Matches: true}
}
