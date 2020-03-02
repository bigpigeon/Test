/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"fmt"
	"github.com/bigpigeon/Test/go/runtime/dep"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if ok == false {
		panic("not support read build info")
	}
	fmt.Printf("%#v\n", info)
	dep.PrintBuildInfo()
}
