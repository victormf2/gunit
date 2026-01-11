package matchers

import "testing"

func TestAnyBoolMatcher(t *testing.T) {
	testCases := []struct {
		desc    string
		value   any
		matches bool
	}{
		{
			desc:    "matches true",
			value:   true,
			matches: true,
		},
		{
			desc:    "matches false",
			value:   false,
			matches: true,
		},
		{
			desc:    "does not match non bools",
			value:   "not a bool",
			matches: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			matcher := &AnyBoolMatcher{}
			result := matcher.Match(tC.value)
			if result.Matches != tC.matches {
				t.Errorf("Expected matches to be %v, but got %v", tC.matches, result.Matches)
			}
		})
	}
}
