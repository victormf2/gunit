package matchers_test

import (
	"testing"

	"github.com/victormf2/gunit/expect/matchers"
	"github.com/victormf2/gunit/mock"
)

func TestCallMatcher(t *testing.T) {
	t.Run("non *MockFunction", func(t *testing.T) {
		match := matchers.NewCallMatcher().Match(func() {})
		if match.Matches {
			t.Errorf("expected non *MockFunction to not match")
		}
	})
	t.Run("default", func(t *testing.T) {
		mockFunctionNeverCalled := mock.NewMockFunction("fn", func() {})
		mockFunctionNeverCalled.Call()

		mockFunction := mock.NewMockFunction("fn", func() {})

		defaultMatcher := matchers.NewCallMatcher()
		match0 := defaultMatcher.Match(mockFunction)

		mockFunction.Call()
		match1 := defaultMatcher.Match(mockFunction)

		mockFunction.Call()
		match2 := defaultMatcher.Match(mockFunction)

		if match0.Matches {
			t.Errorf("expected default matcher to not match a function never called")
		}

		if !match1.Matches {
			t.Errorf("expected default matcher to match a function called once")
		}

		if !match2.Matches {
			t.Errorf("expected default matcher to match a function called more than once")
		}
	})

	t.Run("Times", func(t *testing.T) {
		mockFunction := mock.NewMockFunction("fn", func() {})

		mockFunction.Call()

		match0 := matchers.NewCallMatcher().Times(0).Match(mockFunction)
		match1 := matchers.NewCallMatcher().Times(1).Match(mockFunction)
		match2 := matchers.NewCallMatcher().Times(2).Match(mockFunction)

		if match0.Matches {
			t.Errorf("expected Times(0) to not match a function called once")
		}

		if !match1.Matches {
			t.Errorf("expected Times(1) to match a function called once")
		}

		if match2.Matches {
			t.Errorf("expected Times(2) to not match a function called once")
		}
	})

	t.Run("AtLeast", func(t *testing.T) {
		mockFunction := mock.NewMockFunction("fn", func() {})

		mockFunction.Call()

		match0 := matchers.NewCallMatcher().AtLeast(0).Match(mockFunction)
		match1 := matchers.NewCallMatcher().AtLeast(1).Match(mockFunction)
		match2 := matchers.NewCallMatcher().AtLeast(2).Match(mockFunction)

		if !match0.Matches {
			t.Errorf("expected AtLeast(0) to match a function called once")
		}

		if !match1.Matches {
			t.Errorf("expected AtLeast(1) to match a function called once")
		}

		if match2.Matches {
			t.Errorf("expected AtLeast(2) to not match a function called once")
		}
	})

	t.Run("AtMost", func(t *testing.T) {
		mockFunction := mock.NewMockFunction("fn", func() {})

		mockFunction.Call()

		match0 := matchers.NewCallMatcher().AtMost(0).Match(mockFunction)
		match1 := matchers.NewCallMatcher().AtMost(1).Match(mockFunction)
		match2 := matchers.NewCallMatcher().AtMost(2).Match(mockFunction)

		if match0.Matches {
			t.Errorf("expected AtMost(0) to not match a function called once")
		}

		if !match1.Matches {
			t.Errorf("expected AtMost(1) to match a function called once")
		}

		if !match2.Matches {
			t.Errorf("expected AtMost(2) to match a function called once")
		}
	})

	t.Run("Never", func(t *testing.T) {
		mockFunction := mock.NewMockFunction("fn", func() {})

		match0 := matchers.NewCallMatcher().Never().Match(mockFunction)

		mockFunction.Call()
		match1 := matchers.NewCallMatcher().Never().Match(mockFunction)

		if !match0.Matches {
			t.Errorf("expected Never() to not match a function never called")
		}

		if match1.Matches {
			t.Errorf("expected Never() to not match a function called once")
		}
	})

	t.Run("WithArgs", func(t *testing.T) {

		mockFunction := mock.NewMockFunction("fn", func(a string, b int) {})
		mockFunction.Call("a", 1)

		testCases := []struct {
			desc    string
			matcher matchers.CallMatcher
			matches bool
		}{
			{
				desc:    "matches function called with same args",
				matcher: matchers.NewCallMatcher().WithArgs("a", 1),
				matches: true,
			},
			{
				desc: "matches function called with args matchers",
				matcher: matchers.NewCallMatcher().WithArgs(
					matchers.NewStringMatcher(),
					matchers.NewIntMatcher(),
				),
				matches: true,
			},
			{
				desc:    "doesn't match function called with other args",
				matcher: matchers.NewCallMatcher().WithArgs("b", 1),
				matches: false,
			},
			{
				desc:    "doesn't match function called with less args",
				matcher: matchers.NewCallMatcher().WithArgs("a", 1, true),
				matches: false,
			},
			{
				desc:    "doesn't match function called with more args",
				matcher: matchers.NewCallMatcher().WithArgs("a"),
				matches: false,
			},
			{
				desc:    "doesn't match .Never() called with same args",
				matcher: matchers.NewCallMatcher().WithArgs("a", 1).Never(),
				matches: false,
			},
			{
				desc:    "matches .Never() called with other args",
				matcher: matchers.NewCallMatcher().WithArgs("b", 1).Never(),
				matches: true,
			},
			{
				desc:    "doesn't match .Times() called wrong number of times",
				matcher: matchers.NewCallMatcher().WithArgs("a", 1).Times(2),
				matches: false,
			},
			{
				desc:    "matches .Times() called exact number of times",
				matcher: matchers.NewCallMatcher().WithArgs("a", 1).Times(1),
				matches: true,
			},
		}
		for _, tC := range testCases {
			t.Run(tC.desc, func(t *testing.T) {
				matchResult := tC.matcher.Match(mockFunction)
				if matchResult.Matches != tC.matches {
					t.Errorf("expected matches to be %v, but got %v", tC.matches, matchResult.Matches)
				}
			})
		}
	})
}
