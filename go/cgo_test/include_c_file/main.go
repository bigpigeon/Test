/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

// #include "clib/hello.c"
import "C"

func main() {
	C.Hello()
}
