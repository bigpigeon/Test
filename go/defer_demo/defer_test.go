package main

import (
	"fmt"
	"testing"
)

func printSlice(l []int, t *testing.T) {
	t.Log(l)
}

func TestDefer(t *testing.T) {
	var slice []int

	defer printSlice(slice, t)
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}
}

func TestScopeDefer(t *testing.T) {
	{
		defer func() {
			fmt.Println("out scope 1")
		}()
	}
	{
		defer func() {
			fmt.Println("out scope 2")
		}()
	}
	defer func() {
		fmt.Println("out main scope")
	}()
}

func TestDeferRefence(t *testing.T) {
	var ret int
	defer func() {
		fmt.Println(ret)
	}()

	ret = 2

}

type XX struct {
	A string
}

func (x *XX) Foo() {
	fmt.Println(x.A)
}

func TestDeferMethod(t *testing.T) {
	var x *XX
	for _, v := range []string{
		"AA",
		"BB",
	} {
		x = &XX{A: v}
		defer x.Foo()
	}
}

func TestDeferPanic(t *testing.T) {
	defer func() {
		fmt.Println("aa")
	}()
	panic("bb")
}
