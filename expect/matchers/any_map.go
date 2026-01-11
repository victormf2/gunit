package matchers

import (
	"fmt"
	"reflect"
)

type AnyMapMatcher struct {
	MinLength        *int
	MaxLength        *int
	KeyValueMatchers []keyValueMatcher
}

func (a *AnyMapMatcher) clone() *AnyMapMatcher {
	newMatcher := &AnyMapMatcher{
		MinLength:        a.MinLength,
		MaxLength:        a.MaxLength,
		KeyValueMatchers: make([]keyValueMatcher, len(a.KeyValueMatchers)),
	}
	copy(newMatcher.KeyValueMatchers, a.KeyValueMatchers)
	return newMatcher
}

type keyValueMatcher struct {
	KeyMatcher   Matcher
	ValueMatcher Matcher
}

func (a *AnyMapMatcher) WithLength(length int) *AnyMapMatcher {
	newMatcher := a.clone()
	newMatcher.MinLength = &length
	newMatcher.MaxLength = &length
	return newMatcher
}

func (a *AnyMapMatcher) WithMaxLength(max int) *AnyMapMatcher {
	newMatcher := a.clone()
	newMatcher.MaxLength = &max
	return newMatcher
}

func (a *AnyMapMatcher) WithMinLength(min int) *AnyMapMatcher {
	newMatcher := a.clone()
	newMatcher.MinLength = &min
	return newMatcher
}

func (a *AnyMapMatcher) WithLengthBetween(min int, max int) *AnyMapMatcher {
	newMatcher := a.clone()
	newMatcher.MinLength = &min
	newMatcher.MaxLength = &max
	return newMatcher
}

func (a *AnyMapMatcher) Containing(keyValues ...[]any) *AnyMapMatcher {
	keyValueMatchers := []keyValueMatcher{}
	for _, pair := range keyValues {
		if len(pair) != 2 {
			panic("Each key-value pair must have exactly two elements: key and value")
		}
		keyMatcher, ok := pair[0].(Matcher)
		if !ok {
			keyMatcher = &EqualMatcher{Expected: pair[0]}
		}
		valueMatcher, ok := pair[1].(Matcher)
		if !ok {
			valueMatcher = &EqualMatcher{Expected: pair[1]}
		}
		keyValueMatchers = append(keyValueMatchers, keyValueMatcher{
			KeyMatcher:   keyMatcher,
			ValueMatcher: valueMatcher,
		})
	}
	a.KeyValueMatchers = keyValueMatchers
	return a
}

func (a *AnyMapMatcher) ContainingKeys(keys ...any) *AnyMapMatcher {
	for _, key := range keys {
		keyMatcher, ok := key.(Matcher)
		if !ok {
			keyMatcher = &EqualMatcher{Expected: key}
		}
		a.KeyValueMatchers = append(a.KeyValueMatchers, keyValueMatcher{
			KeyMatcher:   keyMatcher,
			ValueMatcher: &AnyMatcher{},
		})
	}
	return a
}

func (a *AnyMapMatcher) ContainingValues(values ...any) *AnyMapMatcher {
	for _, value := range values {
		valueMatcher, ok := value.(Matcher)
		if !ok {
			valueMatcher = &EqualMatcher{Expected: value}
		}
		a.KeyValueMatchers = append(a.KeyValueMatchers, keyValueMatcher{
			KeyMatcher:   &AnyMatcher{},
			ValueMatcher: valueMatcher,
		})
	}
	return a
}

func (a *AnyMapMatcher) Match(value any) MatchResult {
	mapValue := reflect.ValueOf(value)
	if mapValue.Kind() != reflect.Map {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type map, but got %T", value),
		}
	}
	if a.MinLength != nil && mapValue.Len() < *a.MinLength {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected map length >= %d, but got %d", *a.MinLength, mapValue.Len()),
		}
	}
	if a.MaxLength != nil && mapValue.Len() > *a.MaxLength {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected map length <= %d, but got %d", *a.MaxLength, mapValue.Len()),
		}
	}
	for _, kvMatcher := range a.KeyValueMatchers {
		matched := false
		for _, key := range mapValue.MapKeys() {
			keyInterface := key.Interface()
			valueInterface := mapValue.MapIndex(key).Interface()
			keyMatchResult := kvMatcher.KeyMatcher.Match(keyInterface)
			if !keyMatchResult.Matches {
				continue
			}
			valueMatchResult := kvMatcher.ValueMatcher.Match(valueInterface)
			if !valueMatchResult.Matches {
				continue
			}
			matched = true
			break
		}
		if !matched {
			return MatchResult{
				Matches: false,
				Message: fmt.Sprintf("No key-value pairs matched for matcher: %+v", kvMatcher),
			}
		}
	}
	return MatchResult{Matches: true}
}
