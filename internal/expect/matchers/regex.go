package matchers

import (
	"fmt"
	"regexp"

	"github.com/victormf2/gunit/internal/expect"
)

type RegexMatcher interface {
	expect.Matcher
}

func NewRegexMatcher(regex *regexp.Regexp) RegexMatcher {
	return &regexMatcher{regex: regex}
}

type regexMatcher struct {
	regex *regexp.Regexp
}

func (r *regexMatcher) Match(value any) expect.MatchResult {
	strValue, ok := value.(string)
	if !ok {
		return expect.DoesNotMatch(fmt.Sprintf("Expected type string, but got %T", value), nil)
	}

	if !r.regex.MatchString(strValue) {
		return expect.DoesNotMatch(fmt.Sprintf("Expected string to match regex '%s', but it did not", r.regex), nil)
	}

	return expect.Matches()
}
