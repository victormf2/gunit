package matchers

type MatchResult struct {
	Matches         bool
	Message         string
	ProblemsByField map[any]MatchResult
}

func (m MatchResult) String() string {
	return m.Message
}

type Matcher interface {
	Match(value any) MatchResult
}
