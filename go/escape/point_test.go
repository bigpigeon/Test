/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package escape

import (
	"testing"
)

type EscapeData struct {
	Int *int
}

func EscapeProcess(val *int) {
	ed := &EscapeData{}
	ed.Int = val
}

func NoEscapeProcess() {
	ed := &EscapeData{}
	val := 22
	ed.Int = &val
}

func TestPoint(t *testing.T) {
	x := 22
	px := &x
	EscapeProcess(px)
	NoEscapeProcess()
	//fmt.Println(data)
}
