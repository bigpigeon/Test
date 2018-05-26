package unit_demo

import (
	"testing"
)

func TestMaxInt(t *testing.T) {
	t.Logf("max uintptr %v", ^uintptr(1)/2)
	t.Logf("max uint %v", ^uint(1))
	t.Logf("max int %v", int(^uint(0)>>1))
}
