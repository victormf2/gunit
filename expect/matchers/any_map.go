package matchers

import (
	"fmt"
	"reflect"
)

type AnyMapMatcher struct {
	minLength         *int
	maxLength         *int
	matchAll          bool
	matchNil          bool
	expectedKeyValues []*keyValueMatcher
}

func (a *AnyMapMatcher) clone() *AnyMapMatcher {
	newMatcher := &AnyMapMatcher{
		minLength:         a.minLength,
		maxLength:         a.maxLength,
		matchAll:          a.matchAll,
		matchNil:          a.matchNil,
		expectedKeyValues: make([]*keyValueMatcher, len(a.expectedKeyValues)),
	}
	copy(newMatcher.expectedKeyValues, a.expectedKeyValues)
	return newMatcher
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

func (a *AnyMapMatcher) WithLength(length int) *AnyMapMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &length
	newMatcher.maxLength = &length
	return newMatcher
}

func (a *AnyMapMatcher) WithMaxLength(max int) *AnyMapMatcher {
	newMatcher := a.clone()
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *AnyMapMatcher) WithMinLength(min int) *AnyMapMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	return newMatcher
}

func (a *AnyMapMatcher) WithLengthBetween(min int, max int) *AnyMapMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *AnyMapMatcher) Containing(keyValues ...[]any) *AnyMapMatcher {
	newMatcher := a.containing(false, keyValues)
	return newMatcher
}

func (a *AnyMapMatcher) ContainingAll(keyValues ...[]any) *AnyMapMatcher {
	newMatcher := a.containing(true, keyValues)
	return newMatcher
}

func (a *AnyMapMatcher) ContainingKeys(keys ...any) *AnyMapMatcher {
	keyValues := [][]any{}
	for _, key := range keys {
		keyValues = append(keyValues, []any{key, &AnyMatcher{}})
	}

	newMatcher := a.containing(false, keyValues)
	return newMatcher
}

func (a *AnyMapMatcher) ContainingAllKeys(keys ...any) *AnyMapMatcher {
	keyValues := [][]any{}
	for _, key := range keys {
		keyValues = append(keyValues, []any{key, &AnyMatcher{}})
	}

	newMatcher := a.containing(true, keyValues)
	return newMatcher
}

func (a *AnyMapMatcher) ContainingValues(values ...any) *AnyMapMatcher {
	keyValues := [][]any{}
	for _, value := range values {
		keyValues = append(keyValues, []any{&AnyMatcher{}, value})
	}

	newMatcher := a.containing(false, keyValues)
	return newMatcher
}

func (a *AnyMapMatcher) ContainingAllValues(values ...any) *AnyMapMatcher {
	keyValues := [][]any{}
	for _, value := range values {
		keyValues = append(keyValues, []any{&AnyMatcher{}, value})
	}

	newMatcher := a.containing(true, keyValues)
	return newMatcher
}

func (a *AnyMapMatcher) matching(value any) *AnyMapMatcher {
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
	newMatcher := a.WithLength(mapValue.Len()).containing(true, keyValues)
	return newMatcher
}

func (a *AnyMapMatcher) containing(matchAll bool, keyValues [][]any) *AnyMapMatcher {
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

func (a *AnyMapMatcher) Match(value any) MatchResult {
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

func getMapKeyMatcher(expectedKey any) Matcher {
	switch v := expectedKey.(type) {
	case Matcher:
		return v
	default:
		return (&GeneralMatcher{}).Matching(expectedKey)
	}
}

func getMapValueMatcher(expectedValue any) Matcher {
	switch v := expectedValue.(type) {
	case Matcher:
		return v
	default:
		return (&GeneralMatcher{}).Matching(expectedValue)
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
