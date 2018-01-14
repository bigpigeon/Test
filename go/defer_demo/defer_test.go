package defer_demo

import (
	"testing"
)

func printSlice(l []int, t *testing.T) {
	t.Log(l)
}

func TestDefer(t *testing.T) {
	var slice []int

	defer printSlice(slice, t)
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}
}
