/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"fmt"
	"runtime"
)

type stack []uintptr

func caller() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(1, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

func main() {

	st := caller()
	frames := runtime.CallersFrames(*st)
	for {
		f, more := frames.Next()
		if more == false {
			break
		}
		fmt.Printf("%s:%d \n in func %s \n", f.File, f.Line, f.Function)
		//if f.Func != nil {
		//	file, line := f.Func.FileLine(f.PC - 1)
		//	fmt.Printf("  in func %s line %d file %s\n", f.Func.Name(), line, file)
		//}
	}

}
