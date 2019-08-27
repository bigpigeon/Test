/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package test_lab

import "reflect"

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
