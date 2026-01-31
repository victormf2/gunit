package matchers

import (
	"context"

	"github.com/victormf2/gunit/internal/expect"
)

type ContextMatcher interface {
	expect.Matcher
}

func NewContextMatcher(ctx context.Context) ContextMatcher {
	return &contextMatcher{
		expectedCtx: ctx,
	}
}

type contextMatcher struct {
	expectedCtx context.Context
}

func (c *contextMatcher) Match(value any) expect.MatchResult {
	actualCtx, ok := value.(context.Context)
	if !ok {
		return expect.DoesNotMatch("not a context", nil)
	}

	if actualCtx == c.expectedCtx {
		return expect.Matches()
	}

	actualCtxTracker := actualCtx.Value(contextTracker(0))
	if actualCtxTracker != nil {
		expectedCtxTracker := c.expectedCtx.Value(contextTracker(0))
		if expectedCtxTracker != nil {
			if actualCtxTracker == expectedCtxTracker {
				return expect.Matches()
			}
		}
	}

	return expect.DoesNotMatch("different contexts", nil)
}

type contextTracker int64

func ContextWithTracker(ctx context.Context) context.Context {
	// creating unique value for matching ctx identity
	ctxTracker := &struct{}{}
	ctx = context.WithValue(ctx, contextTracker(0), ctxTracker)
	return ctx
}
