package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

//!+input
const input = `package main
import "net/http"
type A struct{}
func (*A) f(rawReq *http.Request)
type B int
func (B) f(rawReq *http.Request)
func (*B) g(rawReq *http.Request)

`

const ifaceInput = `package main
import "net/http"

type I interface { f(rawReq *http.Request) }
type J interface { g(rawReq *http.Request) }
`

//!-input

func DiffPackage() {
	// Parse one file.
	fset := token.NewFileSet()
	var sourcePkg *types.Package
	var apiPkg *types.Package
	{
		f, err := parser.ParseFile(fset, "input.go", input, 0)
		if err != nil {
			log.Fatal(err) // parse error
		}
		conf := types.Config{Importer: importer.Default()}
		sourcePkg, err = conf.Check("hello", fset, []*ast.File{f}, nil)
		if err != nil {
			log.Fatal(err) // type error
		}
	}
	{
		f, err := parser.ParseFile(fset, "iface.go", ifaceInput, 0)
		if err != nil {
			log.Fatal(err) // parse error
		}
		conf := types.Config{Importer: importer.Default()}
		apiPkg, err = conf.Check("world", fset, []*ast.File{f}, nil)
		if err != nil {
			log.Fatal(err) // type error
		}
	}

	//!+implements
	// Find all named types at package level.
	var allNamed []*types.Named
	for _, name := range sourcePkg.Scope().Names() {
		if obj, ok := sourcePkg.Scope().Lookup(name).(*types.TypeName); ok {
			allNamed = append(allNamed, obj.Type().(*types.Named))
		}
	}
	fmt.Println(apiPkg.Scope().Names())
	iIface := apiPkg.Scope().Lookup("I").Type().Underlying().(*types.Interface)
	jIface := apiPkg.Scope().Lookup("J").Type().Underlying().(*types.Interface)
	fmt.Println(iIface)
	fmt.Println(jIface)
	for _, T := range allNamed {
		if types.AssignableTo(T, iIface) {
			fmt.Printf("%s satisfies %s\n", T, iIface)
		} else if types.AssignableTo(types.NewPointer(T), iIface) {
			fmt.Printf("%s satisfies %s\n", types.NewPointer(T), iIface)
		}
		if types.AssignableTo(T, jIface) {
			fmt.Printf("%s satisfies %s\n", T, jIface)
		} else if types.AssignableTo(types.NewPointer(T), jIface) {
			fmt.Printf("%s satisfies %s\n", types.NewPointer(T), jIface)
		}
	}

}

func SamePackage() {
	// Parse one file.
	var err error
	fset := token.NewFileSet()
	var inputFile, ifaceFile *ast.File
	{
		inputFile, err = parser.ParseFile(fset, "input.go", input, 0)
		if err != nil {
			log.Fatal(err) // parse error
		}
	}
	{
		ifaceFile, err = parser.ParseFile(fset, "iface.go", ifaceInput, 0)
		if err != nil {
			log.Fatal(err) // parse error
		}

	}
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("hello", fset, []*ast.File{inputFile, ifaceFile}, nil)
	if err != nil {
		log.Fatal(err) // type error
	}

	//!+implements
	// Find all named types at package level.
	var allNamed []*types.Named
	for _, name := range pkg.Scope().Names() {
		if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok {
			allNamed = append(allNamed, obj.Type().(*types.Named))
		}
	}
	fmt.Println(pkg.Scope().Names())
	iIface := pkg.Scope().Lookup("I").Type().Underlying().(*types.Interface)
	jIface := pkg.Scope().Lookup("J").Type().Underlying().(*types.Interface)
	fmt.Println(iIface)
	fmt.Println(jIface)
	for _, T := range allNamed {
		if types.AssignableTo(T, iIface) {
			fmt.Printf("%s satisfies %s\n", T, iIface)
		} else if types.AssignableTo(types.NewPointer(T), iIface) {
			fmt.Printf("%s satisfies %s\n", types.NewPointer(T), iIface)
		}
		if types.AssignableTo(T, jIface) {
			fmt.Printf("%s satisfies %s\n", T, jIface)
		} else if types.AssignableTo(types.NewPointer(T), jIface) {
			fmt.Printf("%s satisfies %s\n", types.NewPointer(T), jIface)
		}
	}
	// Test assignability of all distinct pairs of
	// named types (T, U) where U is an interface.

	//!-implements
}

func main() {
	SamePackage()
}
