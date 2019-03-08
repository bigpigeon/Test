package _chan

import (
	"fmt"
	"testing"
)

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
			t.Log(<-a)
			break
		}

	}
}

func TestChanSelectClose(t *testing.T) {
	c := make(chan bool)
	close(c)
	select {
	case <-c:
		fmt.Println("close event")
	}
	fmt.Printf("chan %#v\n", c)
}
