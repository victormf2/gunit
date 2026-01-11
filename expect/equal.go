package expect

import (
	"reflect"
	"testing"
)

func (e *Expector) ToEqual(t *testing.T, expected any) {
	actual := e.value

	matchResult := matchEquals(reflect.ValueOf(actual), reflect.ValueOf(expected))
	if !matchResult.Matches {
		t.Fatal(matchResult.Message)
	}
}
