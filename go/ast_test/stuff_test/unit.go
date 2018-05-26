/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
)

type Stuff struct {
}

func (s Stuff) Stuff() {
	fmt.Printf("stuff")
}
