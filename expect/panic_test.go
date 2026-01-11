package expect_test

import (
	"testing"

	"github.com/victormf2/gunit/expect"
	"github.com/victormf2/gunit/internal"
)

func TestToPanic(t *testing.T) {
	t.Run("only allow function values", func(t *testing.T) {
		mockT := &internal.MockT{}
		expect.It(42).ToPanic(mockT)
		if !mockT.Failed {
			t.Fatal("Expected test to fail when non-function is provided")
		}
	})

	t.Run("fail on non-panic", func(t *testing.T) {
		mockT := &internal.MockT{}
		expect.It(func() {}).ToPanic(mockT)
		if !mockT.Failed {
			t.Fatal("Expected test to fail when function does not panic")
		}
	})
}
