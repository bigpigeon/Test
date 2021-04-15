/*
 * Copyright 2021 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

/*
extern void panic(void);
*/
import "C"

import (
	"fmt"
	_ "github.com/bigpigeon/Test/go/cgo_test/traceback_test/testpanic/cgosymbolizer"
)

func main() {
	fmt.Println("sth")
	C.panic()
}
