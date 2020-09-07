/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"fmt"
	"github.com/bigpigeon/Test/go/ldflags/global_var_replace/submodule"
)

func main() {
	fmt.Printf("replace variable '%s'\n", submodule.ReplaceVal)
}
