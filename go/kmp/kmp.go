/*
 * Copyright 2020. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package kmp

import "fmt"

type Kmp struct {
	Pat string
	dp  [][]int
}

func NewKmp(pat string) *Kmp {
	kmp := &Kmp{
		Pat: pat,
		dp:  make([][]int, len(pat)),
	}
	for i := range kmp.dp {
		kmp.dp[i] = make([]int, 'z'-'a'+1)
	}
	kmp.dp[0][pat[0]-'a'] = 1
	x := 0

	for i := 1; i < len(pat); i++ {
		fmt.Println("x=", x)
		for c := 0; c < 'z'-'a'+1; c++ {
			if pat[i] == byte(c)+'a' {
				kmp.dp[i][c] = i + 1
			} else {
				kmp.dp[i][c] = kmp.dp[x][c]
			}
		}
		x = kmp.dp[x][pat[i]-'a']
	}
	return kmp
}

func (k Kmp) Query(s string) int {
	current := 0
	for i, l := range s {
		current = k.dp[current][l-'a']
		if current == len(k.Pat) {
			return i - current + 1
		}
	}
	return -1
}
