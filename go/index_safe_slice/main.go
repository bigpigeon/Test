/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import "fmt"

func safeSliceAccessInt(sl []int, idx int) (int, bool) {
	defer func() { recover() }()
	return sl[idx], true // default values will be substituted if sl[idx] panics (0, false)
}

func main() {
	a := []int{1, 2}
	b, bool := safeSliceAccessInt(a, 3)
	fmt.Println(b, bool)
}
