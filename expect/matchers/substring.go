package matchers

import (
	"fmt"
	"strings"
)

type SubstringMatcher struct {
	Substring string
}

func (s *SubstringMatcher) Match(value any) MatchResult {
	strValue, ok := value.(string)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type string, but got %T", value),
		}
	}
	if !strings.Contains(strValue, s.Substring) {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected string to contain '%s', but it did not", s.Substring),
		}
	}
	return MatchResult{
		Matches: true,
	}
}
