package matchers

import (
	"fmt"
	"reflect"
)

type AnySliceMatcher struct {
	minLength        *int
	maxLength        *int
	matchNil         bool
	matchInOrder     bool
	matchAll         bool
	expectedElements []Matcher
}

func (a *AnySliceMatcher) clone() *AnySliceMatcher {
	newMatcher := &AnySliceMatcher{
		minLength:        a.minLength,
		maxLength:        a.maxLength,
		matchInOrder:     a.matchInOrder,
		matchAll:         a.matchAll,
		matchNil:         a.matchNil,
		expectedElements: make([]Matcher, len(a.expectedElements)),
	}
	copy(newMatcher.expectedElements, a.expectedElements)
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
	newMatcher := a.containing(false, false, values)
	return newMatcher
}

func (a *AnySliceMatcher) ContainingAll(values ...any) *AnySliceMatcher {
	newMatcher := a.containing(true, false, values)
	return newMatcher
}

func (a *AnySliceMatcher) matching(value any) *AnySliceMatcher {
	if value == nil {
		newMatcher := a.clone()
		newMatcher.matchNil = true
		return newMatcher
	}
	sliceValue, ok := getSlice(value)
	if !ok {
		panic("Matching only accept nil, arrays and alices")
	}
	elements := make([]any, sliceValue.Len())
	for i := range sliceValue.Len() {
		elements[i] = sliceValue.Index(i).Interface()
	}
	newMatcher := a.WithLength(sliceValue.Len()).containing(true, true, elements)
	return newMatcher
}

func (a *AnySliceMatcher) containing(matchAll bool, matchInOrder bool, expectedElements []any) *AnySliceMatcher {
	elementMatchers := []Matcher{}
	for _, element := range expectedElements {
		matcher := getSliceElementMatcher(element)
		elementMatchers = append(elementMatchers, matcher)
	}

	newMatcher := a.clone()
	newMatcher.expectedElements = elementMatchers
	newMatcher.matchAll = matchAll
	newMatcher.matchInOrder = matchInOrder
	return newMatcher
}

func (a *AnySliceMatcher) Match(value any) MatchResult {
	if value == nil {
		if a.matchNil {
			return MatchResult{Matches: true}
		}
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type slice or array, but got nil"),
		}
	}
	sliceValue, ok := getSlice(value)
	if !ok {
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

	if len(a.expectedElements) == 0 {
		return MatchResult{Matches: true}
	}

	foundMatch := false
	problemsByIndex := map[any]MatchResult{}
	for elementIndex, elementMatcher := range a.expectedElements {
		elementMatchResult := a.matchElementInSlice(sliceValue, elementIndex, elementMatcher)
		if elementMatchResult.Matches {
			foundMatch = true
			if !a.matchAll {
				break
			}
			continue
		}
		problemsByIndex[elementIndex] = elementMatchResult
	}

	if !a.matchAll {
		if !foundMatch {
			return MatchResult{
				Matches:         false,
				Message:         "no matches found in slice",
				ProblemsByField: problemsByIndex,
			}
		}
		return MatchResult{Matches: true}
	}

	if len(problemsByIndex) > 0 {
		return MatchResult{
			Matches:         false,
			Message:         "some elements did not match",
			ProblemsByField: problemsByIndex,
		}
	}

	return MatchResult{Matches: true}
}

func getSliceElementMatcher(value any) Matcher {
	switch v := value.(type) {
	case Matcher:
		return v
	default:
		return (&GeneralMatcher{}).Matching(value)
	}
}

func (a *AnySliceMatcher) matchElementInSlice(sliceValue reflect.Value, elementIndex int, elementMatcher Matcher) MatchResult {
	if a.matchInOrder {
		element := sliceValue.Index(elementIndex).Interface()
		matchResult := elementMatcher.Match(element)
		return matchResult
	}
	matched := false
	for i := range sliceValue.Len() {
		element := sliceValue.Index(i).Interface()
		matchResult := elementMatcher.Match(element)
		if matchResult.Matches {
			matched = true
			break
		}
	}
	if !matched {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("No element found in slice for matcher %+v", elementMatcher),
		}
	}
	return MatchResult{Matches: true}
}

func getSlice(value any) (reflect.Value, bool) {
	var zeroValue reflect.Value
	if value == nil {
		return zeroValue, false
	}
	sliceValue := reflect.ValueOf(value)
	if sliceValue.Kind() != reflect.Array && sliceValue.Kind() != reflect.Slice {
		return zeroValue, false
	}
	return sliceValue, true
}
