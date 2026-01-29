package expectors_test

import (
	"testing"

	"github.com/victormf2/gunit/internal"
	"github.com/victormf2/gunit/internal/expectors"
)

func TestToPanic(t *testing.T) {
	t.Run("only allow function values", func(t *testing.T) {
		mockT := &internal.MockT{}
		expectors.NewExpector(42).ToPanic(mockT)
		if !mockT.Failed {
			t.Fatal("Expected test to fail when non-function is provided")
		}
	})

	t.Run("fail on non-panic", func(t *testing.T) {
		mockT := &internal.MockT{}
		expectors.NewExpector(func() {}).ToPanic(mockT)
		if !mockT.Failed {
			t.Fatal("Expected test to fail when function does not panic")
		}
	})
}
