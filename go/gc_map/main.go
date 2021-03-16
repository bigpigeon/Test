/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"runtime"
	"time"
)

func fn() {
	const MAX = 10 * 1000 * 1000
	var m = make(map[int]string, MAX)
	for i := 0; i < MAX; i++ {
		m[i] = fmt.Sprintf("%4d", i)
	}

	time.Sleep(time.Second)
}

func main() {
	const MAX = 10 * 1000 * 1000
	var _map = make(map[int]string, MAX)
	for i := 0; i < MAX; i++ {
		_map[i] = fmt.Sprintf("%4d", i)
	}
	println("begin gc")
	//for i:=0; i<10; i++ {
	//  bs[i] = nil
	//}
	//time.Sleep(time.Second * 5)
	//runtime.GC()
	//time.Sleep(time.Second * 5)
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	fmt.Println("xxx", m.HeapObjects, m.HeapAlloc, m.TotalAlloc)
	for _, v := range m.BySize {
		fmt.Println("size ", v.Size, "malloc", v.Mallocs, "frees", v.Frees)
	}
}
