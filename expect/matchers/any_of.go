package matchers

import "fmt"

type AnyOfMatcher[T any] interface {
	Matcher
}

func NewAnyOfMatcher[T any]() AnyOfMatcher[T] {
	return &anyOfMatcher[T]{}
}

type anyOfMatcher[T any] struct{}

func (a *anyOfMatcher[T]) Match(value any) MatchResult {
	_, ok := value.(T)
	if ok {
		return MatchResult{Matches: true}
	}
	return MatchResult{
		Matches: false,
		Message: fmt.Sprintf("Expected type %T, but got %T", *new(T), value),
	}
}
