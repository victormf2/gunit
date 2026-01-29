package expectors

import (
	"github.com/victormf2/gunit/gunit"
	"github.com/victormf2/gunit/internal/matchers"
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
