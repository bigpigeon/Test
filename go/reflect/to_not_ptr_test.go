/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package reflect

import (
	"reflect"
	"testing"
)

func ToNotPoint(v interface{}) {
	toNotPoint(reflect.ValueOf(v))
}

func toNotPoint(vVal reflect.Value) {
	vTyp := vVal.Type()
	for vVal.Kind() == reflect.Ptr {
		if vVal.IsNil() {
			vVal.Set(reflect.New(vTyp.Elem()))
		}
		vVal = vVal.Elem()
		vTyp = vTyp.Elem()
	}
	if vTyp.Kind() == reflect.Struct {
		for i := 0; i < vVal.NumField(); i++ {
			vField := vVal.Field(i)
			toNotPoint(vField)
		}
	}
}

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
}
