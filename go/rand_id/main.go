/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

var filterID = []int{4444, 4445, 4446, 4447, 4449, 5555}

func init() {
	sort.Ints(filterID)
}

func skipID(id int) int {
	if pos := sort.SearchInts(filterID, id); pos != -1 {
		// calculate offset
		offset := -2
		for i := pos; i >= 0; i-- {
			fmt.Println("previous", filterID[i])
			if id != filterID[i]+i {
				break
			}
			offset += 2
		}
		for i := pos; i < len(filterID); i++ {
			fmt.Println("after", filterID[i])
			if id != filterID[i]-i {
				break
			}
			offset++
		}
		return skipID(id + offset)
	}
	return id
}

func randID(r int) int {
	return skipID(rand.Intn(r - len(filterID)))
}

func main() {
	for i := 4444; i < 4450; i++ {
		fmt.Println(skipID(i))
	}
}
