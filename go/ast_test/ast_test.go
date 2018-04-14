/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package ast_test

import (
	"go/ast"
	"testing"
	//"go/format"
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"strings"
)

func TestCustom(t *testing.T) {
	// src is the input for which we want to print the AST.
	src := `
/*
multiple group
*/
package main
func main() { //test1 //test2
/*
multiple group1
multiple group2
*/
	println("Hello, World!") //test2
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0|parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)
	// Output:
	//
}

func TestParser(t *testing.T) {
	src := `
package main

import (
	"fmt"
)

type Printer interface {
	Print()
}

type myLogger struct {
	name string
}

//test print
func (log *myLogger) Print() { //xxx
	// why
	fmt.Printf("test")
}

func main() {
	loggerMap := map[string]*myLogger{}
	var printer Printer = loggerMap["not exist"]
	if printer == nil {
		fmt.Print("yes printer is nil")
	} else {
		fmt.Print("now have a bug")
	}
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0|parser.ParseComments)
	if err != nil {
		panic(err)
	}
	//ast.Inspect(f, func(node ast.Node) bool {
	//
	//})
	cmap := ast.NewCommentMap(fset, f, f.Comments)
	//ast.Print(fset, f)

	for n, _ := range cmap {
		ast.Print(fset, n)
	}
	for _, decl := range f.Decls {
		if _func, ok := decl.(*ast.FuncDecl); ok {
			if comments := cmap[_func]; len(comments) != 0 {
				fmt.Println("func & comments ")
				ast.Print(fset, _func)
				ast.Print(fset, comments)
			}
		}
		if _gen, ok := decl.(*ast.GenDecl); ok {
			if comments := cmap[_gen]; len(comments) != 0 {
				fmt.Println("gen & comments ")
				ast.Print(fset, _gen)
				ast.Print(fset, comments)
			}
		}
	}
}

//
//func parserFunc(t *testing.T,decl *ast.FuncDecl){
//
//}

func NodeStr(node ast.Node) string {
	switch x := node.(type) {
	case *ast.Ellipsis:
		return fmt.Sprintf("ellipsis %v", x.Elt)
	case *ast.GenDecl:
		return fmt.Sprintf("gen decl %v", x.Doc)
	case *ast.ImportSpec:
		return fmt.Sprintf("import %v", x.Name)
	case *ast.TypeSpec:
		return fmt.Sprintf("type %v", x.Name)
	case *ast.InterfaceType:
		return fmt.Sprintf("interface %v", x.Interface)
	case *ast.Ident:
		if x.Obj != nil {
			fmt.Printf("decl %#v\n", x.Obj.Decl)
		}
		return fmt.Sprintf("indent %v(%v)", x.Name, x.Obj)
	case *ast.BasicLit:
		return fmt.Sprintf("basic %s %s", x.Kind, x.Value)
	case *ast.CompositeLit:
		return fmt.Sprintf("composite %v", x.Type)
	case *ast.ParenExpr:
		return fmt.Sprintf("parent %v", x.X)
	case *ast.FuncLit:
		return fmt.Sprintf("func %v", x.Type)
	case *ast.SelectorExpr:
		return "selector " + NodeStr(x.X) + "." + NodeStr(x.Sel)
	case *ast.IndexExpr:
		return fmt.Sprintf("indent %v", x.Index)
	case *ast.SliceExpr:
		return fmt.Sprintf("slice %v", x.X)
	case *ast.TypeAssertExpr:
		return fmt.Sprintf("type assert %v", x.Type)
	case *ast.CallExpr:
		var argsStr []string
		for _, n := range x.Args {
			argsStr = append(argsStr, NodeStr(n))
		}
		return fmt.Sprintf("call %v(%s)", NodeStr(x.Fun), strings.Join(argsStr, ","))
	case *ast.StarExpr:
		return fmt.Sprintf("star %v", x.X)
	case *ast.UnaryExpr:

		return fmt.Sprintf("unary %v", NodeStr(x.X))
	case *ast.BinaryExpr:
		return fmt.Sprintf("binary %v", x.X)
	case *ast.KeyValueExpr:
		return fmt.Sprintf("key val %v %v", x.Key, x.Value)
	case *ast.FuncDecl:
		return fmt.Sprintf("func decl %v", x.Name)
	case *ast.MapType:
		return fmt.Sprintf("map %v", x.Key)
	case *ast.FieldList:
		return fmt.Sprintf("field list %v", x.NumFields())
	case *ast.Field:
		return fmt.Sprintf("field %v", x.Type)
	case *ast.StructType:
		return fmt.Sprintf("struct %v", x.Fields)
	}
	fmt.Printf("unknow %#v\n", node)
	return ""
}

func TestBuildPackage(t *testing.T) {
	directory := "data/"
	pkg, err := build.Default.ImportDir(directory, 0)
	if err != nil {
		panic(err)
	}
	names := pkg.GoFiles
	for i := range names {
		names[i] = directory + names[i]
	}

	fs := token.NewFileSet()
	var astFiles []*ast.File
	for _, name := range names {
		astFile, err := parser.ParseFile(fs, name, nil, 0)
		if err != nil {
			panic(err)
		}
		ast.Inspect(astFile, func(node ast.Node) bool {
			if node != nil {
				fmt.Printf("%s\n|- %s %s\n", fs.Position(node.Pos()), fs.Position(node.End()), NodeStr(node))
			}
			return true
		})
		astFiles = append(astFiles, astFile)
	}

}
