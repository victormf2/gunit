package matchers

import (
	"fmt"
	"regexp"
)

type RegexMatcher interface {
	Matcher
}

func NewRegexMatcher(regex *regexp.Regexp) RegexMatcher {
	return &regexMatcher{regex: regex}
}

type regexMatcher struct {
	regex *regexp.Regexp
}

func (r *regexMatcher) Match(value any) MatchResult {
	strValue, ok := value.(string)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type string, but got %T", value),
		}
	}

	if !r.regex.MatchString(strValue) {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected string to match regex '%s', but it did not", r.regex),
		}
	}

	return MatchResult{
		Matches: true,
	}
}
