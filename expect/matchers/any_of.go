package matchers

import "fmt"

type AnyOfMatcher[T any] struct{}

func (a *AnyOfMatcher[T]) Match(value any) MatchResult {
	_, ok := value.(T)
	if ok {
		return MatchResult{Matches: true}
	}
	return MatchResult{
		Matches: false,
		Message: fmt.Sprintf("Expected type %T, but got %T", *new(T), value),
	}
}
