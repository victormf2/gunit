package expect

import (
	"testing"
)

func (e *Expector) ToEqual(t *testing.T, expected any) {
	actual := e.value

	if actual != expected {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}
