/*
 * Copyright 2018 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package copy

import (
	"fmt"
	"testing"
)

func TestPointCopy(t *testing.T) {
	{
		a := new(int)
		*a = 2
		b := &(*a)
		*b = 3
		t.Log(a == b)
	}
	{
		a := new(int)
		*a = 2
		b := new(int)
		*b = *a
		t.Log(a == b)
	}
}

func TestStrCopy(t *testing.T) {
	str := "123"
	strB := make([]byte, len(str))
	copy(strB, str)
	strC := str
	strBS := string(strB)
	fmt.Printf("a %p b %p c%p\n", &str, &strBS, &strC)
}
