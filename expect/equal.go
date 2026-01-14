package expect

import (
	"testing"

	"github.com/victormf2/gunit/expect/matchers"
)

func (e *Expector) ToEqual(t *testing.T, expected any) {
	actual := e.value

	matchResult := (&matchers.EqualMatcher{Expected: expected}).Match(actual)
	if !matchResult.Matches {
		t.Fatal(matchResult.Message)
	}
}
