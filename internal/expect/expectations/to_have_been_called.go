package expectations

import (
	"strings"

	"github.com/victormf2/gunit/internal"
	"github.com/victormf2/gunit/internal/expect"
	"github.com/victormf2/gunit/internal/expect/matchers"
	"github.com/victormf2/gunit/internal/mock"
)

func (e *expector) ToHaveBeenCalled(t internal.TestingT, callMatchers ...expect.CallMatcher) {
	t.Helper()

	_, isMockFunction := e.value.(mock.MockFunction)
	if !isMockFunction {
		t.Fatal("ToHaveBeenCalled must be used only with MockFunction")
	}

	if callMatchers == nil {
		callMatchers = []expect.CallMatcher{
			matchers.NewCallMatcher(),
		}
	}

	errors := []string{}
	for _, matcher := range callMatchers {
		matchResult := matcher.Match(e.value)
		if !matchResult.Matches() {
			errors = append(errors, matchResult.String())
		}
	}

	if len(errors) > 0 {
		message := strings.Join(errors, "\n")
		t.Fatal(message)
	}
}
