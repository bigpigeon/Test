package lib

import "fmt"

func Hello() {
	fmt.Println("Hello")
}

// gen
type Response struct {
	Name  string
	Value string
}
