package matchers_test

import (
	"testing"

	"github.com/victormf2/gunit/internal/expect/matchers"
)

func TestUintMatcher(t *testing.T) {
	testCases := []struct {
		desc    string
		value   any
		matches bool
	}{
		{
			desc:    "matches uint",
			value:   uint(42),
			matches: true,
		},
		{
			desc:    "matches uint8",
			value:   uint8(42),
			matches: true,
		},
		{
			desc:    "matches uint16",
			value:   uint16(42),
			matches: true,
		},
		{
			desc:    "matches uint32",
			value:   uint32(42),
			matches: true,
		},
		{
			desc:    "matches uint64",
			value:   uint64(42),
			matches: true,
		},
		{
			desc:    "does not match non-uint",
			value:   12,
			matches: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			matcher := matchers.NewUintMatcher()
			result := matcher.Match(tC.value)
			if result.Matches() != tC.matches {
				t.Errorf("Expected matches to be %v, but got %v", tC.matches, result.Matches())
			}
		})
	}

	t.Run("LessThan", func(t *testing.T) {
		matcher := matchers.NewUintMatcher().LessThan(uint(10))
		result := matcher.Match(uint(9))
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(uint(10))
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("LessThanOrEqualTo", func(t *testing.T) {
		matcher := matchers.NewUintMatcher().LessThanOrEqualTo(uint(10))
		result := matcher.Match(uint(9))
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(uint(10))
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(uint(11))
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("GreaterThan", func(t *testing.T) {
		matcher := matchers.NewUintMatcher().GreaterThan(uint(5))
		result := matcher.Match(uint(6))
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(uint(5))
		if result.Matches() {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("GreaterThanOrEqualTo", func(t *testing.T) {
		matcher := matchers.NewUintMatcher().GreaterThanOrEqualTo(uint(5))
		result := matcher.Match(uint(6))
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match(uint(5))
		if !result.Matches() {
			t.Errorf("Expected matches to be true, but got false")
		}
	})
}
