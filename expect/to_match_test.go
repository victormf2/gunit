package expect_test

import (
	"testing"
	"time"

	"github.com/victormf2/gunit/expect"
	"github.com/victormf2/gunit/internal"
)

func TestMatch(t *testing.T) {
	t.Run("matches nil", func(t *testing.T) {
		expect.It(nil).ToMatch(t, nil)
	})
	t.Run("matches nil different types", func(t *testing.T) {
		actual := (*string)(nil)
		expected := (*int)(nil)
		expect.It(actual).ToMatch(t, expected)
	})
	t.Run("matches scalars", func(t *testing.T) {
		expect.It(42).ToMatch(t, 42)
		expect.It(0).ToMatch(t, 0)
		expect.It(3.14).ToMatch(t, 3.14)
		expect.It(0.0).ToMatch(t, 0.0)
		expect.It("hello").ToMatch(t, "hello")
		expect.It("").ToMatch(t, "")
		expect.It(false).ToMatch(t, false)
		expect.It(true).ToMatch(t, true)
		expect.It(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)).ToMatch(t, time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	})
	t.Run("matches slices", func(t *testing.T) {
		actual := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		expect.It(actual).ToMatch(t, expected)
	})
	t.Run("matches maps", func(t *testing.T) {
		actual := map[string]int{
			"one": 1,
			"two": 2,
		}
		expected := map[string]int{
			"one": 1,
			"two": 2,
		}
		expect.It(actual).ToMatch(t, expected)
	})
	t.Run("matches identical structs", func(t *testing.T) {
		actual := internal.BigStruct{
			String:      "test",
			Number:      42,
			Bool:        true,
			Date:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			SimpleSlice: []string{"a", "b", "c"},
			NestedSlice: []internal.NestedStruct{
				{ID: 1, Value: "one"},
				{ID: 2, Value: "two"},
			},
			SimpleMap: map[string]int{
				"one": 1,
				"two": 2,
			},
			NestedMap: map[string]internal.NestedStruct{
				"first":  {ID: 1, Value: "one"},
				"second": {ID: 2, Value: "two"},
			},
			Struct: internal.NestedStruct{ID: 99, Value: "ninety-nine"},
		}
		expected := &internal.BigStruct{
			String:      "test",
			Number:      42,
			Bool:        true,
			Date:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			SimpleSlice: []string{"a", "b", "c"},
			NestedSlice: []internal.NestedStruct{
				{ID: 1, Value: "one"},
				{ID: 2, Value: "two"},
			},
			SimpleMap: map[string]int{
				"one": 1,
				"two": 2,
			},
			NestedMap: map[string]internal.NestedStruct{
				"first":  {ID: 1, Value: "one"},
				"second": {ID: 2, Value: "two"},
			},
			Struct: internal.NestedStruct{ID: 99, Value: "ninety-nine"},
		}

		expect.It(actual).ToMatch(t, expected)
	})
	t.Run("partial match of structs", func(t *testing.T) {
		actual := internal.BigStruct{
			String:      "test",
			Number:      42,
			Bool:        true,
			Date:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			SimpleSlice: []string{"a", "b", "c"},
			NestedSlice: []internal.NestedStruct{
				{ID: 1, Value: "one"},
				{ID: 2, Value: "two"},
			},
			SimpleMap: map[string]int{
				"one": 1,
				"two": 2,
			},
			NestedMap: map[string]internal.NestedStruct{
				"first":  {ID: 1, Value: "one"},
				"second": {ID: 2, Value: "two"},
			},
			Struct: internal.NestedStruct{ID: 99, Value: "ninety-nine"},
		}
		expected := &internal.BigStruct{
			String: "test",
			Date:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			NestedSlice: []internal.NestedStruct{
				{ID: 1},
				{Value: "two"},
			},
			SimpleMap: map[string]int{
				"one": 1,
				"two": 2,
			},
			NestedMap: map[string]internal.NestedStruct{
				"first":  {ID: 1},
				"second": {Value: "two"},
			},
			Struct: internal.NestedStruct{ID: 99},
		}

		expect.It(actual).ToMatch(t, expected)
	})

	t.Run("required tag", func(t *testing.T) {
		actual := internal.BigStruct{}
		expected := &internal.BigStruct{
			String: "must be set",
		}
		expect.It(actual).ToMatch(t, expected)
	})
}
