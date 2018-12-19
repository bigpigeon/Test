/*
 * Copyright 2018 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import "C"
import (
	"fmt"
	"unsafe"
)

//export StringCmp
func StringCmp(s1, s2, s3 []byte) *C.char {
	fmt.Printf("p %p, %p, %p\n", s1, s2, s3)
	fmt.Printf("v %v, %v, %v\n", s1, s2, s3)
	fmt.Printf("str %v, %v, %v\n", string(s1), string(s2), string(s3))
	s1C := C.CBytes(s1)
	s2C := C.CBytes(s2)
	s3C := C.CBytes(s3)

	fmt.Printf("cv %v, %v, %v\n", s1C, s2C, s3C)
	return nil
}

//export IntSlicePrint
func IntSlicePrint(v []int) *C.longlong {
	fmt.Printf("v %v\n", v)
	fmt.Printf("int size %d\n", C.size_t(unsafe.Sizeof(int(0))))
	cArray := C.malloc(C.size_t(len(v)) * C.size_t(unsafe.Sizeof(int(0))))
	a := (*[1<<30 - 1]C.longlong)(cArray)
	for i, d := range v {
		a[i] = C.longlong(d)
	}
	return (*C.longlong)(cArray)
}

//export StrConv
func StrConv(s1 string) *C.char {
	fmt.Printf("str %s\n", s1)
	s1C := C.CString(s1)
	fmt.Printf("point %v\n", s1C)
	return s1C
}

//export Sum
func Sum(a int, b int) int {
	return a + b
}

func init() {
	fmt.Printf("go init!!\n")
}

func main() {}
