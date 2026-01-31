package mock

import "github.com/victormf2/gunit/internal/mock"

func NewMockFunction(name string, fn any) MockFunction {
	return mock.NewMockFunction(name, fn)
}
