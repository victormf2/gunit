package expect

import (
	"github.com/victormf2/gunit/expect/matchers"
	"github.com/victormf2/gunit/gunit"
)

func (e *expector) ToEqual(t gunit.TestingT, expected any) {
	t.Helper()

	actual := e.value

	matchResult := matchers.NewEqualMatcher(expected).Match(actual)
	if !matchResult.Matches {
		t.Fatal(matchResult.String())
	}
}
