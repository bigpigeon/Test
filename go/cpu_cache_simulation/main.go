/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"fmt"
)

type CacheLines struct {
	data []byte
	ways int
}

func (c CacheLines) Get(n int) CacheLine {
	return CacheLine{
		data: c.data[n*64*c.ways : (n+1)*64*c.ways],
		way:  c.ways,
	}
}

type CacheLine struct {
	data []byte
	way  int
}

func (w CacheLine) Way(n int) Way {
	preWayLen := len(w.data) / w.way
	d := w.data[n*preWayLen : (n+1)*preWayLen]
	fmt.Println(len(d))
	return Way{
		tag:    0,
		index:  0,
		offset: 0,
	}
}

type Way struct {
	tag    int8
	index  int8
	offset int8
}

func NWayCache() CacheLines {
	l1 := make([]byte, 8*1024)
	cacheLines := CacheLines{data: l1, ways: 4}

	return cacheLines
}

func main() {
	lines := NWayCache()
	way := lines.Get(0).Way(0)
	fmt.Printf("way %#v", way)
}
