package matchers

import (
	"fmt"
	"regexp"
)

type AnyStringMatcher struct {
	minLength      *int
	maxLength      *int
	stringMatchers []Matcher
	matchAll       bool
}

func (a *AnyStringMatcher) clone() *AnyStringMatcher {
	newMatcher := &AnyStringMatcher{
		minLength:      a.minLength,
		maxLength:      a.maxLength,
		stringMatchers: make([]Matcher, len(a.stringMatchers)),
		matchAll:       a.matchAll,
	}
	copy(newMatcher.stringMatchers, a.stringMatchers)
	return newMatcher
}

func (a *AnyStringMatcher) WithLength(length int) *AnyStringMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &length
	newMatcher.maxLength = &length
	return newMatcher
}

func (a *AnyStringMatcher) WithMaxLength(max int) *AnyStringMatcher {
	newMatcher := a.clone()
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *AnyStringMatcher) WithMinLength(min int) *AnyStringMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	return newMatcher
}

func (a *AnyStringMatcher) WithLengthBetween(min int, max int) *AnyStringMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *AnyStringMatcher) Matching(values ...any) *AnyStringMatcher {
	return a.matching(false, values...)
}

func (a *AnyStringMatcher) MatchingAll(values ...any) *AnyStringMatcher {
	return a.matching(true, values...)
}

// Alias for Matching
func (a *AnyStringMatcher) Containing(values ...any) *AnyStringMatcher {
	return a.Matching(values...)
}

// Alias for MatchingAll
func (a *AnyStringMatcher) ContainingAll(values ...any) *AnyStringMatcher {
	return a.MatchingAll(values...)
}

func (a *AnyStringMatcher) matching(matchAll bool, values ...any) *AnyStringMatcher {
	newMatcher := a.clone()
	stringMatchers := []Matcher{}
	for _, value := range values {
		matcher, ok := getStringMatcher(value)
		if !ok {
			panic("Containing only accepts strings, *regexp.Regexp, or Matcher instances")
		}
		stringMatchers = append(stringMatchers, matcher)
	}
	newMatcher.stringMatchers = stringMatchers
	newMatcher.matchAll = matchAll
	return newMatcher
}

func (a *AnyStringMatcher) Match(value any) MatchResult {
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
	for _, stringMatcher := range a.stringMatchers {
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
	if !foundMatch && len(a.stringMatchers) > 0 {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("None of the substring matchers matched the string"),
		}
	}
	return MatchResult{Matches: true}
}

func getStringMatcher(value any) (Matcher, bool) {
	switch v := value.(type) {
	case string:
		return &SubstringMatcher{Substring: v}, true
	case *regexp.Regexp:
		return &RegexMatcher{Regex: v}, true
	case Matcher:
		return v, true
	default:
		return nil, false
	}
}
