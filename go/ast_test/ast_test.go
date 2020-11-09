/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package ast_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"go/ast"
	"go/build"
	"go/parser"
	"go/printer"
	"go/token"
	"os"

	//"sort"
	"go/importer"
	"go/types"
	"strings"
	"testing"
)

var directory = "data/"

func TestAstPrint(t *testing.T) {

	// Create the AST by parsing src.
	fs := token.NewFileSet() // positions are relative to fset

	fMap, err := parser.ParseDir(fs, directory, nil, 0|parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	for _, f := range fMap {
		err = ast.Print(fs, f)
		if err != nil {
			panic(err)
		}
	}
}

type visitor struct{ fs *token.FileSet }

func (s visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.BasicLit:
		n.Value = n.Value + "xxx"
	}
	return s
}
func TestAstMergeLinePrint(t *testing.T) {

	// Create the AST by parsing src.
	fs := token.NewFileSet() // positions are relative to fset

	fMap, err := parser.ParseDir(fs, directory, nil, 0|parser.ParseComments)
	if err != nil {
		panic(err)
	}
	cfg := printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}

	// Print the AST.
	for _, f := range fMap {

		ast.Walk(visitor{fs: fs}, f)
		for _, f := range f.Files {
			err = cfg.Fprint(os.Stdout, fs, f)
			if err != nil {
				panic(err)
			}
		}
	}
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

// test prefix print
// test print
func (log *myLogger) Print() { //xxx
// xxx2
	// why
	fmt.Printf("test")
}

func main() {
	// test stmt print
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

	//for n, _ := range cmap {
	//	ast.Print(fset, n)
	//}
	ast.Inspect(f, func(node ast.Node) bool {
		if comments, ok := cmap[node]; ok {
			ast.Print(fset, comments)
		}
		return true
	})

}

func TestAstStruct(t *testing.T) {

	fs := token.NewFileSet()
	astFile, err := parser.ParseFile(fs, directory+"join_example.go", nil, 0|parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, decl := range astFile.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok {
			if gen.Tok == token.TYPE {
				fmt.Printf("%s\n|- %s %s\n", fs.Position(decl.Pos()), fs.Position(decl.End()), NodeStr(gen))
				for _, spec := range gen.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						fmt.Printf("%s\n|- %s %s\n", fs.Position(decl.Pos()), fs.Position(decl.End()), NodeStr(typeSpec))
						if structType, ok := typeSpec.Type.(*ast.StructType); ok {
							fmt.Printf("%s\n|- %s struct  %s\n", fs.Position(decl.Pos()), fs.Position(decl.End()), NodeStr(structType))
							for _, field := range structType.Fields.List {
								fmt.Printf("field %#v %v %s\n", field.Type, field.Tag, field.Names)
							}
						}
					}
				}
			}
		} else {
			fmt.Printf("%s\n|- %s %s\n", fs.Position(decl.Pos()), fs.Position(decl.End()), NodeStr(decl))
		}
	}

}

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
		return fmt.Sprintf("selector (%s).(%s)", NodeStr(x.Sel), NodeStr(x.X))
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

func TestFindAllModel(t *testing.T) {
	fs := token.NewFileSet()
	astFile, err := parser.ParseFile(fs, directory+"join_example.go", nil, 0|parser.ParseComments)
	if err != nil {
		panic(err)
	}
	modelMap := map[*ast.CallExpr]ast.Expr{}
	ast.Inspect(astFile, func(node ast.Node) bool {
		if callExpr, ok := node.(*ast.CallExpr); ok {
			if selector, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				if selector.Sel.Name == "Model" {
					if ident, ok := selector.X.(*ast.Ident); ok && ident.Name == "toy" {
						t.Logf("%s\n|- %s %s\n", fs.Position(node.Pos()), fs.Position(node.End()), NodeStr(node))
						arg := callExpr.Args[0]
					ForExpr:
						for true {
							switch x := arg.(type) {
							case *ast.StarExpr:
								arg = x.X
							case *ast.UnaryExpr:
								arg = x.X
							default:
								break ForExpr
							}
						}
						modelMap[callExpr] = arg
						t.Logf("%s\n|- %s %s\n", fs.Position(arg.Pos()), fs.Position(arg.End()), NodeStr(arg))
					}
				} else if selector.Sel.Name == "Join" {
					t.Logf("%s\n|- %s join\n", fs.Position(node.Pos()), fs.Position(node.End()))
				}
			}
		}
		return true
	})
}

