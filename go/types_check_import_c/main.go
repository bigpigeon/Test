/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

func main() {
	src := `
package main

import (
	"C"
	"fmt"
)

const cint1 C.int = 0xff

func main() {
	var cint2 C.int = 22
	var cint3 C.int = 0xff
	fmt.Printf("%#v\n", cint1)
	fmt.Printf("%#v\n", cint2)
	fmt.Printf("%#v\n", cint3)
}
`
	fs := token.NewFileSet()
	parserFile, err := parser.ParseFile(fs, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	config := types.Config{Importer: importer.For("source", nil), FakeImportC: true}
	_, err = config.Check("", fs, []*ast.File{parserFile}, info)
	if err != nil {
		panic(err)
	}
	for expr, t := range info.Types {
		fmt.Printf("expr: %s type %s\n", fs.Position(expr.Pos()), t.Type)
	}
}
