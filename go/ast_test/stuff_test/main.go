/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"encoding/json"
	"fmt"
)

func Foo() {

}

func main() {
	Foo()
	m := MainStuff{}
	m.Stuff()
	s := Stuff{}
	s.Stuff()
	e := json.Encoder{}
	e.Encode(s)

}

type MainStuff struct {
}

func (i MainStuff) Stuff() {
	fmt.Printf("stuff")
}
