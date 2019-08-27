/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

func main() {
	c := make(chan int)
	close(c)
	go func() {
		<-c
		receive := 0
		send := 0
		for i := 0; i < 1000; i++ {
			select {
			case <-c:
				receive++
			case c <- 2:
				send++
			}
		}
	}()
}
