package matchers_test

import (
	"testing"

	"github.com/victormf2/gunit/expect/matchers"
)

func TestIntMatcher(t *testing.T) {
	testCases := []struct {
		desc    string
		value   any
		matches bool
	}{
		{
			desc:    "matches int",
			value:   42,
			matches: true,
		},
		{
			desc:    "matches int8",
			value:   int8(42),
			matches: true,
		},
		{
			desc:    "matches int16",
			value:   int16(42),
			matches: true,
		},
		{
			desc:    "matches int32",
			value:   int32(42),
			matches: true,
		},
		{
			desc:    "matches int64",
			value:   int64(42),
			matches: true,
		},
		{
			desc:    "does not match non-int",
			value:   12.34,
			matches: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			matcher := matchers.NewIntMatcher()
			result := matcher.Match(tC.value)
			if result.Matches != tC.matches {
				t.Errorf("Expected matches to be %v, but got %v", tC.matches, result.Matches)
			}
		})
	}

	t.Run("LessThan", func(t *testing.T) {
		matcher := matchers.NewIntMatcher().LessThan(10)
		result := matcher.Match(9)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(10)
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("LessThanOrEqualTo", func(t *testing.T) {
		matcher := matchers.NewIntMatcher().LessThanOrEqualTo(10)
		result := matcher.Match(9)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(10)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(11)
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("GreaterThan", func(t *testing.T) {
		matcher := matchers.NewIntMatcher().GreaterThan(5)
		result := matcher.Match(6)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(5)
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("GreaterThanOrEqualTo", func(t *testing.T) {
		matcher := matchers.NewIntMatcher().GreaterThanOrEqualTo(5)
		result := matcher.Match(6)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(5)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
	})
}
