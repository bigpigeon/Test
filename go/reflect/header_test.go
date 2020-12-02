/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package reflect

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestSliceHeader(t *testing.T) {
	months := []string{1: "January", 2: "February"}
	t.Logf("months fetch      addr %p,\n", &months)
	t.Logf("months            addr  %p,\n", months)
	t.Logf("months first elem addr  %p,\n", &months[0])

	header := (*reflect.SliceHeader)(unsafe.Pointer(&months))
	t.Logf("header          addr %p,\n", header)
	t.Logf("len             addr %p\n", &header.Len)
	t.Logf("cap             addr %p\n", &header.Cap)
	t.Logf("data            addr 0x%x,\n", header.Data)
	t.Logf("data first elem addr 0x%x,\n", header.Data+0)
}

func TestStringHeader(t *testing.T) {
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 128)
	}

	str := string(data)
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	dataHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	t.Logf("str header addr %p,\n", strHeader)
	t.Logf("str data  0x%x\n", strHeader.Data)
	t.Logf("data header addr %p\n", dataHeader)
	t.Logf("byte data 0x%x\n", dataHeader.Data)
}
