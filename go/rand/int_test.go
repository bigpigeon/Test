/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package rand

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestIntN(t *testing.T) {
	count := map[int]int{}
	for i := 0; i < 1000; i++ {
		count[rand.Intn(2)]++
	}
	fmt.Println(count)
}
