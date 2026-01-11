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
		matcher := &AnyMapMatcher{}
		matcher.WithLength(2)
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
		matcher := &AnyMapMatcher{}
		matcher.WithMaxLength(2)
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
		matcher := &AnyMapMatcher{}
		matcher.WithMinLength(2)
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
		matcher := &AnyMapMatcher{}
		matcher.WithLengthBetween(2, 4)
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
}
