/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"github.com/bigpigeon/Test/go/module_demo/depend1"
	"github.com/bigpigeon/Test/go/module_demo/depend2"
	"github.com/bigpigeon/Test/go/module_demo/depend3"
	"reflect"
)

func main() {
	fmt.Printf("%#v\n", depend1.DB)
	fmt.Printf("%#v\n", depend2.DB)
	fmt.Printf("%v\n", reflect.TypeOf(depend1.DB) == reflect.TypeOf(depend2.DB))
	depend3.SomeOutput()
}
