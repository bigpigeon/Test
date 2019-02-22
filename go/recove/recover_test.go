package recove

import (
	"fmt"
	"runtime/debug"
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
