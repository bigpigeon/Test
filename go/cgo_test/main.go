/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"C"
	"fmt"
)

const cint1 C.int = 0xff

func main() {
	var cint2 C.int = 22
	var cint3 C.int = 0xff
	fmt.Printf("%#v\n", cint1)
	fmt.Printf("%#v\n", cint2)
	fmt.Printf("%#v\n", cint3)
}
