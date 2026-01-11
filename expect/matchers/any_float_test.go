package matchers

import "testing"

func TestAnyFloatMatcher(t *testing.T) {
	testCases := []struct {
		desc    string
		value   any
		matches bool
	}{
		{
			desc:    "matches float32",
			value:   float32(3.14),
			matches: true,
		},
		{
			desc:    "matches float64",
			value:   float64(2.718),
			matches: true,
		},
		{
			desc:    "does not match non-float",
			value:   42,
			matches: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			matcher := &AnyFloatMatcher{}
			result := matcher.Match(tC.value)
			if result.Matches != tC.matches {
				t.Errorf("Expected matches to be %v, but got %v", tC.matches, result.Matches)
			}
		})
	}

	t.Run("LessThan", func(t *testing.T) {
		matcher := &AnyFloatMatcher{}
		matcher.LessThan(10.0)
		result := matcher.Match(9.0)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(10.0)
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("LessThanOrEqualTo", func(t *testing.T) {
		matcher := &AnyFloatMatcher{}
		matcher.LessThanOrEqualTo(10.0)
		result := matcher.Match(9.0)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(10.0)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(11.0)
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("GreaterThan", func(t *testing.T) {
		matcher := &AnyFloatMatcher{}
		matcher.GreaterThan(5.0)
		result := matcher.Match(6.0)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(5.0)
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("GreaterThanOrEqualTo", func(t *testing.T) {
		matcher := &AnyFloatMatcher{}
		matcher.GreaterThanOrEqualTo(5.0)
		result := matcher.Match(6.0)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(5.0)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(4.0)
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("CloseTo", func(t *testing.T) {
		matcher := &AnyFloatMatcher{}
		matcher.CloseTo(10.0, 0.5)
		result := matcher.Match(10.3)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(10.6)
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})
}
