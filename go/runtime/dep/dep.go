/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package dep

import (
	"fmt"
	"runtime/debug"
)

func PrintBuildInfo() {
	info, ok := debug.ReadBuildInfo()
	if ok == false {
		panic("not support read build info")
	}
	fmt.Printf("%#v\n", info)
	dep.Read
}
