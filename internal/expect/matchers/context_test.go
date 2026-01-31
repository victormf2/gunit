package matchers_test

import (
	"context"
	"testing"

	"github.com/victormf2/gunit/internal/expect/matchers"
)

func TestContextMatcher(t *testing.T) {
	t.Run("should match equal contexts", func(t *testing.T) {
		expectedCtx := context.TODO()
		actualCtx := context.TODO()

		matcher := matchers.NewContextMatcher(expectedCtx)
		matchResult := matcher.Match(actualCtx)

		if !matchResult.Matches() {
			t.Errorf("expected matches to be true")
		}
	})

	t.Run("should match contexts with same tracker", func(t *testing.T) {
		expectedCtx := matchers.ContextWithTracker(context.TODO())
		actualCtx := context.WithValue(expectedCtx, "a", 1)

		matcher := matchers.NewContextMatcher(expectedCtx)
		matchResult := matcher.Match(actualCtx)

		if !matchResult.Matches() {
			t.Errorf("expected matches to be true")
		}
	})

	t.Run("should not match different contexts", func(t *testing.T) {
		expectedCtx := matchers.ContextWithTracker(context.TODO())
		actualCtx := context.TODO()

		matcher := matchers.NewContextMatcher(expectedCtx)
		matchResult := matcher.Match(actualCtx)

		if matchResult.Matches() {
			t.Errorf("expected matches to be false")
		}
	})
}
