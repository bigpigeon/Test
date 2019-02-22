/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func fn() {
	const MAX = 10 * 1000 * 1000
	var m = make(map[int]string, MAX)
	for i := 0; i < MAX; i++ {
		m[i] = fmt.Sprint(i)
	}

	time.Sleep(time.Second)
}

func main() {

	fn()
	println("begin gc")
	//for i:=0; i<10; i++ {
	//  bs[i] = nil
	//}
	//time.Sleep(time.Second * 5)
	//runtime.GC()
	//time.Sleep(time.Second * 5)
	runtime.GC()
	debug.FreeOSMemory()
	time.Sleep(time.Second)
	runtime.GC()
	time.Sleep(time.Second)
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	fmt.Println("xxx", m.HeapObjects, m.HeapAlloc, m.TotalAlloc)
}
