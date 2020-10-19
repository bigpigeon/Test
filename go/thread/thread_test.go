/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package thread

import (
	"testing"
	"time"
)

func TestThread(t *testing.T) {
	goThreadBasic(func(callback func()) {
		v := 0
		time.Sleep(1 * time.Second)
		for i := 0; i < 100; i++ {
			v++
		}
		callback()
	})
	time.Sleep(2 * time.Second)
}

func TestCgoThread(t *testing.T) {
	testCgoThread(t)
}

func TestCgoThreadManager(t *testing.T) {
	testCgoThreadManager(t)
}

func TestCgoThreadManagerAndGoroutine(t *testing.T) {
	go testCgoThreadManager(t)
	goThreadBasic(func(callback func()) {
		v := 0
		time.Sleep(1 * time.Second)
		for i := 0; i < 100; i++ {
			v++
		}
		callback()
	})
	time.Sleep(100 * time.Second)
}
