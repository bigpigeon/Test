/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"testing"
	"time"
)

func resultToStr(r []bool) string {
	d := make([]byte, len(r))
	for i, v := range r {
		if v {
			d[i] = '1'
		} else {
			d[i] = '0'
		}
	}
	return string(d)
}

func TestTokenBucket(t *testing.T) {
	bucket := TokenBucket(100, time.Millisecond*10)
	result := make([]bool, 1000)
	for i := 0; i < 1000; i++ {
		result[i] = bucket()
		time.Sleep(time.Millisecond)
	}
	t.Log(resultToStr(result))
}
