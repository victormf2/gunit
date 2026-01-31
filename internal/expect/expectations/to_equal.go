package expectations

import (
	"github.com/victormf2/gunit/internal"
	"github.com/victormf2/gunit/internal/expect/matchers"
)

func (e *expector) ToEqual(t internal.TestingT, expected any) {
	t.Helper()

	actual := e.value

	matchResult := matchers.NewEqualMatcher(expected).Match(actual)
	if !matchResult.Matches() {
		t.Fatal(matchResult.String())
	}
}
