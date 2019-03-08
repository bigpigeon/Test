/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"crypto/rand"
	"testing"
)

var s string

func init() {
	bytes := make([]byte, 256*100000)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	s = string(bytes)
}
func BenchmarkSlice(b *testing.B) {
	a := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		a[i] = s[i]
	}
}

func BenchmarkSliceMemclr(b *testing.B) {
	a := make([]byte, len(s))
	for i := range s {
		a[i] = s[i]
	}
}
