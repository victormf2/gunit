package matchers

import (
	"fmt"
	"strings"
)

type SubstringMatcher interface {
	Matcher
}

func NewSubstringMatcher(substring string) SubstringMatcher {
	return &substringMatcher{substring: substring}
}

type substringMatcher struct {
	substring string
}

func (s *substringMatcher) Match(value any) MatchResult {
	strValue, ok := value.(string)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type string, but got %T", value),
		}
	}
	if !strings.Contains(strValue, s.substring) {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected string to contain '%s', but it did not", s.substring),
		}
	}
	return MatchResult{
		Matches: true,
	}
}
