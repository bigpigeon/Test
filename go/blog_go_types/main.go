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
	"reflect"
	"sort"
)

// 排序规则order by Pos(), End()
func sortNodes(nodes []ast.Node) {
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Pos() == nodes[j].Pos() {
			return nodes[i].End() < nodes[j].End()
		}
		return nodes[i].Pos() < nodes[j].Pos()
	})
}

// map中的元素是无序的，对key排序打印更好查看
func getSortedKeys(m interface{}) []ast.Node {
	mValue := reflect.ValueOf(m)
	nodes := make([]ast.Node, mValue.Len())
	keys := mValue.MapKeys()
	for i := range keys {
		nodes[i] = keys[i].Interface().(ast.Node)
	}
	sortNodes(nodes)
	return nodes
}

func main() {
	fset := token.NewFileSet() // 位置是相对于节点
	// 用ParseFile把文件解析成*ast.File节点
	f, err := parser.ParseFile(fset, "data/src.go.temp", nil, 0)
	if err != nil {
		panic(err)
	}

	// 使用types check
	// 构造config
	config := types.Config{
		// 加载包的方式，可以通过源码或编译好的包，其中编译好的包分为gc和gccgo,前者应该是
		Importer:    importer.For("source", nil),
		FakeImportC: true,
	}

	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
		InitOrder:  make([]*types.Initializer, 0, 0),
	}
	_, err = config.Check("", fset, []*ast.File{f}, info)
	if err != nil {
		panic(err)
	}
	// 打印types
	fmt.Println("------------ types -----------")
	for _, node := range getSortedKeys(info.Types) {
		expr := node.(ast.Expr)
		typeValue := info.Types[expr]
		fmt.Printf("%s - %s it's value %v type %s\n",
			fset.Position(expr.Pos()),
			fset.Position(expr.End()),
			typeValue.Value,
			typeValue.Type.String(),
		)
	}
	// 打印defs
	fmt.Println("------------ def -----------")
	for _, node := range getSortedKeys(info.Defs) {
		ident := node.(*ast.Ident)
		object := info.Defs[ident]
		fmt.Printf("%s - %s",
			fset.Position(ident.Pos()),
			fset.Position(ident.End()),
		)
		if object != nil {
			fmt.Printf(" it's object %s type %s",
				object,
				object.Type().String(),
			)
		}
		fmt.Println()
	}
	// 打印Uses
	fmt.Println("------------ uses -----------")
	for _, node := range getSortedKeys(info.Uses) {
		ident := node.(*ast.Ident)
		object := info.Uses[ident]
		fmt.Printf("%s - %s",
			fset.Position(ident.Pos()),
			fset.Position(ident.End()),
		)
		if object != nil {
			fmt.Printf(" it's object %s type %s",
				object,
				object.Type().String(),
			)
		}
		fmt.Println()
	}
	// 打印Implicits
	fmt.Println("------------ implicits -----------")
	for _, node := range getSortedKeys(info.Implicits) {
		object := info.Implicits[node]
		fmt.Printf("%s - %s it's object %s\n",
			fset.Position(node.Pos()),
			fset.Position(node.End()),
			object,
		)
	}
	// 打印Selections
	fmt.Println("------------ selections -----------")
	for _, node := range getSortedKeys(info.Selections) {
		sel := node.(*ast.SelectorExpr)
		typeSel := info.Selections[sel]
		fmt.Printf("%s - %s it's recv %s kind %v object %s type %s\n",
			fset.Position(sel.Pos()),
			fset.Position(sel.End()),
			typeSel.Recv(),
			typeSel.Kind(),
			typeSel.String(),
			typeSel.Type().String(),
		)
	}
	// 打印Scopes
	fmt.Println("------------ scopes -----------")
	for _, node := range getSortedKeys(info.Scopes) {
		scope := info.Scopes[node]
		fmt.Printf("%s - %s it's scope %s\n",
			fset.Position(node.Pos()),
			fset.Position(node.End()),
			scope.String(),
		)
	}
	// 打印InitOrder
	fmt.Println("------------ init -----------")
	for _, init := range info.InitOrder {
		fmt.Printf("init %s\n",
			init.String(),
		)
	}
}
