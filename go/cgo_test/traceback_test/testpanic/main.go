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

import _ "github.com/ianlancetaylor/cgosymbolizer"

func main() {
	C.panic()
}
