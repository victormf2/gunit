package matchers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/victormf2/gunit/internal/expect"
)

func NewStructMatcher() expect.StructMatcher {
	return &structMatcher{}
}

type structMatcher struct {
	expectedValue reflect.Value
}

func (a *structMatcher) clone() *structMatcher {
	newMatcher := &structMatcher{
		expectedValue: a.expectedValue,
	}

	return newMatcher
}

func (a *structMatcher) matching(value any) expect.StructMatcher {
	matchingValue, ok := getStructValue(value)
	if !ok {
		panic("Like only accepts structs")
	}

	newMatcher := a.clone()
	newMatcher.expectedValue = matchingValue
	return newMatcher
}

func (a *structMatcher) Match(value any) expect.MatchResult {
	structValue, ok := getStructValue(value)
	if !ok {
		return expect.DoesNotMatch(fmt.Sprintf("Expected type struct, but got %T", value), nil)
	}

	problemsByFieldName := map[string]string{}
	// Compare field by field
	for i := range a.expectedValue.NumField() {
		expectedField := a.expectedValue.Type().Field(i)

		// Skip private (unexported) fields
		if !expectedField.IsExported() {
			continue
		}

		expectedFieldValue := a.expectedValue.Field(i)

		gunitTag := expectedField.Tag.Get("gunit")

		// Skip zero values in expected
		if expectedFieldValue.IsZero() && !strings.Contains(gunitTag, "required") {
			continue
		}

		fieldName := expectedField.Name
		actualFieldValue := structValue.FieldByName(fieldName)
		if !actualFieldValue.IsValid() {
			problemsByFieldName[fieldName] = "Field not found in actual value"
			continue
		}
		fieldMatchResult := NewGeneralMatcher(expectedFieldValue.Interface()).
			Match(actualFieldValue.Interface())
		if !fieldMatchResult.Matches() {
			problemsByFieldName[fieldName] = fieldMatchResult.Message()
		}
	}
	if len(problemsByFieldName) > 0 {
		message := "Struct fields did not match:\n"
		for fieldName, problem := range problemsByFieldName {
			message += fmt.Sprintf("  %s: %s\n", fieldName, problem)
		}
		return expect.DoesNotMatch(message, nil)
	}
	return expect.Matches()

}

func getStructValue(value any) (reflect.Value, bool) {
	structValue := reflect.ValueOf(value)
	if structValue.Kind() == reflect.Pointer {
		structValue = structValue.Elem()
	}

	if structValue.Kind() == reflect.Struct {
		return structValue, true
	}

	return reflect.Value{}, false
}
