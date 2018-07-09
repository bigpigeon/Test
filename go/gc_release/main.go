/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"runtime"
	"time"
)

type B struct {
	bb []int
}

func NewB() *B {
	return new(B)
}

func main() {
	var bs = make([]*B, 10)
	for i := 0; i < 10; i++ {
		bs[i] = NewB()
		bs[i].bb = make([]int, 1000000)
	}

	time.Sleep(time.Second)
	println("begin gc")
	//for i:=0; i<10; i++ {
	//  bs[i] = nil
	//}
	bs = nil
	runtime.GC()
	time.Sleep(time.Second * 2)
	runtime.GC()
	time.Sleep(time.Second * 2)
}
