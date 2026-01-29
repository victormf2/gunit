package expectors

import (
	"strings"

	"github.com/victormf2/gunit/gunit"
	"github.com/victormf2/gunit/internal/matchers"
	"github.com/victormf2/gunit/mock"
)

func (e *expector) ToHaveBeenCalled(t gunit.TestingT, callMatchers ...matchers.CallMatcher) {
	t.Helper()

	_, isMockFunction := e.value.(*mock.MockFunction)
	if !isMockFunction {
		t.Fatal("ToHaveBeenCalled must be used only with *MockFunction")
	}

	if callMatchers == nil {
		callMatchers = []matchers.CallMatcher{
			matchers.NewCallMatcher(),
		}
	}

	errors := []string{}
	for _, matcher := range callMatchers {
		matchResult := matcher.Match(e.value)
		if !matchResult.Matches {
			errors = append(errors, matchResult.String())
		}
	}

	if len(errors) > 0 {
		message := strings.Join(errors, "\n")
		t.Fatal(message)
	}
}
