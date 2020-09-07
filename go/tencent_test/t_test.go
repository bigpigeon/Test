/*
 * Copyright 2020. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

/*
* 给定一个整数数组nums和一个目标值target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
* 假设每种输入只会对应一个答案。但是，不能重复利用这个数组中同样的元素。
*
* 示例:
* 给定 nums = [2, 7, 11, 15], target = 9
* 因为 nums[0] + nums[1] = 2 + 7 = 9
* 所以返回 [0, 1]
 */

package tencent_test

import (
	"encoding/json"
	"math"
	"os"
	"testing"
)

func findNumberSum(nums []int, target int) (int, int) {
	if len(nums) < 2 {
		return -1, -1
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			return left, right

		}
	}
	return -1, -1

}

func Test1(t *testing.T) {
	t.Log(findNumberSum([]int{2, 7, 11, 15}, 9))
	t.Log(findNumberSum([]int{2, 7, 11, 15}, 18))
}

/*
实现函数int  atoi(string & src, int & result),字符串转换为整数   例如：
 "+123"->123,
"-123"-> -123,
 "abs"->0,
"123nndnd34"->123
如果超出整型范围，返回-1
*/

func Atoi(s string) int {
	if len(s) < 1 {
		return -1
	}
	sign := 1
	val := 0
	if s[0] == '-' {
		s = s[1:]
		sign = -1
	} else if s[0] == '+' {
		s = s[1:]
	}
	maxInt := math.MaxInt64
	for _, c := range s {
		if c < '9' && c > '0' {

			if maxInt-val*10 < int(c)-'9' {
				return -1
			}
			val = val*10 + int(c) - '0'

		} else {
			break
		}
	}
	return sign * val
}

func Test2(t *testing.T) {
	t.Log(Atoi("-123"))
	t.Log(Atoi("+123"))
	t.Log(Atoi("abs"))
	t.Log(Atoi("123nndnd34"))
	t.Log(Atoi("999999999999999"))
}

/*
用ch将strTarget分割成数组，保存到vctStr
	void Split(const std::string &strTarget, char ch, std::vector<std::string> &vctStr);
*/

func SplitStr(str string, spl byte) []string {
	val := []string{}
	start := 0
	for i, c := range str {
		if byte(c) == spl {
			val = append(val, str[start:i])
			start = i + 1
		}
	}
	if start != len(str) {
		val = append(val, str[start:])
	}
	return val
}

func PrintStrArr(s []string) {
	json.NewEncoder(os.Stdout).Encode(s)
}

func Test3(t *testing.T) {
	PrintStrArr(SplitStr("1-2-3-4", '-'))
	PrintStrArr(SplitStr("1---4", '-'))
	PrintStrArr(SplitStr("", '-'))
}

func findNum(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		if arr[0] == target {
			return 0
		}
	}

	mid := len(arr) / 2
	if arr[mid] < target {
		pos := findNum(arr[mid+1:], target)
		if pos != -1 {
			return pos + mid + 1
		}
		return -1
	} else if arr[mid] > target {
		pos := findNum(arr[:mid], target)
		if pos != -1 {
			return pos
		}
		return -1
	} else {
		return mid
	}

}

func Test4(t *testing.T) {
	t.Log(findNum([]int{1, 3, 5, 7, 9, 11, 13}, 10))
	t.Log(findNum([]int{1, 3, 5, 7, 9, 11, 13}, 11))
	t.Log(findNum([]int{1, 3, 5, 7, 9, 11, 13}, 3))
}
