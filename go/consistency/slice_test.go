package consistency

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

	sli := []bool{true}
	doneChan := make(chan struct{})
	go func() {
		for i := 0; i < 1000*1000; i++ {
			sli = append(sli, true)
		}
		doneChan <- struct{}{}
	}()
	for {
		select {
		case <-doneChan:
			return
		default:
			sliLen := len(sli)
			if sli[sliLen-1] != true {
				fmt.Println("consistency error", sliLen)
			}
		}
	}
}
