package _chan

import "testing"

func TestChanClose(t *testing.T) {
	a := make(chan int, 15)
	for i := 0; i < 10; i++ {
		a <- i
	}
	close(a)
	//for i := 0; i < 10; i++ {
	//	t.Log(<-a)
	//}

	for {
		if i, more := <-a; more {
			t.Log(i)
		} else {
			break
		}

	}
}
