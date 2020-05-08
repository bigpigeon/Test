/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package _map

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[int]int{}
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	for k, v := range m {
		require.Equal(t, m[k], v)
		delete(m, k)
	}
}

func TestMapIter(t *testing.T) {
	m := map[int]int{}
	for i := 0; i < 10; i++ {
		m[i] = i
	}
	for k, v := range m {
		t.Log(k, v)
	}
	t.Log("second")
	for k, v := range m {
		t.Log(k, v)
	}
}
