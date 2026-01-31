package matchers_test

import (
	"testing"
	"time"

	"github.com/victormf2/gunit/internal"
	"github.com/victormf2/gunit/internal/expect/matchers"
)

func TestGeneralMatcher(t *testing.T) {
	t.Run("scalars", func(t *testing.T) {
		testCases := []struct {
			desc          string
			actualValue   any
			expectedValue any
			matches       bool
		}{
			{
				desc:          "matches nil",
				actualValue:   nil,
				expectedValue: nil,
				matches:       true,
			},
			{
				desc:          "does not match non-nil",
				actualValue:   "non nil",
				expectedValue: nil,
				matches:       false,
			},
			{
				desc:          "matches string",
				actualValue:   "hello",
				expectedValue: "hello",
				matches:       true,
			},
			{
				desc:          "does not match different strings",
				actualValue:   "hello",
				expectedValue: "hey",
				matches:       false,
			},
			{
				desc:          "does not match non-string",
				actualValue:   "hello",
				expectedValue: 42,
				matches:       false,
			},
			{
				desc:          "matches int",
				actualValue:   42,
				expectedValue: 42,
				matches:       true,
			},
			{
				desc:          "does not match different ints",
				actualValue:   35,
				expectedValue: 42,
				matches:       false,
			},
			{
				desc:          "does not match non-int",
				actualValue:   42.0,
				expectedValue: 42,
				matches:       false,
			},
			{
				desc:          "matches uint",
				actualValue:   uint(42),
				expectedValue: uint(42),
				matches:       true,
			},
			{
				desc:          "does not match different uints",
				actualValue:   uint(35),
				expectedValue: uint(42),
				matches:       false,
			},
			{
				desc:          "does not match non-uint",
				actualValue:   42,
				expectedValue: uint(42),
				matches:       false,
			},
			{
				desc:          "matches float",
				actualValue:   42.0,
				expectedValue: 42.0,
				matches:       true,
			},
			{
				desc:          "does not match different floats",
				actualValue:   42.1,
				expectedValue: 42.0,
				matches:       false,
			},
			{
				desc:          "does not match non-float",
				actualValue:   42,
				expectedValue: 42.0,
				matches:       false,
			},
			{
				desc:          "matches true",
				actualValue:   true,
				expectedValue: true,
				matches:       true,
			},
			{
				desc:          "matches false",
				actualValue:   false,
				expectedValue: false,
				matches:       true,
			},
			{
				desc:          "does not match different bools",
				actualValue:   true,
				expectedValue: false,
				matches:       false,
			},
			{
				desc:          "does not match non-bool",
				actualValue:   "non-bool",
				expectedValue: false,
				matches:       false,
			},
		}
		for _, tC := range testCases {
			t.Run(tC.desc, func(t *testing.T) {
				matcher := matchers.NewGeneralMatcher(tC.expectedValue)
				matchResult := matcher.Match(tC.actualValue)
				if matchResult.Matches() != tC.matches {
					t.Errorf("Expected matches to be %v but got %v", tC.matches, matchResult.Matches())
				}
			})
		}
	})

	t.Run("slices", func(t *testing.T) {
		testCases := []struct {
			desc          string
			actualValue   any
			expectedValue any
			matches       bool
		}{
			{
				desc:          "matches zero length slices",
				actualValue:   []any{},
				expectedValue: []any{},
				matches:       true,
			},
			{
				desc:          "doesn't match slices with less length",
				actualValue:   []any{1},
				expectedValue: []any{1, 2},
				matches:       false,
			},
			{
				desc:          "doesn't match slices with more length",
				actualValue:   []any{1, 2, 3},
				expectedValue: []any{1, 2},
				matches:       false,
			},
			{
				desc:          "matches slices with scalars",
				actualValue:   []any{1, 2},
				expectedValue: []any{1, 2},
				matches:       true,
			},
			{
				desc:          "matches arrays with slices",
				actualValue:   [2]any{1, 2},
				expectedValue: []any{1, 2},
				matches:       true,
			},
			{
				desc:          "matches slices with arrays",
				actualValue:   []any{1, 2},
				expectedValue: [2]any{1, 2},
				matches:       true,
			},
			{
				desc:          "matches slices with different types",
				actualValue:   []int{1, 2},
				expectedValue: []any{1, 2},
				matches:       true,
			},
			{
				desc: "matches slices with structs",
				actualValue: []any{
					internal.BigStruct{
						String: "hello",
						Number: 12,
						Bool:   true,
					},
				},
				expectedValue: []any{
					internal.BigStruct{
						String: "hello",
						Bool:   true,
					},
				},
				matches: true,
			},
			{
				desc:        "matches slices with custom element matchers",
				actualValue: []any{1, 2},
				expectedValue: []any{
					1,
					matchers.NewIntMatcher().GreaterThan(1),
				},
				matches: true,
			},
		}
		for _, tC := range testCases {
			t.Run(tC.desc, func(t *testing.T) {
				matcher := matchers.NewGeneralMatcher(tC.expectedValue)
				matchResult := matcher.Match(tC.actualValue)
				if matchResult.Matches() != tC.matches {
					t.Errorf("Expected matches to be %v but got %v", tC.matches, matchResult.Matches())
				}
			})
		}
	})

	t.Run("maps", func(t *testing.T) {
		testCases := []struct {
			desc          string
			actualValue   any
			expectedValue any
			matches       bool
		}{
			{
				desc:          "matches empty maps",
				actualValue:   map[any]any{},
				expectedValue: map[any]any{},
				matches:       true,
			},
			{
				desc:          "doesn't match maps with less length",
				actualValue:   map[any]any{"a": 1},
				expectedValue: map[any]any{"a": 1, "b": 2},
				matches:       false,
			},
			{
				desc:          "matches partial maps",
				actualValue:   map[any]any{"a": 1, "b": 2, "c": 3},
				expectedValue: map[any]any{"a": 1, "b": 2},
				matches:       true,
			},
			{
				desc:          "matches equal maps",
				actualValue:   map[any]any{"a": 1, "b": 2},
				expectedValue: map[any]any{"a": 1, "b": 2},
				matches:       true,
			},
			{
				desc:          "matches maps with nil values",
				actualValue:   map[any]any{"a": nil},
				expectedValue: map[any]any{"a": nil},
				matches:       true,
			},
			{
				desc:          "matches maps with custom value matchers",
				actualValue:   map[any]any{"a": 1},
				expectedValue: map[any]any{"a": matchers.NewIntMatcher().GreaterThan(0)},
				matches:       true,
			},
			{
				desc:          "matches maps with custom key matchers",
				actualValue:   map[any]any{"a": 1},
				expectedValue: map[any]any{matchers.NewStringMatcher().ContainingAny("a"): 1},
				matches:       true,
			},
			{
				desc:          "matches maps of different types",
				actualValue:   map[string]int{"a": 1},
				expectedValue: map[any]any{"a": 1},
				matches:       true,
			},
		}
		for _, tC := range testCases {
			t.Run(tC.desc, func(t *testing.T) {
				matcher := matchers.NewGeneralMatcher(tC.expectedValue)
				matchResult := matcher.Match(tC.actualValue)
				if matchResult.Matches() != tC.matches {
					t.Errorf("Expected matches to be %v but got %v", tC.matches, matchResult.Matches())
				}
			})
		}
	})

	t.Run("structs", func(t *testing.T) {
		testCases := []struct {
			desc          string
			actualValue   any
			expectedValue any
			matches       bool
		}{
			{
				desc: "matches equal structs",
				actualValue: internal.BigStruct{
					String:      "hello",
					Number:      42,
					Bool:        true,
					Date:        time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
					SimpleSlice: []string{"a", "b"},
					NestedSlice: []internal.NestedStruct{
						{ID: 1, Value: "a"},
						{ID: 2, Value: "b"},
					},
					SimpleMap: map[string]int{
						"a": 1,
						"b": 2,
					},
					NestedMap: map[string]internal.NestedStruct{
						"a": {ID: 1, Value: "aa"},
						"b": {ID: 2, Value: "bb"},
					},
					Struct: internal.NestedStruct{
						ID:    42,
						Value: "hello",
					},
				},
				expectedValue: internal.BigStruct{
					String:      "hello",
					Number:      42,
					Bool:        true,
					Date:        time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
					SimpleSlice: []string{"a", "b"},
					NestedSlice: []internal.NestedStruct{
						{ID: 1, Value: "a"},
						{ID: 2, Value: "b"},
					},
					SimpleMap: map[string]int{
						"a": 1,
						"b": 2,
					},
					NestedMap: map[string]internal.NestedStruct{
						"a": {ID: 1, Value: "aa"},
						"b": {ID: 2, Value: "bb"},
					},
					Struct: internal.NestedStruct{
						ID:    42,
						Value: "hello",
					},
				},
				matches: true,
			},
			{
				desc: "matches partial structs",
				actualValue: internal.BigStruct{
					String:      "hello",
					Number:      42,
					Bool:        true,
					Date:        time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
					SimpleSlice: []string{"a", "b"},
					NestedSlice: []internal.NestedStruct{
						{ID: 1, Value: "a"},
						{ID: 2, Value: "b"},
					},
					SimpleMap: map[string]int{
						"a": 1,
						"b": 2,
					},
					NestedMap: map[string]internal.NestedStruct{
						"a": {ID: 1, Value: "aa"},
						"b": {ID: 2, Value: "bb"},
					},
					Struct: internal.NestedStruct{
						ID:    42,
						Value: "hello",
					},
				},
				expectedValue: internal.BigStruct{
					String:      "hello",
					SimpleSlice: []string{"a", "b"},
					NestedSlice: []internal.NestedStruct{
						{ID: 1},
						{Value: "b"},
					},
					SimpleMap: map[string]int{
						"b": 2,
					},
					NestedMap: map[string]internal.NestedStruct{
						"a": {Value: "aa"},
					},
					Struct: internal.NestedStruct{
						ID:    42,
						Value: "hello",
					},
				},
				matches: true,
			},
		}
		for _, tC := range testCases {
			t.Run(tC.desc, func(t *testing.T) {
				matcher := matchers.NewGeneralMatcher(tC.expectedValue)
				matchResult := matcher.Match(tC.actualValue)
				if matchResult.Matches() != tC.matches {
					t.Errorf("Expected matches to be %v but got %v", tC.matches, matchResult.Matches())
				}
			})
		}
	})
}