func TestIdentObj(t *testing.T) {
	fs := token.NewFileSet() // positions are relative to fset
	astFile, err := parser.ParseDir(fs, directory, nil, 0|parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, file := range astFile {

		ast.Inspect(file, func(node ast.Node) bool {
			if x, ok := node.(*ast.Ident); ok {
				if x.Name != "toy" {
					return true
				}
				fmt.Printf("%s\n|- %s ident\n", fs.Position(node.Pos()), fs.Position(node.End()))
				if x.Obj != nil {
					fmt.Printf("type %v data %v decl %#v \n", x.Obj.Type, x.Obj.Data, x.Obj.Decl)
					if valueSpec, ok := x.Obj.Decl.(*ast.ValueSpec); ok {
						ast.Print(fs, valueSpec.Type)
					} else if assignStmt, ok := x.Obj.Decl.(*ast.AssignStmt); ok {
						if id, ok := assignStmt.Lhs[0].(*ast.Ident); ok {
							if valueSpec, ok := id.Obj.Decl.(*ast.ValueSpec); ok {
								ast.Print(fs, valueSpec.Type)
							}
						}
						ast.Print(fs, assignStmt)
					}
				}
			}
			return true
		})
	}
}

type Context struct {
	Stack []*ast.StructType
}

func TestTypesCheck(t *testing.T) {
	fs := token.NewFileSet()
	parserFile, err := parser.ParseFile(fs, "types_check/main.go", nil, 0|parser.ParseComments)
	if err != nil {
		panic(err)
	}
	config := types.Config{Importer: importer.For("source", nil), FakeImportC: true}
	info := &types.Info{
		Uses:       map[*ast.Ident]types.Object{},
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Selections: map[*ast.SelectorExpr]*types.Selection{},
		Scopes:     map[ast.Node]*types.Scope{},
		Defs:       make(map[*ast.Ident]types.Object),
	}
	_, err = config.Check("types_check/", fs, []*ast.File{parserFile}, info)
	if err != nil {
		panic(err)
	}
	//ast.Print(fs, parserFile)
	//ast.Inspect(parserFile, func(node ast.Node) bool {
	//	switch x := node.(type) {
	//	case *ast.Ident:
	//		if x.Name == "modelPreload" {
	//			fmt.Printf("pos %s\n", fs.Position(x.Pos()))
	//			if x.Obj != nil {
	//				ast.Print(fs, x.Obj.Decl)
	//			}
	//		}
	//	}
	//	return true
	//})
	//fmt.Printf("types\n")
	//for k, v := range info.Types {
	//	fmt.Printf("%s \n |- %s %#v : %#v \n%#v\n", fs.Position(k.Pos()), fs.Position(k.End()), k, v, v.Type.String())
	//	if p, ok := v.Type.(*types.Pointer); ok {
	//		fmt.Printf("elem %v\n", p.Elem())
	//	}
	//	fmt.Println("value", v.Value)
	//	if sType, ok := v.Type.(*types.Struct); ok {
	//		for i := 0; i < sType.NumFields(); i++ {
	//			fmt.Printf("field: %v\n", sType.Field(i))
	//		}
	//	}
	//	if sType, ok := v.Type.(*types.Named); ok {
	//		for i := 0; i < sType.NumMethods(); i++ {
	//			fmt.Printf("method: %v\n", sType.Method(i))
	//		}
	//	}
	//}
	//
	fmt.Printf("defs\n")
	for k, v := range info.Defs {
		fmt.Printf("%s \n |- %s %v : %#v\n", fs.Position(k.Pos()), fs.Position(k.End()), k, v)

		if v != nil {
			if sType, ok := v.Type().(*types.Struct); ok {
				for i := 0; i < sType.NumFields(); i++ {
					fmt.Printf("field: %v\n", sType.Field(i))
				}
			} else if sType, ok := v.Type().(*types.Signature); ok {
				fmt.Printf("signature: %v\n", sType)
				fmt.Println(sType)
			}

		}
	}
	//
	//fmt.Printf("scopes\n")
	//for k, v := range info.Scopes {
	//	fmt.Printf("%s \n |- %s %v : %v\n", fs.Position(k.Pos()), fs.Position(k.End()), k, v)
	//}
	//fmt.Printf("selectors\n")
	//for k, v := range info.Selections {
	//	fmt.Printf("%s \n |- %s %v : %v\n", fs.Position(k.Pos()), fs.Position(k.End()), k, v)
	//	fmt.Printf("recrive %v\n", v.Type())
	//	fmt.Printf("obj pos %s %s\n", fs.Position(v.Obj().Pos()), v.Obj())
	//}
	//fmt.Printf("users\n")
	//for k, v := range info.Uses {
	//	fmt.Printf("%s \n |- %s %v : %s \n %v \n %#v\n", fs.Position(k.Pos()), fs.Position(k.End()),
	//		k, v, v, v)
	//	fmt.Printf("point %p\n", v)
	//	fmt.Printf("obj %v\n", v.String())
	//}
}

func TestStructPrint(t *testing.T) {
	src := `
package main
type GlobalStruct struct {
	Example string
	Embed struct {
		EmbedField string
	}
}

func main() {
    data :=  struct {
		VarField string
	} {"123"}
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	require.NoError(t, err)
	//ast.Inspect(f, func(node ast.Node) bool {
	//
	//})

	err = ast.Print(fset, f)
	require.NoError(t, err)
}
