/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package _chan

import (
	"fmt"
	"testing"
	"time"
)

func ignoreReload(reloadChan chan bool, timeout *time.Timer) {
	for {
		select {
		case _, ok := <-reloadChan:
			if ok == false {
				return
			}
			fmt.Println("ignore once")
		case <-timeout.C:
			fmt.Println("ignore timeout")
			return
		}

	}
}

func TestInterval(t *testing.T) {

	reloadChan := make(chan bool, 10)
	go func() {
		reloadChan <- true
		reloadChan <- true
		reloadChan <- true
		reloadChan <- true
		reloadChan <- true
		time.Sleep(time.Second)
		fmt.Println("push 8")
		reloadChan <- true

		close(reloadChan)
	}()
	var previousReload time.Time
	for {
		select {
		case _, ok := <-reloadChan:
			if ok == false {
				return
			}
			ignoreReload(reloadChan, time.NewTimer(time.Second))
			previousReload = time.Now()
			fmt.Println("reload once", previousReload)
		}
	}
}
