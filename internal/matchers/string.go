package matchers

import (
	"fmt"
	"regexp"
)

type StringMatcher interface {
	Matcher
	WithLength(length int) StringMatcher
	WithMaxLength(max int) StringMatcher
	WithMinLength(min int) StringMatcher
	WithLengthBetween(min int, max int) StringMatcher
	MatchingAny(values ...any) StringMatcher
	MatchingAll(values ...any) StringMatcher
	ContainingAny(values ...any) StringMatcher
	ContainingAll(values ...any) StringMatcher
}

func NewStringMatcher() StringMatcher {
	return &stringMatcher{}
}

type stringMatcher struct {
	minLength        *int
	maxLength        *int
	expectedPatterns []Matcher
	matchAll         bool
}

func (a *stringMatcher) clone() *stringMatcher {
	newMatcher := &stringMatcher{
		minLength:        a.minLength,
		maxLength:        a.maxLength,
		expectedPatterns: make([]Matcher, len(a.expectedPatterns)),
		matchAll:         a.matchAll,
	}
	copy(newMatcher.expectedPatterns, a.expectedPatterns)
	return newMatcher
}

func (a *stringMatcher) WithLength(length int) StringMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &length
	newMatcher.maxLength = &length
	return newMatcher
}

func (a *stringMatcher) WithMaxLength(max int) StringMatcher {
	newMatcher := a.clone()
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *stringMatcher) WithMinLength(min int) StringMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	return newMatcher
}

func (a *stringMatcher) WithLengthBetween(min int, max int) StringMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *stringMatcher) MatchingAny(values ...any) StringMatcher {
	return a.matching(false, values...)
}

func (a *stringMatcher) MatchingAll(values ...any) StringMatcher {
	return a.matching(true, values...)
}

// Alias for Matching
func (a *stringMatcher) ContainingAny(values ...any) StringMatcher {
	return a.MatchingAny(values...)
}

// Alias for MatchingAll
func (a *stringMatcher) ContainingAll(values ...any) StringMatcher {
	return a.MatchingAll(values...)
}

func (a *stringMatcher) matching(matchAll bool, values ...any) *stringMatcher {
	newMatcher := a.clone()
	expectedPatterns := []Matcher{}
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

func (a *stringMatcher) Match(value any) MatchResult {
	strValue, ok := value.(string)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type string, but got %T", value),
		}
	}
	if a.minLength != nil && len(strValue) < *a.minLength {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected string length >= %d, but got %d", *a.minLength, len(strValue)),
		}
	}
	if a.maxLength != nil && len(strValue) > *a.maxLength {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected string length <= %d, but got %d", *a.maxLength, len(strValue)),
		}
	}
	foundMatch := false
	for _, stringMatcher := range a.expectedPatterns {
		matchResult := stringMatcher.Match(strValue)
		if matchResult.Matches {
			foundMatch = true
			if !a.matchAll {
				break
			}
			continue
		}
		if a.matchAll {
			return MatchResult{
				Matches: false,
				Message: matchResult.Message,
			}
		}
	}
	if !foundMatch && len(a.expectedPatterns) > 0 {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("None of the substring matchers matched the string"),
		}
	}
	return MatchResult{Matches: true}
}

func getPatternMatcher(value any) (Matcher, bool) {
	switch v := value.(type) {
	case string:
		return NewSubstringMatcher(v), true
	case *regexp.Regexp:
		return NewRegexMatcher(v), true
	case Matcher:
		return v, true
	default:
		return nil, false
	}
}
