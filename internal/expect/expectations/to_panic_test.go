package expectations_test

import (
	"testing"

	"github.com/victormf2/gunit/internal"
	"github.com/victormf2/gunit/internal/expect/expectations"
)

func TestToPanic(t *testing.T) {
	t.Run("only allow function values", func(t *testing.T) {
		mockT := &internal.MockT{}
		expectations.NewExpector(42).ToPanic(mockT)
		if !mockT.Failed {
			t.Fatal("Expected test to fail when non-function is provided")
		}
	})

	t.Run("fail on non-panic", func(t *testing.T) {
		mockT := &internal.MockT{}
		expectations.NewExpector(func() {}).ToPanic(mockT)
		if !mockT.Failed {
			t.Fatal("Expected test to fail when function does not panic")
		}
	})
}
