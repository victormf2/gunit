package expect

import "maps"

func Matches() *matchResult {
	return &matchResult{matches: true}
}

func DoesNotMatch(message string, problemsByField map[any]MatchResult) *matchResult {
	return &matchResult{matches: false, message: message, problemsByField: problemsByField}
}

type matchResult struct {
	matches         bool
	message         string
	problemsByField map[any]MatchResult
}

func (r *matchResult) Matches() bool {
	return r.matches
}

func (r *matchResult) Message() string {
	return r.message
}

func (r *matchResult) ProblemsByField() map[any]MatchResult {
	return maps.Clone(r.problemsByField)
}
func (r *matchResult) String() string {
	return r.message
}
