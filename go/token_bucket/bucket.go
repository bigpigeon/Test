/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"sync/atomic"
	"time"
)

func TokenBucket(num int, interval time.Duration) func() bool {
	bucket := int64(num)
	go func() {
		// 每隔一秒往桶里+1
		tick := time.Tick(interval)
		for {
			select {
			case _, ok := <-tick:
				if ok == false {
					return
				}
				if bucket < int64(num) {
					atomic.AddInt64(&bucket, 1)
				}
			}
		}
	}()
	return func() bool {
		for {
			current := bucket
			// 使用cas尝试获取token
			if current > 0 {
				swap := atomic.CompareAndSwapInt64(&bucket, current, current-1)
				if swap == false {
					continue
				}
				return true
			} else {
				return false
			}

		}
	}
}
