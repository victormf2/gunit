package matchers

import (
	"fmt"
	"reflect"
)

type AnySliceMatcher struct {
	minLength       *int
	maxLength       *int
	elementMatchers []Matcher
	matchAll        bool
}

func (a *AnySliceMatcher) clone() *AnySliceMatcher {
	newMatcher := &AnySliceMatcher{
		minLength:       a.minLength,
		maxLength:       a.maxLength,
		elementMatchers: make([]Matcher, len(a.elementMatchers)),
		matchAll:        a.matchAll,
	}
	copy(newMatcher.elementMatchers, a.elementMatchers)
	return newMatcher
}

func (a *AnySliceMatcher) WithLength(length int) *AnySliceMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &length
	newMatcher.maxLength = &length
	return newMatcher
}

func (a *AnySliceMatcher) WithMaxLength(max int) *AnySliceMatcher {
	newMatcher := a.clone()
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *AnySliceMatcher) WithMinLength(min int) *AnySliceMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	return newMatcher
}

func (a *AnySliceMatcher) WithLengthBetween(min int, max int) *AnySliceMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *AnySliceMatcher) Containing(values ...any) *AnySliceMatcher {
	newMatcher := a.containing(false, values)
	return newMatcher
}

func (a *AnySliceMatcher) ContainingAll(values ...any) *AnySliceMatcher {
	newMatcher := a.containing(true, values)
	return newMatcher
}

func (a *AnySliceMatcher) containing(matchAll bool, values []any) *AnySliceMatcher {
	elementMatchers := []Matcher{}
	for _, value := range values {
		matcher := getSliceValueMatcher(value)
		elementMatchers = append(elementMatchers, matcher)
	}

	newMatcher := a.clone()
	newMatcher.elementMatchers = elementMatchers
	newMatcher.matchAll = matchAll
	return newMatcher
}

func (a *AnySliceMatcher) Match(value any) MatchResult {
	sliceValue := reflect.ValueOf(value)
	if sliceValue.Kind() != reflect.Slice && sliceValue.Kind() != reflect.Array {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type slice or array, but got %T", value),
		}
	}

	if a.minLength != nil && sliceValue.Len() < *a.minLength {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected slice length >= %d, but got %d", *a.minLength, sliceValue.Len()),
		}
	}
	if a.maxLength != nil && sliceValue.Len() > *a.maxLength {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected slice length <= %d, but got %d", *a.maxLength, sliceValue.Len()),
		}
	}
	foundMatch := false
	for _, elementMatcher := range a.elementMatchers {
		matched := matchElementInSlice(sliceValue, elementMatcher)
		if matched {
			foundMatch = true
			if !a.matchAll {
				break
			}
			continue
		}
		if a.matchAll {
			return MatchResult{
				Matches: false,
				Message: fmt.Sprintf("No element matched for matcher: %+v", elementMatcher),
			}
		}
	}
	if !foundMatch && len(a.elementMatchers) > 0 {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("None of the element matchers matched the slice"),
		}
	}
	return MatchResult{Matches: true}
}

func getSliceValueMatcher(value any) Matcher {
	switch v := value.(type) {
	case Matcher:
		return v
	default:
		return &EqualMatcher{Expected: v}
	}
}

func matchElementInSlice(sliceValue reflect.Value, elementMatcher Matcher) bool {
	matched := false
	for i := range sliceValue.Len() {
		element := sliceValue.Index(i).Interface()
		matchResult := elementMatcher.Match(element)
		if matchResult.Matches {
			matched = true
			break
		}
	}
	return matched
}
