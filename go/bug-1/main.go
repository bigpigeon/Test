package main

import (
	"fmt"
)

type Printer interface {
	Print()
}

type myLogger struct {
	name string
}

func (log *myLogger) Print() {
	fmt.Printf("test")
}

func main() {
	loggerMap := map[string]*myLogger{}
	var printer Printer = loggerMap["not exist"]
	if printer == nil {
		fmt.Print("yes printer is nil")
	} else {
		fmt.Print("now have a bug")
	}
}
