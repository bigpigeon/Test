/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package multi_return

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func multiReturn() (int, float32) {
	return 1, 2
}

func toSlice(a ...interface{}) []interface{} {
	return a
}

func TestMultiReturn(t *testing.T) {
	assert.Equal(t, []interface{}{int(1), float32(2)}, toSlice(multiReturn()))
}
