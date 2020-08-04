/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package reflect

import (
	"reflect"
	"testing"
)

func GetPointElemAndNew(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		v = v.Elem()
	}
	return v
}

func TestAddressable(t *testing.T) {
	type Point struct {
		X int
		Y int
	}
	type Data struct {
		P *Point
	}

	val := reflect.ValueOf(&Data{})
	val = GetPointElemAndNew(val)
	pVal := val.Field(0)
	t.Log("P field addressable ? ", pVal.CanAddr())
	pVal = GetPointElemAndNew(pVal)
	t.Logf("%#v\n", pVal.Interface())
}
