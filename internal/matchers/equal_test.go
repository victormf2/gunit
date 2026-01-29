package matchers_test

import (
	"testing"

	"github.com/victormf2/gunit/internal/matchers"
)

func TestEqualMatcher(t *testing.T) {
	testCases := []struct {
		desc     string
		expected any
		actual   any
		matches  bool
	}{
		{
			desc:     "matches equal values",
			expected: 42,
			actual:   42,
			matches:  true,
		},
		{
			desc:     "does not match unequal values",
			expected: 42,
			actual:   43,
			matches:  false,
		},
		{
			desc:     "matches nil values",
			expected: nil,
			actual:   nil,
			matches:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			matcher := matchers.NewEqualMatcher(tC.expected)
			result := matcher.Match(tC.actual)
			if result.Matches != tC.matches {
				t.Errorf("Expected match result to be %v, but got %v", tC.matches, result.Matches)
			}
		})
	}
}
