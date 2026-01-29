package matchers

import (
	"fmt"
	"reflect"
)

type MapMatcher interface {
	Matcher
	WithLength(length int) MapMatcher
	WithMaxLength(max int) MapMatcher
	WithMinLength(min int) MapMatcher
	WithLengthBetween(min int, max int) MapMatcher
	Containing(keyValues ...[]any) MapMatcher
	ContainingAll(keyValues ...[]any) MapMatcher
	ContainingKeys(keys ...any) MapMatcher
	ContainingAllKeys(keys ...any) MapMatcher
	ContainingValues(values ...any) MapMatcher
	ContainingAllValues(values ...any) MapMatcher
}

func NewMapMatcher() MapMatcher {
	return &mapMatcher{}
}

type mapMatcher struct {
	minLength         *int
	maxLength         *int
	matchAll          bool
	matchNil          bool
	expectedKeyValues []*keyValueMatcher
}

func (a *mapMatcher) clone() *mapMatcher {
	newMatcher := &mapMatcher{
		minLength:         a.minLength,
		maxLength:         a.maxLength,
		matchAll:          a.matchAll,
		matchNil:          a.matchNil,
		expectedKeyValues: make([]*keyValueMatcher, len(a.expectedKeyValues)),
	}
	copy(newMatcher.expectedKeyValues, a.expectedKeyValues)
	return newMatcher
}

func (a *mapMatcher) WithLength(length int) MapMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &length
	newMatcher.maxLength = &length
	return newMatcher
}

func (a *mapMatcher) WithMaxLength(max int) MapMatcher {
	newMatcher := a.clone()
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *mapMatcher) WithMinLength(min int) MapMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	return newMatcher
}

func (a *mapMatcher) WithLengthBetween(min int, max int) MapMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *mapMatcher) Containing(keyValues ...[]any) MapMatcher {
	newMatcher := a.containing(false, keyValues)
	return newMatcher
}

func (a *mapMatcher) ContainingAll(keyValues ...[]any) MapMatcher {
	newMatcher := a.containing(true, keyValues)
	return newMatcher
}

func (a *mapMatcher) ContainingKeys(keys ...any) MapMatcher {
	keyValues := [][]any{}
	for _, key := range keys {
		keyValues = append(keyValues, []any{key, NewAnyMatcher()})
	}

	newMatcher := a.containing(false, keyValues)
	return newMatcher
}

func (a *mapMatcher) ContainingAllKeys(keys ...any) MapMatcher {
	keyValues := [][]any{}
	for _, key := range keys {
		keyValues = append(keyValues, []any{key, NewAnyMatcher()})
	}

	newMatcher := a.containing(true, keyValues)
	return newMatcher
}

func (a *mapMatcher) ContainingValues(values ...any) MapMatcher {
	keyValues := [][]any{}
	for _, value := range values {
		keyValues = append(keyValues, []any{NewAnyMatcher(), value})
	}

	newMatcher := a.containing(false, keyValues)
	return newMatcher
}

func (a *mapMatcher) ContainingAllValues(values ...any) MapMatcher {
	keyValues := [][]any{}
	for _, value := range values {
		keyValues = append(keyValues, []any{NewAnyMatcher(), value})
	}

	newMatcher := a.containing(true, keyValues)
	return newMatcher
}

func (a *mapMatcher) matching(value any) MapMatcher {
	if value == nil {
		newMatcher := a.clone()
		newMatcher.matchNil = true
		return newMatcher
	}
	mapValue, ok := getMap(value)
	if !ok {
		panic("Matching only accepts nil or maps")
	}

	keys := mapValue.MapKeys()
	keyValues := [][]any{}
	for _, key := range keys {
		value := mapValue.MapIndex(key)
		keyValues = append(keyValues, []any{key.Interface(), value.Interface()})
	}
	newMatcher := a.WithMinLength(mapValue.Len()).(*mapMatcher).containing(true, keyValues)
	return newMatcher
}

func (a *mapMatcher) containing(matchAll bool, keyValues [][]any) MapMatcher {
	keyValueMatchers := []*keyValueMatcher{}
	for _, pair := range keyValues {
		if len(pair) != 2 {
			panic("Each key-value pair must have exactly two elements: key and value")
		}
		keyMatcher := getMapKeyMatcher(pair[0])
		valueMatcher := getMapValueMatcher(pair[1])
		keyValueMatchers = append(keyValueMatchers, &keyValueMatcher{
			keyMatcher:   keyMatcher,
			valueMatcher: valueMatcher,
		})
	}

	newMatcher := a.clone()
	newMatcher.expectedKeyValues = keyValueMatchers
	newMatcher.matchAll = matchAll
	return newMatcher
}

func (a *mapMatcher) Match(value any) MatchResult {
	if value == nil {
		if a.matchNil {
			return MatchResult{Matches: true}
		}
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type map, but got nil"),
		}
	}
	mapValue, ok := getMap(value)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type map, but got %T", value),
		}
	}

	if a.minLength != nil && mapValue.Len() < *a.minLength {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected map length >= %d, but got %d", *a.minLength, mapValue.Len()),
		}
	}
	if a.maxLength != nil && mapValue.Len() > *a.maxLength {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected map length <= %d, but got %d", *a.maxLength, mapValue.Len()),
		}
	}

	if len(a.expectedKeyValues) == 0 {
		return MatchResult{Matches: true}
	}

	foundMatch := false
	for _, kvMatcher := range a.expectedKeyValues {
		matched := matchKeyValueInMap(mapValue, kvMatcher)
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
				Message: fmt.Sprintf("No key-value pairs matched for matcher: %+v", kvMatcher),
			}
		}
	}
	if !foundMatch && len(a.expectedKeyValues) > 0 {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("None of the key-value matchers matched the map"),
		}
	}
	return MatchResult{Matches: true}
}

type keyValueMatcher struct {
	keyMatcher   Matcher
	valueMatcher Matcher
}

func (k *keyValueMatcher) match(expectedKey any, expectedValue any) MatchResult {
	keyResult := k.keyMatcher.Match(expectedKey)
	if !keyResult.Matches {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Key did not match: %s", keyResult.Message),
		}
	}
	valueResult := k.valueMatcher.Match(expectedValue)
	if !valueResult.Matches {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Value did not match: %s", valueResult.Message),
		}
	}
	return MatchResult{Matches: true}
}

func getMapKeyMatcher(expectedKey any) Matcher {
	switch v := expectedKey.(type) {
	case Matcher:
		return v
	default:
		return NewGeneralMatcher(expectedKey)
	}
}

func getMapValueMatcher(expectedValue any) Matcher {
	switch v := expectedValue.(type) {
	case Matcher:
		return v
	default:
		return NewGeneralMatcher(expectedValue)
	}
}

func matchKeyValueInMap(mapValue reflect.Value, kvMatcher *keyValueMatcher) bool {
	matched := false
	for _, key := range mapValue.MapKeys() {
		keyInterface := key.Interface()
		valueInterface := mapValue.MapIndex(key).Interface()
		matchResult := kvMatcher.match(keyInterface, valueInterface)
		if !matchResult.Matches {
			continue
		}
		matched = true
		break
	}
	return matched
}

func getMap(value any) (reflect.Value, bool) {
	var zeroValue reflect.Value
	if value == nil {
		return zeroValue, false
	}
	mapValue := reflect.ValueOf(value)
	if mapValue.Kind() != reflect.Map {
		return zeroValue, false
	}
	return mapValue, true
}
