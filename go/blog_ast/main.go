/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := `
package main

// 该声明为GenDecl TOK=token.IMPORT
import "fmt"

// 该声明为GenDecl TOK=token.TYPE
type Product struct {
	Name string
}

// 该声明为GenDecl TOK=token.VAR
var product Product

// 该声明为FunDecl
func main() { //test1
	fmt.Println("Hello, World!") //test2
	a := []int{1,2,3}
	a[1],a[2] = 5,6
}
`
	fset := token.NewFileSet() // 位置是相对于节点
	// 用ParseFile把文件解析成*ast.File节点
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// 打印ast节点
	ast.Print(fset, f)

}
