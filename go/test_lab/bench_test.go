/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package test_lab

import (
	"testing"
)

func TestToNotPtr(t *testing.T) {
	type TestSub struct {
		Data   *string
		Name   *string
		IntVal *int
	}
	type TestData struct {
		Sub  *TestSub
		Text *string
	}

	var data *TestData
	ToNotPoint(&data)
	t.Logf("data %#v\n", *data)
	t.Logf("data %#v\n", *data.Sub)
	t.Logf("data %#v\n", *data.Text)
	t.Logf("data %#v\n", *data.Sub.Data)
	t.Logf("data %#v\n", *data.Sub.Name)
	t.Logf("data %#v\n", *data.Sub.IntVal)
}

func BenchmarkToNotPtr(b *testing.B) {
	// 函数的一些比较耗时的初始化内容写在这
	// 开启计时器
	b.StartTimer()
	// b.N 就是调用次数，需要自己组织for循环
	// 基准测试会不断的增加调用次数，直到调用耗时超过1秒
	b.Log(b.N)
	for i := 0; i < b.N; i++ {
		type TestSub struct {
			Data   *string
			Name   *string
			IntVal *int
		}
		type TestData struct {
			Sub  *TestSub
			Text *string
		}

		var data *TestData
		ToNotPoint(&data)
	}
	b.StopTimer()
	// 资源释放在计时器停止之后

}
