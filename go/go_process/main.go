package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	var i int
	go func() {
		fmt.Println("1")
	}()
	for {
		i++
	}
}
