package internal

import "time"

type BigStruct struct {
	String      string `gunit:"required"`
	Number      int
	Bool        bool
	Date        time.Time
	SimpleSlice []string
	NestedSlice []NestedStruct
	SimpleMap   map[string]int
	NestedMap   map[string]NestedStruct
	Struct      NestedStruct
}
type NestedStruct struct {
	ID    int
	Value string
}
