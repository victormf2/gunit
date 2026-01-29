package expectors

import (
	"context"

	"github.com/victormf2/gunit/internal/matchers"
)

func Context(ctx context.Context) context.Context {
	return matchers.ContextWithTracker(ctx)
}
