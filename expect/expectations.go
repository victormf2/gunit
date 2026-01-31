package expect

import (
	"context"
	"testing"

	"github.com/victormf2/gunit/internal/expect/expectations"
)

func It(value any) Expector {
	return expectations.NewExpector(value)
}

func Context(t *testing.T) context.Context {
	ctx := expectations.Context(t.Context())
	return ctx
}
