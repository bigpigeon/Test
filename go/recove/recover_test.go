package recove

import (
	"fmt"
	"runtime/debug"
	"sync"
	"testing"
	"time"
)

func doSth() {
	fmt.Println("23")
	panic("test error")
}

func TestRecover(t *testing.T) {
	go func() {
		for {
			fmt.Println("123")
			time.Sleep(1)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("runtime error: %s\n", r)
			debug.PrintStack()
		}
	}()
	doSth()
}

func TestRecoverWithLock(t *testing.T) {
	l := sync.Mutex{}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("runtime error: %s\n", r)
			t.Logf("l status %#v\n", l)
		}
	}()
	t.Logf("l status %#v\n", l)
	l.Lock()
	t.Logf("l status %#v\n", l)
	defer l.Unlock()
	defer func() {
		fmt.Println("before recover")
	}()
	panic("panic!!!")
}
