/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeSwitch(t *testing.T) {
	v1 := 1
	v2 := "2"
	v3 := struct {
		Val  int
		Name string
	}{3, "3"}

	fn := func(v interface{}) {
		switch val := v.(type) {
		case int:
			fmt.Println("int", val)
		case string:
			fmt.Println("string", val)
		case reflect.Value:
			fmt.Println("valuer", val)
		}
	}
	fn(v1)
	fn(v2)
	fn(v3)

}
