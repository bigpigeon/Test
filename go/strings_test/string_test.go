/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package strings_test

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	vv := strings.Split("", ",")
	t.Log(vv, len(vv))
}

func TestConvert(t *testing.T) {
	t.Logf(`"%s"`, string([]byte(nil)))
	t.Logf(`"%s"`, string([]byte{}))
}
