package matchers

import (
	"fmt"

	"github.com/victormf2/gunit/internal/expect"
)

func NewAnyOfMatcher[T any]() expect.AnyOfMatcher[T] {
	return &anyOfMatcher[T]{}
}

type anyOfMatcher[T any] struct{}

func (a *anyOfMatcher[T]) Match(value any) expect.MatchResult {
	_, ok := value.(T)
	if ok {
		return expect.Matches()
	}
	return expect.DoesNotMatch(fmt.Sprintf("Expected type %T, but got %T", *new(T), value), nil)
}
