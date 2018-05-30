package main

// implicitly import
import "fmt"

import (
	. "unsafe"
)

type Product struct {
	Name string
}

func (p Product) String() string {
	return "Product:" + p.Name
}

func ImplicitlyNode() {
	var d interface{} = 5
	switch x := d.(type) {
	case int:
		fmt.Println(x)
	default:
		fmt.Println(x)
	}
	var e func(int)
	fmt.Println(e)
}

func SelectionNode() {
	p := Product{Name: "t011"}
	fmt.Println(p.Name)
	fmt.Println(p.String())
	fmt.Println(Product.String(p))
	fmt.Println(Offsetof(p.Name))
}

const MaxRoutines = 100

var CurrentRoutines = 1

func main() { //test1
	fmt.Println("Hello, World!") //test2
	a := []int{1, 2, 3}
	fmt.Println(a)
	b := map[int]string{
		1: "a",
		2: "b",
	}
	fmt.Println(b)

	d := make(chan int, 5)
	fmt.Println(d)

	fmt.Println(MaxRoutines)
}
