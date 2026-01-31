package matchers

import (
	"fmt"
	"reflect"

	"github.com/victormf2/gunit/internal/expect"
)

func NewMapMatcher() expect.MapMatcher {
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

func (a *mapMatcher) WithLength(length int) expect.MapMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &length
	newMatcher.maxLength = &length
	return newMatcher
}

func (a *mapMatcher) WithMaxLength(max int) expect.MapMatcher {
	newMatcher := a.clone()
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *mapMatcher) WithMinLength(min int) expect.MapMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	return newMatcher
}

func (a *mapMatcher) WithLengthBetween(min int, max int) expect.MapMatcher {
	newMatcher := a.clone()
	newMatcher.minLength = &min
	newMatcher.maxLength = &max
	return newMatcher
}

func (a *mapMatcher) ContainingAny(keyValues ...[]any) expect.MapMatcher {
	newMatcher := a.containing(false, keyValues)
	return newMatcher
}

func (a *mapMatcher) ContainingAll(keyValues ...[]any) expect.MapMatcher {
	newMatcher := a.containing(true, keyValues)
	return newMatcher
}

func (a *mapMatcher) ContainingAnyKeys(keys ...any) expect.MapMatcher {
	keyValues := [][]any{}
	for _, key := range keys {
		keyValues = append(keyValues, []any{key, NewAnyMatcher()})
	}

	newMatcher := a.containing(false, keyValues)
	return newMatcher
}

func (a *mapMatcher) ContainingAllKeys(keys ...any) expect.MapMatcher {
	keyValues := [][]any{}
	for _, key := range keys {
		keyValues = append(keyValues, []any{key, NewAnyMatcher()})
	}

	newMatcher := a.containing(true, keyValues)
	return newMatcher
}

func (a *mapMatcher) ContainingAnyValues(values ...any) expect.MapMatcher {
	keyValues := [][]any{}
	for _, value := range values {
		keyValues = append(keyValues, []any{NewAnyMatcher(), value})
	}

	newMatcher := a.containing(false, keyValues)
	return newMatcher
}

func (a *mapMatcher) ContainingAllValues(values ...any) expect.MapMatcher {
	keyValues := [][]any{}
	for _, value := range values {
		keyValues = append(keyValues, []any{NewAnyMatcher(), value})
	}

	newMatcher := a.containing(true, keyValues)
	return newMatcher
}

func (a *mapMatcher) matching(value any) expect.MapMatcher {
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

func (a *mapMatcher) containing(matchAll bool, keyValues [][]any) expect.MapMatcher {
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

func (a *mapMatcher) Match(value any) expect.MatchResult {
	if value == nil {
		if a.matchNil {
			return expect.Matches()
		}
		return expect.DoesNotMatch(fmt.Sprintf("Expected type map, but got nil"), nil)
	}
	mapValue, ok := getMap(value)
	if !ok {
		return expect.DoesNotMatch(fmt.Sprintf("Expected type map, but got %T", value), nil)
	}

	if a.minLength != nil && mapValue.Len() < *a.minLength {
		return expect.DoesNotMatch(fmt.Sprintf("Expected map length >= %d, but got %d", *a.minLength, mapValue.Len()), nil)
	}
	if a.maxLength != nil && mapValue.Len() > *a.maxLength {
		return expect.DoesNotMatch(fmt.Sprintf("Expected map length <= %d, but got %d", *a.maxLength, mapValue.Len()), nil)
	}

	if len(a.expectedKeyValues) == 0 {
		return expect.Matches()
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
			return expect.DoesNotMatch(fmt.Sprintf("No key-value pairs matched for matcher: %+v", kvMatcher), nil)
		}
	}
	if !foundMatch && len(a.expectedKeyValues) > 0 {
		return expect.DoesNotMatch(fmt.Sprintf("None of the key-value matchers matched the map"), nil)
	}
	return expect.Matches()
}

type keyValueMatcher struct {
	keyMatcher   expect.Matcher
	valueMatcher expect.Matcher
}

func (k *keyValueMatcher) match(expectedKey any, expectedValue any) expect.MatchResult {
	keyResult := k.keyMatcher.Match(expectedKey)
	if !keyResult.Matches() {
		return expect.DoesNotMatch(fmt.Sprintf("Key did not match: %s", keyResult.Message()), nil)
	}
	valueResult := k.valueMatcher.Match(expectedValue)
	if !valueResult.Matches() {
		return expect.DoesNotMatch(fmt.Sprintf("Value did not match: %s", valueResult.Message()), nil)
	}
	return expect.Matches()
}

func getMapKeyMatcher(expectedKey any) expect.Matcher {
	switch v := expectedKey.(type) {
	case expect.Matcher:
		return v
	default:
		return NewGeneralMatcher(expectedKey)
	}
}

func getMapValueMatcher(expectedValue any) expect.Matcher {
	switch v := expectedValue.(type) {
	case expect.Matcher:
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
		if !matchResult.Matches() {
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
