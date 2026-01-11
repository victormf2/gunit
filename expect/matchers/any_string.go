package matchers

import (
	"fmt"
)

type AnyStringMatcher struct {
	MinLength         *int
	MaxLength         *int
	SubstringMatchers []Matcher
	MatchAll          bool
}

func (a *AnyStringMatcher) clone() *AnyStringMatcher {
	newMatcher := &AnyStringMatcher{
		MinLength:         a.MinLength,
		MaxLength:         a.MaxLength,
		SubstringMatchers: make([]Matcher, len(a.SubstringMatchers)),
		MatchAll:          a.MatchAll,
	}
	copy(newMatcher.SubstringMatchers, a.SubstringMatchers)
	return newMatcher
}

func (a *AnyStringMatcher) WithLength(length int) *AnyStringMatcher {
	newMatcher := a.clone()
	newMatcher.MinLength = &length
	newMatcher.MaxLength = &length
	return newMatcher
}

func (a *AnyStringMatcher) WithMaxLength(max int) *AnyStringMatcher {
	newMatcher := a.clone()
	newMatcher.MaxLength = &max
	return newMatcher
}

func (a *AnyStringMatcher) WithMinLength(min int) *AnyStringMatcher {
	newMatcher := a.clone()
	newMatcher.MinLength = &min
	return newMatcher
}

func (a *AnyStringMatcher) WithLengthBetween(min int, max int) *AnyStringMatcher {
	newMatcher := a.clone()
	newMatcher.MinLength = &min
	newMatcher.MaxLength = &max
	return newMatcher
}

func (a *AnyStringMatcher) ContainingAll(values ...any) *AnyStringMatcher {
	newMatcher := a.clone()
	substringMatchers := []Matcher{}
	for _, value := range values {
		matcher, ok := value.(Matcher)
		if !ok {
			substring, ok := value.(string)
			if !ok {
				panic("ContainingAll only accepts strings or Matcher instances")
			}
			matcher = &SubstringMatcher{Substring: substring}
		}
		substringMatchers = append(substringMatchers, matcher)
	}
	newMatcher.SubstringMatchers = substringMatchers
	newMatcher.MatchAll = true
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
	foundMatch := false
	for _, substringMatcher := range a.SubstringMatchers {
		matchResult := substringMatcher.Match(strValue)
		if matchResult.Matches {
			foundMatch = true
			continue
		}
		if a.MatchAll {
			return MatchResult{
				Matches: false,
				Message: matchResult.Message,
			}
		}
	}
	if !foundMatch && len(a.SubstringMatchers) > 0 {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("None of the substring matchers matched the string"),
		}
	}
	return MatchResult{Matches: true}
}
