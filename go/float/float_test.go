/*
 * Copyright 2018 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package float

import (
	"math"
	"reflect"
	"testing"
)

func TestFloatRound(t *testing.T) {
	a := 10.123333
	t.Log(a)
	t.Log(math.Round(a))
}

func TestConst(t *testing.T) {
	t.Log(50 * 0.5)
}

func TestInf(t *testing.T) {
	a := float32(1) / reflect.ValueOf(float32(0)).Interface().(float32)
	t.Logf("val %f is Nan? %v", a, a != a)
	b := float32(2)
	t.Logf("val %f is Nan? %v", b, b != b)
}

func TestNaN(t *testing.T) {
	a := math.Float32frombits(0xffc00000)
	t.Logf("val %f is Nan? %v", a, a != a)
	b := math.Float64frombits(0xfff8000000000000)
	t.Logf("val %f is NaN? %v", b, b != b)
}

func TestNegativeZero(t *testing.T) {
	a := reflect.ValueOf(float64(0)).Interface().(float64) / reflect.ValueOf(float64(-1.0)).Interface().(float64)
	b := float64(0)
	t.Log("a:", a, "b:", b)
	t.Log("a < b ?", a < b)
	t.Log("a == b ?", a == b)
	t.Log("a is -0 ?", a == -0)
	t.Log("b is -0 ?", b == -0)
	v := map[float64]string{}
	v[a] = "negative zero"
	v[b] = "positive zero"
	t.Log(v)
}
