/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import "plugin"

func main() {
	p, err := plugin.Open("myplugin/myplugin.so")
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*int) = 7
	f.(func())() // prints "Hello, number 7"
	{
		f2, err := p.Lookup("F2")
		if err != nil {
			panic(err)
		}
		f2.(func(name string))("jia")
	}
}
