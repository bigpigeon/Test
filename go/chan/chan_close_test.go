package _chan

import (
	"fmt"
	"testing"
	"time"
)

func TestChanClose(t *testing.T) {
	a := make(chan int, 10)
	for i := 0; i < 10; i++ {
		a <- i
	}
	close(a)

	for i := 0; i < 12; i++ {
		t.Log(<-a)
	}

	//for {
	//	if i, more := <-a; more {
	//		t.Log(i)
	//	} else {
	//		t.Log(<-a)
	//		break
	//	}
	//
	//}
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

func TestChanFullSend(t *testing.T) {
	c := make(chan int)
	select {
	case c <- 2:
		t.Log("send successful")
	default:
		t.Log("queue full")
	}

	select {
	case c <- 2:
		t.Log("send successful")
	case <-time.After(time.Millisecond):
		t.Log("timeout")
	}
}

func TestCloseSend(t *testing.T) {
	c := make(chan int)
	close(c)
	go func() {
		<-c
		receive := 0
		defaultNum := 0
		for i := 0; i < 1000; i++ {
			select {
			case <-c:
				receive++
			default:
				defaultNum++
			case c <- 2:
			}
		}
		t.Log("receive num", receive)
		t.Log("default num", defaultNum)
	}()
	time.Sleep(1 * time.Second)
}
