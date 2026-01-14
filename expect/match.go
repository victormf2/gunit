package expect

import (
	"testing"

	"github.com/victormf2/gunit/expect/matchers"
)

type MatchResult = matchers.MatchResult

func (e *Expector) ToMatch(t *testing.T, expected any) {
	actual := e.value

	matchResult := (&matchers.GeneralMatcher{}).Matching(expected).Match(actual)
	if !matchResult.Matches {
		t.Fatal(matchResult.String())
	}
}
