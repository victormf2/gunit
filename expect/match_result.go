package expect

import "github.com/victormf2/gunit/internal/expect"

func Matches() MatchResult {
	return expect.Matches()
}

func DoesNotMatch(message string, problemsByField map[any]MatchResult) MatchResult {
	return expect.DoesNotMatch(message, problemsByField)
}
