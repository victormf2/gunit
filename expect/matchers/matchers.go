package matchers

type MatchResult struct {
	Matches bool
	Message string
}

type Matcher interface {
	Match(value any) MatchResult
}
