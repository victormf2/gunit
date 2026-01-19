package expectors

import (
	"github.com/victormf2/gunit/expect/matchers"
	"github.com/victormf2/gunit/gunit"
)

type MatchResult = matchers.MatchResult

func (e *expector) ToMatch(t gunit.TestingT, expected any) {
	t.Helper()

	actual := e.value

	matchResult := matchers.NewGeneralMatcher(expected).Match(actual)
	if !matchResult.Matches {
		t.Fatal(matchResult.String())
	}
}
