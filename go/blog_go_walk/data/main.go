package lib

import (
	"C"
	"bytes"
	"fmt"
	"io"
)

func Hello(name string) (int, error) {
	return fmt.Println("Hello", name)
}

type Stringer interface {
	String() string
}

// named 类型包含它的method信息
type Response struct {
	Name  string
	Value string
	Buff  bytes.Buffer
}

func (r Response) String() string {
	return r.Name + " " + r.Value
}

func (r *Response) Write(w io.Writer) {
	r.Buff.WriteTo(w)
}

var v1 int
var v2 string
var v3 chan int
var v4 interface {
	Stringer
	Write(io.Writer)
	setAttr(s string)
}
var v5 []int
var v6 [2]int
var v7 map[int]int
var v8 C.int
var v9 *int

var idName1 struct {
	ID   int
	Name string
}

var idName2 struct {
	ID   int
	Name string
}

var idName3 struct {
	ID   int
	Name []byte
}

type IDName struct {
	ID   int
	Name string
}

var idName4 IDName

var intVal int

type CustomInt int

var intVal2 CustomInt
