package matchers

import (
	"fmt"
	"strings"
)

type AnyStringMatcher struct {
	MinLength  *int
	MaxLength  *int
	Substrings []string
}

func (a *AnyStringMatcher) WithLength(length int) *AnyStringMatcher {
	*a.MinLength = length
	*a.MaxLength = length
	return a
}

func (a *AnyStringMatcher) WithMaxLength(max int) *AnyStringMatcher {
	*a.MaxLength = max
	return a
}

func (a *AnyStringMatcher) WithMinLength(min int) *AnyStringMatcher {
	*a.MinLength = min
	return a
}

func (a *AnyStringMatcher) WithLengthBetween(min int, max int) *AnyStringMatcher {
	*a.MinLength = min
	*a.MaxLength = max
	return a
}

func (a *AnyStringMatcher) Containing(substrings ...string) *AnyStringMatcher {
	a.Substrings = substrings
	return a
}

func (a *AnyStringMatcher) Match(value any) MatchResult {
	strValue, ok := value.(string)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type string, but got %T", value),
		}
	}
	if a.MinLength != nil && len(strValue) < *a.MinLength {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected string length >= %d, but got %d", *a.MinLength, len(strValue)),
		}
	}
	if a.MaxLength != nil && len(strValue) > *a.MaxLength {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected string length <= %d, but got %d", *a.MaxLength, len(strValue)),
		}
	}
	for _, substring := range a.Substrings {
		if !strings.Contains(strValue, substring) {
			return MatchResult{
				Matches: false,
				Message: fmt.Sprintf("Expected string to contain '%s', but it did not", substring),
			}
		}
	}
	return MatchResult{Matches: true}
}
