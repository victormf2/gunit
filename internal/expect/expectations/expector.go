package expectations

import (
	"github.com/victormf2/gunit/internal/expect"
)

func NewExpector(value any) expect.Expector {
	return &expector{
		value: value,
	}
}

type expector struct {
	value any
}
