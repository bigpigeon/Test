/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package stack

import (
	"fmt"
	"runtime/debug"
	"sync"
	"testing"
)

func TestGetStack(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("test")
		debug.PrintStack()
	}()
	wg.Wait()
}
