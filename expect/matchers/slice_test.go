package matchers_test

import (
	"testing"

	"github.com/victormf2/gunit/expect/matchers"
)

func TestSliceMatcher(t *testing.T) {
	testCases := []struct {
		desc    string
		value   any
		matches bool
	}{
		{
			desc:    "matches slice",
			value:   []int{1, 2, 3},
			matches: true,
		},
		{
			desc:    "matches array",
			value:   [3]int{1, 2, 3},
			matches: true,
		},
		{
			desc:    "does not match non-slice",
			value:   42,
			matches: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			matcher := matchers.NewSliceMatcher()
			result := matcher.Match(tC.value)
			if result.Matches != tC.matches {
				t.Errorf("Expected matches to be %v, but got %v", tC.matches, result.Matches)
			}
		})
	}

	t.Run("WithLength", func(t *testing.T) {
		matcher := matchers.NewSliceMatcher().WithLength(3)
		result := matcher.Match([]int{1, 2, 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match([]int{1, 2})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}

		result = matcher.Match([]int{1, 2, 3, 4})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("WithMinLength", func(t *testing.T) {
		matcher := matchers.NewSliceMatcher().WithMinLength(2)
		result := matcher.Match([]int{1, 2, 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match([]int{1, 2})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match([]int{1})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("WithMaxLength", func(t *testing.T) {
		matcher := matchers.NewSliceMatcher().WithMaxLength(2)
		result := matcher.Match([]int{1})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match([]int{1, 2})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match([]int{1, 2, 3})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("WithLengthBetween", func(t *testing.T) {
		matcher := matchers.NewSliceMatcher().WithLengthBetween(2, 4)
		result := matcher.Match([]int{1})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match([]int{1, 2})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match([]int{1, 2, 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match([]int{1, 2, 3, 4})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match([]int{1, 2, 3, 4, 5})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("Containing", func(t *testing.T) {
		matcher := matchers.NewSliceMatcher().Containing(
			matchers.NewStringMatcher(),
			42,
		)
		result := matcher.Match([]any{"hello", 42, 3.14})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match([]any{"hello"})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match([]any{42})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match([]any{3.14})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("ContainingAll", func(t *testing.T) {
		matcher := matchers.NewSliceMatcher().ContainingAll(
			matchers.NewStringMatcher(),
			42,
		)
		result := matcher.Match([]any{"hello", 42, 3.14})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match([]any{"hello", 3.14})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match([]any{42, 3.14})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})
}
