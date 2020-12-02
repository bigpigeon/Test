/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import "fmt"

var V int

func F() { fmt.Printf("Hello, number %d\n", V) }

func F2(name string) {
	fmt.Printf("Hello, %s\n", name)
}
