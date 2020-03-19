/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import "math/bits"

type CacheLines struct {
	data []byte
	way  int
}

func (c CacheLines) GetWay(n int) Way {
	return Way{
		data: c.data[n*64*8 : (n+1)*64*8],
		way:  c.way,
	}
}

type Cacheline struct {
	data []byte
	way  int
}

func (w Cacheline) Way(n int) {

}

type Way struct {
	data []byte
}

type Bits struct {
	data []byte
	size int
}

func NewBits(data []byte, offset, size int) Bits {
	return Bits{
		data: data[offset : offset+size],
		size: size,
	}
}

func NWayCache(n int) {
	l1 := make([]byte, 32*1024)
	cacheLines := CacheLines{data: l1, way: 8}

}

func main() {

}
