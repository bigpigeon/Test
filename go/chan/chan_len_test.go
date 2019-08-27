package _chan

import "testing"

func TestChan(t *testing.T) {
	c := make(chan int, 100)
	for i := 0; i < 34; i++ {
		c <- 0
	}
	t.Log(len(c))
	t.Log(cap(c))
}
