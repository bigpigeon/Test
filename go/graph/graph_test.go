package graph

import (
	"testing"
)

type SubSubData struct {
	Data string
}

type SubData struct {
	Data string
	Sub  []SubSubData
}

type MainData struct {
	ID   int
	Data string
	Sub  []SubData
}

func TestGraph(t *testing.T) {
	data := []MainData{
		{1, "main data 1", []SubData{
			{"data 1 sub 1", []SubSubData{
				{"data 1 sub 1 sub 1"},
				{"data 1 sub 1 sub 2"},
			}},
			{"data 1 sub 2", []SubSubData{
				{"data 1 sub 2 sub 1"},
				{"data 1 sub 2 sub 2"},
			}},
		}},
		{2, "main data 2", []SubData{
			{"data 2 sub 1", []SubSubData{
				{"data 2 sub 1 sub 1"},
				{"data 2 sub 1 sub 2"},
			}},
			{"data 2 sub 2", []SubSubData{
				{"data 2 sub 2 sub 1"},
				{"data 2 sub 2 sub 2"},
			}},
		}},
	}
}
