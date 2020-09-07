/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"fmt"
	"os"
	"time"
)

func panicFunc() {
	panic("panic test after 60 second")
}

func main() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Fprintln(os.Stdout, `{"level":"info","ts":1584010119.8292112,"caller":"v3@v3.1.0/default.go:34","msg":"test log","svcName":"XXX"}`)
		}
	}()
	time.Sleep(60 * time.Second)

}
