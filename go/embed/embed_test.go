/*
 * Copyright 2021 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package embed

import _ "embed"
import "testing"

//go:embed hello.txt
var s string

func TestEmbed(t *testing.T) {
	t.Log(s)
}
