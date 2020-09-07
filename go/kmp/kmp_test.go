/*
 * Copyright 2020. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package kmp

import (
	"fmt"
	"testing"
)

func TestNewKmp(t *testing.T) {
	//kmp := NewKmp("abababc")
	//t.Log(kmp.dp)
	{
		kmp := NewKmp("ababc")
		t.Log(kmp.Query("abababc"))
	}
}

func printKmp(pat string, v [][]int) {
	charList := []byte{}
	for i := 'a'; i <= 'z'; i++ {
		charList = append(charList, byte(i))
	}
	fmt.Println("  ", string(charList))
	for _, vv := range v {
		fmt.Println(" ", vv)

	}

}
