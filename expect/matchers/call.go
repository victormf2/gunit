package matchers

import (
	"fmt"

	"github.com/victormf2/gunit/mock"
)

type CallMatcher interface {
	Matcher
	WithArgs(args ...any) CallMatcher
	Times(times int) CallMatcher
	AtLeast(times int) CallMatcher
	AtMost(times int) CallMatcher
	Never() CallMatcher
}

func NewCallMatcher() CallMatcher {
	defaultMinCalls := 1
	return &callMatcher{
		minCalls: &defaultMinCalls,
	}
}

type callMatcher struct {
	minCalls     *int
	maxCalls     *int
	expectedArgs []Matcher
}

func (m *callMatcher) clone() *callMatcher {
	newMatcher := &callMatcher{
		minCalls:     m.minCalls,
		maxCalls:     m.maxCalls,
		expectedArgs: make([]Matcher, len(m.expectedArgs)),
	}
	copy(newMatcher.expectedArgs, m.expectedArgs)
	return newMatcher
}

func (m *callMatcher) WithArgs(args ...any) CallMatcher {
	newMatcher := m.clone()
	expectedArgs := []Matcher{}
	for _, arg := range args {
		argMatcher := getArgMatcher(arg)
		expectedArgs = append(expectedArgs, argMatcher)
	}
	newMatcher.expectedArgs = expectedArgs
	return newMatcher
}

func (m *callMatcher) AtLeast(times int) CallMatcher {
	newMatcher := m.clone()
	newMatcher.minCalls = &times
	return newMatcher
}

func (m *callMatcher) AtMost(times int) CallMatcher {
	newMatcher := m.clone()
	newMatcher.maxCalls = &times
	return newMatcher
}

func (m *callMatcher) Times(times int) CallMatcher {
	newMatcher := m.clone()
	newMatcher.minCalls = &times
	newMatcher.maxCalls = &times
	return newMatcher
}

func (m *callMatcher) Never() CallMatcher {
	zeroCalls := 0
	newMatcher := m.clone()
	newMatcher.minCalls = &zeroCalls
	newMatcher.maxCalls = &zeroCalls
	return newMatcher
}

func (m *callMatcher) Match(value any) MatchResult {
	mockFunction, ok := value.(*mock.MockFunction)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected a *MockFunction, but got %T", value),
		}
	}

	matchedCalls := mockFunction.Calls()
	if len(m.expectedArgs) > 0 {
		matchedCalls = []mock.Call{}
		for _, call := range mockFunction.Calls() {
			if len(call.Args) != len(m.expectedArgs) {
				continue
			}
			matchedAllArgs := true
			for i, actualArg := range call.Args {
				matcher := m.expectedArgs[i]
				matchResult := matcher.Match(actualArg)
				if !matchResult.Matches {
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
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected to be called at least %d, got %d", *m.minCalls, len(matchedCalls)),
		}
	}

	if m.maxCalls != nil && len(matchedCalls) > *m.maxCalls {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected to be called at most %d, got %d", *m.maxCalls, len(matchedCalls)),
		}
	}

	return MatchResult{Matches: true}
}

func getArgMatcher(arg any) Matcher {
	switch v := arg.(type) {
	case Matcher:
		return v
	default:
		return NewEqualMatcher(v)
	}
}
