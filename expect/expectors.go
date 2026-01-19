package expect

import "github.com/victormf2/gunit/expect/expectors"

type Expector = expectors.Expector

func It(value any) Expector {
	return expectors.NewExpector(value)
}
