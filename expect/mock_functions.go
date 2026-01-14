package expect

import (
	"testing"

	"github.com/victormf2/gunit/mock"
)

func (e *Expector) ToHaveBeenCalled(t *testing.T, opts ...CallMatchOption) {
	fn := e.value.(*mock.MockFunction)
	callExpectations := &callExpectations{
		min: 1,
		max: -1,
	}
	for _, opt := range opts {
		opt(callExpectations)
	}
	callExpectations.assert(t, fn)
}

type callExpectations struct {
	min  int
	max  int
	args []any
}

func (e *callExpectations) assert(t *testing.T, fn *mock.MockFunction) {
	matchingCalls := e.matchingCalls(fn)

	callCount := len(matchingCalls)
	if callCount < e.min {
		t.Fatalf("expected function to have been called at least %d times, actual call count was %d", e.min, callCount)
	}

	if e.max >= 0 && callCount > e.max {
		t.Fatalf("expected function to have been called at most %d times, actual call count was %d", e.max, callCount)
	}
}

func (e *callExpectations) matchingCalls(fn *mock.MockFunction) []mock.Call {
	if e.args == nil {
		return fn.Calls()
	}

	matchingCalls := []mock.Call{}
	for _, call := range fn.Calls() {
		if len(call.Args) != len(e.args) {
			continue
		}

		matches := true
		for i, actualArg := range call.Args {
			expectedArg := e.args[i]
			matchResult := Matching(expectedArg).Match(actualArg)
			if !matchResult.Matches {
				matches = false
			}
		}
		if matches {
			matchingCalls = append(matchingCalls, call)
		}
	}

	return matchingCalls
}

type CallMatchOption func(ce *callExpectations)

func With(args ...any) CallMatchOption {
	return func(ce *callExpectations) {
		ce.args = args
	}
}

func Times(n int) CallMatchOption {
	return func(ce *callExpectations) {
		ce.min = n
		ce.max = n
	}
}

func TimesAtLeast(n int) CallMatchOption {
	return func(ce *callExpectations) {
		ce.min = n
	}
}

func TimesAtMost(n int) CallMatchOption {
	return func(ce *callExpectations) {
		ce.max = n
	}
}
