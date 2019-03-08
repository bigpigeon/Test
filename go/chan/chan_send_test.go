/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package _chan

import "testing"

func TestChanSend(t *testing.T) {
	c := make(chan bool, 1)

	for i := 0; i < 2; i++ {
		select {
		case c <- true:
			t.Logf("the %d's chan sent scuess", i)
		default:
			t.Logf("the %d's chan sent failure", i)
		}
	}
}
