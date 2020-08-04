/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package enum

import (
	"fmt"
	"testing"
)

const (
	a = iota
	b
	c = "agent"
	d = "go"
	e
)

func TestIota(t *testing.T) {
	t.Log(e)
}

// 下列函数输出
func TestDefer(t *testing.T) {
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	panic("err")
}
