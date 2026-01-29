package matchers_test

import (
	"testing"

	"github.com/victormf2/gunit/internal/matchers"
)

func TestAnyOf(t *testing.T) {
	t.Run("matches correct type", func(t *testing.T) {
		matcher := matchers.NewAnyOfMatcher[int]()
		result := matcher.Match(42)
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match("not an int")
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("does not match incorrect type", func(t *testing.T) {
		matcher := matchers.NewAnyOfMatcher[string]()
		result := matcher.Match(42)
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match("a string")
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
	})
}
