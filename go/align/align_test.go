/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package align

import (
	"testing"
	"unsafe"
)

func TestAlign(t *testing.T) {
	type NotAlignData struct {
		A bool
		B int32
		C int8
		D int64
		E byte
	}
	v := NotAlignData{}
	t.Log(unsafe.Sizeof(v), unsafe.Alignof(v))

	type AlignData struct {
		A bool
		C int8
		E byte
		B int32
		D int64
	}
	v2 := AlignData{}
	t.Log(unsafe.Sizeof(v2), unsafe.Alignof(v2))
}
