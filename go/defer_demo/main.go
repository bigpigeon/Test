/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("aa")
	}()
	panic("bb")
}
