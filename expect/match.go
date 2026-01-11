package expect

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type MatchResult struct {
	Matches bool
	Message string
}

func (e *Expector) ToMatch(t *testing.T, expected any) {
	actual := e.value

	actualValue := reflect.ValueOf(actual)
	expectedValue := reflect.ValueOf(expected)

	matchResult := match(actualValue, expectedValue)
	if !matchResult.Matches {
		t.Fatal(matchResult.Message)
	}
}

func match(actualValue, expectedValue reflect.Value) MatchResult {
	// Handle nil pointers
	if isNil(actualValue) {
		if isNil(expectedValue) {
			return MatchResult{Matches: true} // Both nil, match
		}
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected %v, but got nil", expectedValue.Interface()),
		}
	}
	if isNil(expectedValue) {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected nil, but got %v", actualValue.Interface()),
		}
	}

	// Dereference pointers
	if actualValue.Kind() == reflect.Ptr {
		actualValue = actualValue.Elem()
	}
	if expectedValue.Kind() == reflect.Ptr {
		expectedValue = expectedValue.Elem()
	}

	if actualValue.Kind() != expectedValue.Kind() {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Type mismatch: expected %s, got %s", expectedValue.Kind(), actualValue.Kind()),
		}
	}

	switch actualValue.Kind() {
	case reflect.Struct:
		return matchStructs(actualValue, expectedValue)
	case reflect.Map:
		return matchMaps(actualValue, expectedValue)
	case reflect.Array, reflect.Slice:
		return matchSlices(actualValue, expectedValue)
	default:
		return matchEquals(actualValue, expectedValue)
	}
}

func matchStructs(actualStruct, expectedStruct reflect.Value) MatchResult {
	problemsByFieldName := map[string]string{}
	// Compare field by field
	for i := range expectedStruct.NumField() {
		expectedField := expectedStruct.Type().Field(i)

		// Skip private (unexported) fields
		if !expectedField.IsExported() {
			continue
		}

		expectedFieldValue := expectedStruct.Field(i)

		gunitTag := expectedField.Tag.Get("gunit")

		// Skip zero values in expected
		if expectedFieldValue.IsZero() && !strings.Contains(gunitTag, "required") {
			continue
		}

		fieldName := expectedField.Name
		actualFieldValue := actualStruct.FieldByName(fieldName)
		if !actualFieldValue.IsValid() {
			problemsByFieldName[fieldName] = "Field not found in actual value"
			continue
		}
		fieldMatchResult := match(actualFieldValue, expectedFieldValue)
		if !fieldMatchResult.Matches {
			problemsByFieldName[fieldName] = fieldMatchResult.Message
		}
	}
	if len(problemsByFieldName) > 0 {
		message := "Struct fields did not match:\n"
		for fieldName, problem := range problemsByFieldName {
			message += fmt.Sprintf("  %s: %s\n", fieldName, problem)
		}
		return MatchResult{
			Matches: false,
			Message: message,
		}
	}
	return MatchResult{Matches: true}
}

func matchMaps(actualMap, expectedMap reflect.Value) MatchResult {
	problemsByKey := map[string]string{}
	for _, key := range expectedMap.MapKeys() {
		expectedValue := expectedMap.MapIndex(key)
		actualValue := actualMap.MapIndex(key)
		if !actualValue.IsValid() {
			problemsByKey[key.String()] = "Key not found in actual map"
			continue
		}
		keyMatchResult := match(actualValue, expectedValue)
		if !keyMatchResult.Matches {
			problemsByKey[key.String()] = keyMatchResult.Message
		}
	}
	if len(problemsByKey) > 0 {
		message := "Map keys did not match:\n"
		for key, problem := range problemsByKey {
			message += fmt.Sprintf("  %s: %s\n", key, problem)
		}
		return MatchResult{
			Matches: false,
			Message: message,
		}
	}
	return MatchResult{Matches: true}
}

func matchSlices(actualSlice, expectedSlice reflect.Value) MatchResult {
	if actualSlice.Len() != expectedSlice.Len() {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Length mismatch: expected %d, got %d", expectedSlice.Len(), actualSlice.Len()),
		}
	}

	problemsByIndex := map[int]string{}
	for i := 0; i < actualSlice.Len(); i++ {
		elementMatchResult := match(actualSlice.Index(i), expectedSlice.Index(i))
		if !elementMatchResult.Matches {
			problemsByIndex[i] = elementMatchResult.Message
		}
	}
	if len(problemsByIndex) > 0 {
		message := "Slice elements did not match:\n"
		for index, problem := range problemsByIndex {
			message += fmt.Sprintf("  %d: %s\n", index, problem)
		}
		return MatchResult{
			Matches: false,
			Message: message,
		}
	}
	return MatchResult{Matches: true}
}

func matchEquals(actualValue, expectedValue reflect.Value) MatchResult {
	if equals(actualValue, expectedValue) {
		return MatchResult{Matches: true}
	}
	return MatchResult{
		Matches: false,
		Message: fmt.Sprintf("Expected %v, got %v", expectedValue.Interface(), actualValue.Interface()),
	}
}

func equals(actualValue, expectedValue reflect.Value) bool {
	return reflect.DeepEqual(actualValue.Interface(), expectedValue.Interface())
}

func isNil(value reflect.Value) bool {
	if !value.IsValid() {
		return true
	}
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		if value.IsNil() {
			return true
		}
	}
	return false
}
