package matchers_test

import (
	"testing"

	"github.com/victormf2/gunit/internal/expect/matchers"
)

func TestMapMatcher(t *testing.T) {
	testCases := []struct {
		desc    string
		value   any
		matches bool
	}{
		{
			desc:    "matches map",
			value:   map[string]int{"a": 1, "b": 2},
			matches: true,
		},
		{
			desc:    "does not match non-map",
			value:   42,
			matches: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			matcher := matchers.NewMapMatcher()
			result := matcher.Match(tC.value)
			if result.Matches() != tC.matches {
				t.Errorf("Expected matches to be %v, but got %v", tC.matches, result.Matches)
			}
		})
	}

	t.Run("WithLength", func(t *testing.T) {
		matcher := matchers.NewMapMatcher().WithLength(2)
		result := matcher.Match(map[string]int{"a": 1, "b": 2})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("WithMaxLength", func(t *testing.T) {
		matcher := matchers.NewMapMatcher().WithMaxLength(2)
		result := matcher.Match(map[string]int{"a": 1})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
	})
	t.Run("WithMinLength", func(t *testing.T) {
		matcher := matchers.NewMapMatcher().WithMinLength(2)
		result := matcher.Match(map[string]int{"a": 1})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
	})
	t.Run("WithLengthBetween", func(t *testing.T) {
		matcher := matchers.NewMapMatcher().WithLengthBetween(2, 4)
		result := matcher.Match(map[string]int{"a": 1})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("ContainingAny", func(t *testing.T) {
		matcher := matchers.NewMapMatcher().ContainingAny(
			[]any{"a", 1},
			[]any{"b", 2},
			[]any{matchers.NewStringMatcher().ContainingAny("d"), 4},
			[]any{"e", matchers.NewIntMatcher().GreaterThan(4)},
		)
		result := matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "c": 3})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"doo": 4})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"e": 5})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"b": 3})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"c": 1})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("ContainingAll", func(t *testing.T) {
		matcher := matchers.NewMapMatcher().ContainingAll(
			[]any{"a", 1},
			[]any{"b", 2},
			[]any{matchers.NewStringMatcher().ContainingAny("c"), matchers.NewIntMatcher().GreaterThan(2)},
		)
		result := matcher.Match(map[string]int{"a": 1, "b": 2, "coo": 3})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "c": 3})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 3, "c": 3})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("ContainingKeys", func(t *testing.T) {
		matcher := matchers.NewMapMatcher().ContainingAnyKeys(
			"a",
			matchers.NewStringMatcher().ContainingAny("b"),
		)
		result := matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"boo": 2})
		if !result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"c": 3})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("ContainingAllKeys", func(t *testing.T) {
		matcher := matchers.NewMapMatcher().ContainingAllKeys(
			"a",
			matchers.NewStringMatcher().ContainingAny("b"),
		)
		result := matcher.Match(map[string]int{"a": 1, "boo": 2, "c": 3})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"boo": 2})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("ContainingValues", func(t *testing.T) {
		matcher := matchers.NewMapMatcher().ContainingAnyValues(
			1,
			matchers.NewIntMatcher().GreaterThan(2),
		)
		result := matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"c": 3})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"b": 2})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("ContainingAllValues", func(t *testing.T) {
		matcher := matchers.NewMapMatcher().ContainingAllValues(
			1,
			matchers.NewIntMatcher().GreaterThan(2),
		)
		result := matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "c": 3})
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"b": 2})
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
	})
}
