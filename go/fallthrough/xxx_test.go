/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package _fallthrough

import (
	"testing"
)

func TestSwitch(t *testing.T) {
	v := "jia"
	switch v {
	case "jia":
		fallthrough
	case "val":
		t.Log("test")
	}
}
