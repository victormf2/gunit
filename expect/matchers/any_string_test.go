package matchers

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc    string
		value   any
		matches bool
	}{
		{
			desc:    "matches string",
			value:   "hello",
			matches: true,
		},
		{
			desc:    "does not match non-string",
			value:   42,
			matches: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			matcher := &AnyStringMatcher{}
			result := matcher.Match(tC.value)
			if result.Matches != tC.matches {
				t.Errorf("Expected matches to be %v, but got %v", tC.matches, result.Matches)
			}
		})
	}

	t.Run("WithLength", func(t *testing.T) {
		matcher := &AnyStringMatcher{}
		matcher.WithLength(5)
		result := matcher.Match("hello")
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match("hi")
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match("helloo")
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("WithMaxLength", func(t *testing.T) {
		matcher := &AnyStringMatcher{}
		matcher.WithMaxLength(5)
		result := matcher.Match("hi")
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match("hello")
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match("helloo")
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("WithMinLength", func(t *testing.T) {
		matcher := &AnyStringMatcher{}
		matcher.WithMinLength(3)
		result := matcher.Match("hi")
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match("hey")
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match("hello")
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
	})

	t.Run("WithLengthBetween", func(t *testing.T) {
		matcher := &AnyStringMatcher{}
		matcher.WithLengthBetween(3, 5)
		result := matcher.Match("hi")
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match("hey")
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match("hell")
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match("hello")
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match("helloo")
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})

	t.Run("Containing", func(t *testing.T) {
		matcher := &AnyStringMatcher{}
		matcher.ContainingAll("foo", "bar")
		result := matcher.Match("foobar")
		if !result.Matches {
			t.Errorf("Expected matches to be true, but got false")
		}
		result = matcher.Match("foo")
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match("bar")
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
		result = matcher.Match("baz")
		if result.Matches {
			t.Errorf("Expected matches to be false, but got true")
		}
	})
}
