package matchers

import (
	"fmt"

	"github.com/victormf2/gunit/internal/expect"
	"github.com/victormf2/gunit/mock"
)

func NewCallMatcher() expect.CallMatcher {
	defaultMinCalls := 1
	return &callMatcher{
		minCalls: &defaultMinCalls,
	}
}

type callMatcher struct {
	minCalls     *int
	maxCalls     *int
	expectedArgs []expect.Matcher
}

func (m *callMatcher) clone() *callMatcher {
	newMatcher := &callMatcher{
		minCalls:     m.minCalls,
		maxCalls:     m.maxCalls,
		expectedArgs: make([]expect.Matcher, len(m.expectedArgs)),
	}
	copy(newMatcher.expectedArgs, m.expectedArgs)
	return newMatcher
}

func (m *callMatcher) WithArgs(args ...any) expect.CallMatcher {
	newMatcher := m.clone()
	expectedArgs := []expect.Matcher{}
	for _, arg := range args {
		argMatcher := getArgMatcher(arg)
		expectedArgs = append(expectedArgs, argMatcher)
	}
	newMatcher.expectedArgs = expectedArgs
	return newMatcher
}

func (m *callMatcher) AtLeast(times int) expect.CallMatcher {
	newMatcher := m.clone()
	newMatcher.minCalls = &times
	return newMatcher
}

func (m *callMatcher) AtMost(times int) expect.CallMatcher {
	newMatcher := m.clone()
	newMatcher.maxCalls = &times
	return newMatcher
}

func (m *callMatcher) Times(times int) expect.CallMatcher {
	newMatcher := m.clone()
	newMatcher.minCalls = &times
	newMatcher.maxCalls = &times
	return newMatcher
}

func (m *callMatcher) Never() expect.CallMatcher {
	zeroCalls := 0
	newMatcher := m.clone()
	newMatcher.minCalls = &zeroCalls
	newMatcher.maxCalls = &zeroCalls
	return newMatcher
}

func (m *callMatcher) Match(value any) expect.MatchResult {
	mockFunction, ok := value.(mock.MockFunction)
	if !ok {
		return expect.DoesNotMatch(fmt.Sprintf("Expected a MockFunction, but got %T", value), nil)
	}

	matchedCalls := mockFunction.Calls()
	if len(m.expectedArgs) > 0 {
		matchedCalls = []mock.Call{}
		for _, call := range mockFunction.Calls() {
			args := call.Args()
			if len(args) != len(m.expectedArgs) {
				continue
			}
			matchedAllArgs := true
			for i, actualArg := range args {
				matcher := m.expectedArgs[i]
				matchResult := matcher.Match(actualArg)
				if !matchResult.Matches() {
					matchedAllArgs = false
					break
				}
			}
			if matchedAllArgs {
				matchedCalls = append(matchedCalls, call)
			}
		}
	}

	if m.minCalls != nil && len(matchedCalls) < *m.minCalls {
		return expect.DoesNotMatch(fmt.Sprintf("Expected to be called at least %d, got %d", *m.minCalls, len(matchedCalls)), nil)
	}

	if m.maxCalls != nil && len(matchedCalls) > *m.maxCalls {
		return expect.DoesNotMatch(fmt.Sprintf("Expected to be called at most %d, got %d", *m.maxCalls, len(matchedCalls)), nil)
	}

	return expect.Matches()
}

func getArgMatcher(arg any) expect.Matcher {
	switch v := arg.(type) {
	case expect.Matcher:
		return v
	default:
		return NewEqualMatcher(v)
	}
}
