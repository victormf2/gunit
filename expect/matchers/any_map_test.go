package matchers

import "testing"

func TestAnyMap(t *testing.T) {
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
			matcher := &AnyMapMatcher{}
			result := matcher.Match(tC.value)
			if result.Matches != tC.matches {
				t.Errorf("Expected matches to be %v, but got %v", tC.matches, result.Matches)
			}
		})
	}

	t.Run("WithLength", func(t *testing.T) {
		matcher := (&AnyMapMatcher{}).WithLength(2)
		result := matcher.Match(map[string]int{"a": 1, "b": 2})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("WithMaxLength", func(t *testing.T) {
		matcher := (&AnyMapMatcher{}).WithMaxLength(2)
		result := matcher.Match(map[string]int{"a": 1})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})
	t.Run("WithMinLength", func(t *testing.T) {
		matcher := (&AnyMapMatcher{}).WithMinLength(2)
		result := matcher.Match(map[string]int{"a": 1})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
	})
	t.Run("WithLengthBetween", func(t *testing.T) {
		matcher := (&AnyMapMatcher{}).WithLengthBetween(2, 4)
		result := matcher.Match(map[string]int{"a": 1})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("Containing", func(t *testing.T) {
		matcher := (&AnyMapMatcher{}).Containing(
			[]any{"a", 1},
			[]any{"b", 2},
			[]any{(&AnyStringMatcher{}).Containing("d"), 4},
			[]any{"e", (&AnyIntMatcher{}).GreaterThan(4)},
		)
		result := matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "c": 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"doo": 4})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"e": 5})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"b": 3})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"c": 1})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("ContainingAll", func(t *testing.T) {
		matcher := (&AnyMapMatcher{}).ContainingAll(
			[]any{"a", 1},
			[]any{"b", 2},
			[]any{(&AnyStringMatcher{}).Containing("c"), (&AnyIntMatcher{}).GreaterThan(2)},
		)
		result := matcher.Match(map[string]int{"a": 1, "b": 2, "coo": 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "c": 3})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 3, "c": 3})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("ContainingKeys", func(t *testing.T) {
		matcher := (&AnyMapMatcher{}).ContainingKeys(
			"a",
			(&AnyStringMatcher{}).Containing("b"),
		)
		result := matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"boo": 2})
		if !result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"c": 3})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("ContainingAllKeys", func(t *testing.T) {
		matcher := (&AnyMapMatcher{}).ContainingAllKeys(
			"a",
			(&AnyStringMatcher{}).Containing("b"),
		)
		result := matcher.Match(map[string]int{"a": 1, "boo": 2, "c": 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "b": 2})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match(map[string]int{"boo": 2})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("ContainingValues", func(t *testing.T) {
		matcher := (&AnyMapMatcher{}).ContainingValues(
			1,
			(&AnyIntMatcher{}).GreaterThan(2),
		)
		result := matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"c": 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"b": 2})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("ContainingAllValues", func(t *testing.T) {
		matcher := (&AnyMapMatcher{}).ContainingAllValues(
			1,
			(&AnyIntMatcher{}).GreaterThan(2),
		)
		result := matcher.Match(map[string]int{"a": 1, "b": 2, "c": 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"a": 1, "c": 3})
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(map[string]int{"b": 2})
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})
}
