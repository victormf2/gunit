package matchers

import (
	"fmt"
	"reflect"
)

type AnySliceMatcher struct {
	MinLength       *int
	MaxLength       *int
	ElementMatchers []Matcher
}

func (a *AnySliceMatcher) clone() *AnySliceMatcher {
	newMatcher := &AnySliceMatcher{
		MinLength:       a.MinLength,
		MaxLength:       a.MaxLength,
		ElementMatchers: make([]Matcher, len(a.ElementMatchers)),
	}
	copy(newMatcher.ElementMatchers, a.ElementMatchers)
	return newMatcher
}

func (a *AnySliceMatcher) WithLength(length int) *AnySliceMatcher {
	newMatcher := a.clone()
	newMatcher.MinLength = &length
	newMatcher.MaxLength = &length
	return newMatcher
}

func (a *AnySliceMatcher) WithMaxLength(max int) *AnySliceMatcher {
	newMatcher := a.clone()
	newMatcher.MaxLength = &max
	return newMatcher
}

func (a *AnySliceMatcher) WithMinLength(min int) *AnySliceMatcher {
	newMatcher := a.clone()
	newMatcher.MinLength = &min
	return newMatcher
}

func (a *AnySliceMatcher) WithLengthBetween(min int, max int) *AnySliceMatcher {
	newMatcher := a.clone()
	newMatcher.MinLength = &min
	newMatcher.MaxLength = &max
	return newMatcher
}

func (a *AnySliceMatcher) Containing(items ...any) *AnySliceMatcher {
	for _, item := range items {
		matcher, ok := item.(Matcher)
		if ok {
			a.ElementMatchers = append(a.ElementMatchers, matcher)
			continue
		}
		a.ElementMatchers = append(a.ElementMatchers, &EqualMatcher{Expected: item})
	}
	return a
}

func (a *AnySliceMatcher) Match(value any) MatchResult {
	sliceValue := reflect.ValueOf(value)
	if sliceValue.Kind() != reflect.Slice && sliceValue.Kind() != reflect.Array {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type slice or array, but got %T", value),
		}
	}

	if a.MinLength != nil && sliceValue.Len() < *a.MinLength {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected slice length >= %d, but got %d", *a.MinLength, sliceValue.Len()),
		}
	}
	if a.MaxLength != nil && sliceValue.Len() > *a.MaxLength {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected slice length <= %d, but got %d", *a.MaxLength, sliceValue.Len()),
		}
	}
	for _, matcher := range a.ElementMatchers {
		matched := false
		for i := range sliceValue.Len() {
			element := sliceValue.Index(i).Interface()
			matchResult := matcher.Match(element)
			if matchResult.Matches {
				matched = true
				break
			}
		}
		if !matched {
			return MatchResult{
				Matches: false,
				Message: fmt.Sprintf("No elements matched for matcher: %+v", matcher),
			}
		}
	}
	return MatchResult{Matches: true}
}
