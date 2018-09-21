package wait_group

import (
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Done()
	wg.Done()
	t.Logf("wg %v", wg)
}
