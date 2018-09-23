/*
 * Copyright 2018 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package float

import (
	"math"
	"testing"
)

func TestFloatRound(t *testing.T) {
	a := 10.123333
	t.Log(a)
	t.Log(math.Round(a))
}
