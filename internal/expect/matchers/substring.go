package matchers

import (
	"fmt"
	"strings"

	"github.com/victormf2/gunit/internal/expect"
)

type SubstringMatcher interface {
	expect.Matcher
}

func NewSubstringMatcher(substring string) SubstringMatcher {
	return &substringMatcher{substring: substring}
}

type substringMatcher struct {
	substring string
}

func (s *substringMatcher) Match(value any) expect.MatchResult {
	strValue, ok := value.(string)
	if !ok {
		return expect.DoesNotMatch(fmt.Sprintf("Expected type string, but got %T", value), nil)
	}
	if !strings.Contains(strValue, s.substring) {
		return expect.DoesNotMatch(fmt.Sprintf("Expected string to contain '%s', but it did not", s.substring), nil)
	}
	return expect.Matches()
}
