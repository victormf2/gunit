package expectations

import (
	"github.com/victormf2/gunit/internal"
	"github.com/victormf2/gunit/internal/expect"
	"github.com/victormf2/gunit/internal/expect/matchers"
)

func (e *expector) ToMatch(t internal.TestingT, expected any) {
	t.Helper()

	actual := e.value

	matcher, ok := expected.(expect.Matcher)
	if ok {
		matchResult := matcher.Match(actual)
		if !matchResult.Matches() {
			t.Fatal(matchResult.String())
		}
	}

	matchResult := matchers.NewGeneralMatcher(expected).Match(actual)
	if !matchResult.Matches() {
		t.Fatal(matchResult.String())
	}
}
