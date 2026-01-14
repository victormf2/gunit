package matchers

import (
	"fmt"
	"reflect"
	"strings"
)

type AnyStructMatcher struct {
	matchingValue reflect.Value
}

func (a *AnyStructMatcher) clone() *AnyStructMatcher {
	newMatcher := &AnyStructMatcher{
		matchingValue: a.matchingValue,
	}

	return newMatcher
}

func (a *AnyStructMatcher) matching(value any) *AnyStructMatcher {
	matchingValue, ok := getStructValue(value)
	if !ok {
		panic("Like only accepts structs")
	}

	newMatcher := a.clone()
	newMatcher.matchingValue = matchingValue
	return newMatcher
}

func (a *AnyStructMatcher) Match(value any) MatchResult {
	structValue, ok := getStructValue(value)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: fmt.Sprintf("Expected type struct, but got %T", value),
		}
	}

	problemsByFieldName := map[string]string{}
	// Compare field by field
	for i := range a.matchingValue.NumField() {
		expectedField := a.matchingValue.Type().Field(i)

		// Skip private (unexported) fields
		if !expectedField.IsExported() {
			continue
		}

		expectedFieldValue := a.matchingValue.Field(i)

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
		fieldMatchResult := (&GeneralMatcher{}).
			Matching(expectedFieldValue.Interface()).
			Match(actualFieldValue.Interface())
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
