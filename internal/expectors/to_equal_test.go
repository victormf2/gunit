package expectors_test

import (
	"testing"

	"github.com/victormf2/gunit/internal"
	"github.com/victormf2/gunit/internal/expectors"
)

func TestToEqual(t *testing.T) {
	testCases := []struct {
		desc        string
		actualValue any
		failed      bool
	}{
		{
			desc: "doesn't fail when values are equal",
			actualValue: internal.BigStruct{
				String: "a",
				Number: 2,
			},
			failed: false,
		},
		{
			desc: "fails when values are not equal",
			actualValue: internal.BigStruct{
				String: "a",
				Number: 2,
				Bool:   true,
			},
			failed: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			expectedValue := internal.BigStruct{
				String: "a",
				Number: 2,
			}

			mockT := &internal.MockT{}
			expectors.NewExpector(tC.actualValue).ToEqual(mockT, expectedValue)
			if tC.failed != mockT.Failed {
				t.Errorf("Expect it to fail the test: %v, actually failed: %v", tC.failed, mockT.Failed)
			}
		})
	}

}
