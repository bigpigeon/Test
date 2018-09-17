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
	fset := token.NewFileSet() // 位置是相对于节点
	// 用ParseFile把文件解析成*ast.File节点
	f, err := parser.ParseFile(fset, "data/main.go", nil, 0)
	if err != nil {
		panic(err)
	}

	// 使用types check
	// 构造config
	config := types.Config{
		// 加载包的方式，可以通过源码或编译好的包，其中编译好的包分为gc和gccgo,前者应该是
		Importer: importer.For("source", nil),
		// 表示允许包里面加载c库 import "c"
		FakeImportC: true,
	}

	info := &types.Info{
		// 表达式对应的类型
		Types: make(map[ast.Expr]types.TypeAndValue),
		// 被定义的标示符
		Defs: make(map[*ast.Ident]types.Object),
		// 被使用的标示符
		Uses: make(map[*ast.Ident]types.Object),
		// 隐藏节点，匿名import包，type-specific时的case对应的当前类型，声明函数的匿名参数如var func(int)
		Implicits: make(map[ast.Node]types.Object),
		// 选择器,只能针对类型/对象.字段/method的选择，package.API这种不会记录在这里
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		// scope 记录当前库scope下的所有域，*ast.File/*ast.FuncType/... 都属于scope，详情看Scopes说明
		// scope关系: 最外层Universe scope,之后Package scope，其他子scope
		Scopes: make(map[ast.Node]*types.Scope),
		// 记录所有package级的初始化值
		InitOrder: make([]*types.Initializer, 0, 0),
	}
	// 这里Check的第一个path参数是当前pkg前缀，和FileSet的文件路径是无关的
	pkg, err := config.Check("", fset, []*ast.File{f}, info)
	if err != nil {
		panic(err)
	}
	for _, name := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(name)
		fmt.Printf("%s name %s type %s %T\n", fset.Position(obj.Pos()), name, obj.Type(), obj.Type())
		switch x := obj.Type().(type) {
		case *types.Array:
			fmt.Printf("array %s len %d elem %s \n", name, x.Len(), x.Elem())
		case *types.Basic:
			fmt.Printf("basic %s kind %d info %v \n", name, x.Kind(), x.Info())
		case *types.Chan:
			fmt.Printf("chan %s dir %v elem %v \n", name, x.Dir(), x.Elem())
		case *types.Interface:
			fmt.Printf("interface %s  it's method \n", name)
			for i := 0; i < x.NumMethods(); i++ {
				method := x.Method(i)
				fmt.Printf("%s\n", method)
			}
			fmt.Println()
		case *types.Map:
			fmt.Printf("map %s  it's key %s it's elem %s \n", name, x.Key(), x.Elem())
		case *types.Named:
			fmt.Printf("named %s it's underlying %s it's method \n", name, x.Underlying())
			for i := 0; i < x.NumMethods(); i++ {
				method := x.Method(i)
				fmt.Printf("%s\n", method)
			}
		case *types.Pointer:
			fmt.Printf("point %s it's elem %s\n", name, x.Elem())
		case *types.Signature:
			fmt.Printf("func %s params %s %T\n", name, x.Params(), x.Params())
			fmt.Printf("func %s result %s %T\n", name, x.Results(), x.Results())
		case *types.Slice:
			fmt.Printf("slice %s elem %s \n", name, x.Elem())
		case *types.Struct:
			fmt.Printf("struct %s it's fields \n", name)
			for i := 0; i < x.NumFields(); i++ {
				field := x.Field(i)
				fmt.Printf("%s\n", field)
			}
		}
	}
	// use Identical to compare type
	idName1 := pkg.Scope().Lookup("idName1").Type()
	idName2 := pkg.Scope().Lookup("idName2").Type()
	idName3 := pkg.Scope().Lookup("idName3").Type()
	idName4 := pkg.Scope().Lookup("idName4").Type()
	fmt.Printf("idName1 %s idName2 %s they is same type %v\n", idName1, idName2, types.Identical(idName1, idName2))
	fmt.Printf("idName1 %s idName3 %s they is same type %v\n", idName1, idName3, types.Identical(idName1, idName3))
	fmt.Printf("idName1 %s idName4 %s they is same type %v\n", idName1, idName4, types.Identical(idName1, idName4))

	intVal := pkg.Scope().Lookup("intVal").Type()
	intVal2 := pkg.Scope().Lookup("intVal2").Type()
	fmt.Printf("intVal %s intVal2 %s they is same type %v\n", intVal, intVal2, types.Identical(intVal, intVal2))
	// use Implicate to match type and interface
	stringer := pkg.Scope().Lookup("Stringer").Type().Underlying().(*types.Interface)
	response := pkg.Scope().Lookup("Response").Type()
	fmt.Printf("type %s is implicate interface %s? %v\n", response, stringer, types.Implements(response, stringer))
}
