package expectations

import (
	"context"

	"github.com/victormf2/gunit/internal/expect/matchers"
)

func Context(ctx context.Context) context.Context {
	return matchers.ContextWithTracker(ctx)
}
