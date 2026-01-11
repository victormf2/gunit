package expect

import (
	"reflect"

	"github.com/victormf2/gunit/gunit"
)

func (e *Expector) ToPanic(t gunit.TestingT) {
	functionValue := reflect.ValueOf(e.value)
	if functionValue.Kind() != reflect.Func {
		t.Fatalf("ToPanic expects a function, but got %s", functionValue.Kind())
	}

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expected function to panic, but it did not")
		}
	}()

	// Call the function
	functionValue.Call(nil)
}
