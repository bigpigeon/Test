/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package sync_pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func putMap(pool *sync.Pool) {
	v := 2
	pool.Put(&v)
	pool.Put(&v)
	pool.Put(&v)
}

func TestSyncPool(t *testing.T) {
	// 初始化一个pool
	pool := &sync.Pool{}
	v := 1
	pool.Put(&v)
	pool.Put(&v)
	pool.Put(&v)

	fmt.Printf("%v\n", pool.Get())
	fmt.Printf("%v\n", pool.Get())
	fmt.Printf("%v\n", pool.Get())
	fmt.Printf("%v\n", pool.Get())
	putMap(pool)
	runtime.GC()
	time.Sleep(1 * time.Second)
	fmt.Printf("%v\n", pool.Get())
	fmt.Printf("%v\n", pool.Get())
	fmt.Printf("%v\n", pool.Get())
	fmt.Printf("%v\n", pool.Get())
}
