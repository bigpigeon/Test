/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import "github.com/bigpigeon/toyorm"
import "unsafe"
import "fmt"

type Product struct {
	toyorm.ModelDefault
	Data string
	Name string
}

var (
	TestVal int = 2
)

var GlobalProduct Product

var testFunc = func() string {
	return fmt.Sprintf("2222")
}

func main() {
	toy, err := toyorm.Open("", "")
	if err != nil {
		panic(err)
	}
	{
		// test comments

		model := toy.Model(Product{}).Limit(2).Debug().WhereGroup(toyorm.ExprEqual, map[uintptr]interface{}{
			unsafe.Offsetof(Product{}.ID): 2,
		})
		var model2 = model.Debug()
		model2.Swap()
		model = model.Debug()
		current := GlobalProduct
		if model != nil {
			fmt.Println("do something")
		}
		model.Find(&current)
		_ = model.Preload
		_ = unsafe.Offsetof(Product{}.ID)
		modelPreload := model.Preload
		fmt.Printf("select %v\n", modelPreload)
		modelPreload(unsafe.Offsetof(Product{}.Data))
		modelPreload("Data")
	}
	{
		var tab Product
		model := toy.Model(tab).Limit(2).Debug()
		model.Find(&Product{})
	}

	var a interface{} = 2
	fmt.Println(a)
	b := a.(int)
	fmt.Println(b)
}
