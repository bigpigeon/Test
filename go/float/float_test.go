/*
 * Copyright 2018 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package float

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"
	"unsafe"
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

}

func F32ToBin(v float32) string {
	ival := (*uint32)(unsafe.Pointer(&v))
	return fmt.Sprintf("%.32b", *ival)
}

func BinToF32(s string) float32 {

	v, err := strconv.ParseUint(s, 2, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.32b\n", v)
	fval := (*float32)(unsafe.Pointer(&v))
	return *fval
}

func TestFraction(t *testing.T) {

	for i := -10; i < 10; i++ {
		fmt.Println(float32(i), "\t", F32ToBin(float32(i)))
	}

	for i := 0; i < 8388608; i += 10000 {
		fmt.Println(float32(i), "\t", F32ToBin(float32(i)))
	}
}
