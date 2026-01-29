package matchers

import "testing"

func TestAnyMatcher(t *testing.T) {
	t.Run("matches any value", func(t *testing.T) {
		matcher := NewAnyMatcher()
		values := []any{
			42,
			"hello",
			nil,
			struct{}{},
			[]int{1, 2, 3},
			map[string]int{"a": 1},
		}
		for _, value := range values {
			result := matcher.Match(value)
			if !result.Matches {
				t.Errorf("Expected any value to match, but got: %+v", result)
			}
		}
	})
}
