/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import "fmt"

type Resource interface {
	Name() string
}

type Product interface {
	Resource
	Price() int
}

type StuffFood struct {
	name  string
	price int
}

func (s *StuffFood) Name() string {
	return s.name
}

func (s *StuffFood) Price() int {
	return s.price
}

func NamePrint(res interface{}) {
	switch x := res.(type) {
	case []Product:
		for _, r := range x {
			fmt.Println(r)
		}
	case []Resource:
		for _, r := range x {
			fmt.Println(r)
		}
	}

}

func main() {
	foods := []Product{
		&StuffFood{"rice", 2},
		&StuffFood{"noodle", 3},
		&StuffFood{"bread", 1},
	}
	// cannot convert
	NamePrint(foods)
}
