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

func printDp(pat string, dp [][]int) {
	fmt.Print("  ")
	for _, s := range pat {
		fmt.Printf("%c ", s)
	}
	fmt.Println()
	for i := 0; i < 'z'-'a'+1; i++ {
		fmt.Printf("%c ", byte(i+'a'))
		for _, v := range dp {
			fmt.Printf("%d ", v[i])
		}
		fmt.Println()
	}

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
		for c := 0; c < 'z'-'a'+1; c++ {
			if pat[i] == byte(c)+'a' {
				kmp.dp[i][c] = i + 1
			} else {
				kmp.dp[i][c] = kmp.dp[x][c]
			}
		}
		// example world abcabc, when iter 4'th index word ,x become 1
		x = kmp.dp[x][pat[i]-'a']
	}
	printDp(pat, kmp.dp)
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
