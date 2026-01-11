package matchers

import (
	"fmt"
	"regexp"
)

type RegexMatcher struct {
	Regex *regexp.Regexp
}

func (r *RegexMatcher) Match(value any) MatchResult {
	strValue, ok := value.(string)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type string, but got %T", value),
		}
	}

	if !r.Regex.MatchString(strValue) {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected string to match regex '%s', but it did not", r.Regex),
		}
	}

	return MatchResult{
		Matches: true,
	}
}
