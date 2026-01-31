package matchers

import (
	"fmt"
	"regexp"

	"github.com/victormf2/gunit/internal/expect"
)

func NewStringMatcher() expect.StringMatcher {
	return &stringMatcher{}
}

type stringMatcher struct {
	minLength        *int
	maxLength        *int
	expectedPatterns []expect.Matcher
	matchAll         bool
}

func (a *stringMatcher) clone() *stringMatcher {
	newMatcher := &stringMatcher{
		minLength:        a.minLength,
		maxLength:        a.maxLength,
		expectedPatterns: make([]expect.Matcher, len(a.expectedPatterns)),
		matchAll:         a.matchAll,
	}
	copy(newMatcher.expectedPatterns, a.expectedPatterns)
	return newMatcher
}

func (a *stringMatcher) WithLength(length int) expect.StringMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &length
	newMatcher.maxLength = &length
	return newMatcher
}

func (a *stringMatcher) WithMaxLength(max int) expect.StringMatcher {
	newMatcher := a.clone()
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *stringMatcher) WithMinLength(min int) expect.StringMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	return newMatcher
}

func (a *stringMatcher) WithLengthBetween(min int, max int) expect.StringMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *stringMatcher) MatchingAny(values ...any) expect.StringMatcher {
	return a.matching(false, values...)
}

func (a *stringMatcher) MatchingAll(values ...any) expect.StringMatcher {
	return a.matching(true, values...)
}

// Alias for Matching
func (a *stringMatcher) ContainingAny(values ...any) expect.StringMatcher {
	return a.MatchingAny(values...)
}

// Alias for MatchingAll
func (a *stringMatcher) ContainingAll(values ...any) expect.StringMatcher {
	return a.MatchingAll(values...)
}

func (a *stringMatcher) matching(matchAll bool, values ...any) *stringMatcher {
	newMatcher := a.clone()
	expectedPatterns := []expect.Matcher{}
	for _, value := range values {
		matcher, ok := getPatternMatcher(value)
		if !ok {
			panic("Containing only accepts strings, *regexp.Regexp, or Matcher instances")
		}
		expectedPatterns = append(expectedPatterns, matcher)
	}
	newMatcher.expectedPatterns = expectedPatterns
	newMatcher.matchAll = matchAll
	return newMatcher
}

func (a *stringMatcher) Match(value any) expect.MatchResult {
	strValue, ok := value.(string)
	if !ok {
		return expect.DoesNotMatch(fmt.Sprintf("Expected type string, but got %T", value), nil)
	}
	if a.minLength != nil && len(strValue) < *a.minLength {
		return expect.DoesNotMatch(fmt.Sprintf("Expected string length >= %d, but got %d", *a.minLength, len(strValue)), nil)
	}
	if a.maxLength != nil && len(strValue) > *a.maxLength {
		return expect.DoesNotMatch(fmt.Sprintf("Expected string length <= %d, but got %d", *a.maxLength, len(strValue)), nil)
	}
	foundMatch := false
	for _, stringMatcher := range a.expectedPatterns {
		matchResult := stringMatcher.Match(strValue)
		if matchResult.Matches() {
			foundMatch = true
			if !a.matchAll {
				break
			}
			continue
		}
		if a.matchAll {
			return expect.DoesNotMatch(matchResult.Message(), nil)
		}
	}
	if !foundMatch && len(a.expectedPatterns) > 0 {
		return expect.DoesNotMatch(fmt.Sprintf("None of the substring matchers matched the string"), nil)
	}
	return expect.Matches()
}

func getPatternMatcher(value any) (expect.Matcher, bool) {
	switch v := value.(type) {
	case string:
		return NewSubstringMatcher(v), true
	case *regexp.Regexp:
		return NewRegexMatcher(v), true
	case expect.Matcher:
		return v, true
	default:
		return nil, false
	}
}
