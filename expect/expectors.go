package expect

import (
	"context"
	"testing"

	"github.com/victormf2/gunit/internal/expectors"
)

type Expector = expectors.Expector

func It(value any) Expector {
	return expectors.NewExpector(value)
}

func Context(t *testing.T) context.Context {
	ctx := expectors.Context(t.Context())
	return ctx
}
