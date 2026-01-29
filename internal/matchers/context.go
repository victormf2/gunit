package matchers

import (
	"context"
)

type ContextMatcher interface {
	Matcher
}

func NewContextMatcher(ctx context.Context) ContextMatcher {
	return &contextMatcher{
		expectedCtx: ctx,
	}
}

type contextMatcher struct {
	expectedCtx context.Context
}

func (c *contextMatcher) Match(value any) MatchResult {
	actualCtx, ok := value.(context.Context)
	if !ok {
		return MatchResult{
			Matches: false,
			Message: "not a context",
		}
	}

	if actualCtx == c.expectedCtx {
		return MatchResult{Matches: true}
	}

	actualCtxTracker := actualCtx.Value(contextTracker(0))
	if actualCtxTracker != nil {
		expectedCtxTracker := c.expectedCtx.Value(contextTracker(0))
		if expectedCtxTracker != nil {
			if actualCtxTracker == expectedCtxTracker {
				return MatchResult{Matches: true}
			}
		}
	}

	return MatchResult{
		Matches: false,
		Message: "different contexts",
	}
}

type contextTracker int64

func ContextWithTracker(ctx context.Context) context.Context {
	ctxTracker := &struct{}{}
	ctx = context.WithValue(ctx, contextTracker(0), ctxTracker)
	return ctx
}
