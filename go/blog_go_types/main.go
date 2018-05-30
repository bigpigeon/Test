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
	f, err := parser.ParseFile(fset, "data/src.go", nil, 0)
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
	// 这里第一个path参数觉得当前pkg前缀，和FileSet的文件路径是无关的
	pkg, err := config.Check("", fset, []*ast.File{f}, info)

	if err != nil {
		panic(err)
	}
	// 打印types
	fmt.Println("------------ types -----------")
	for _, node := range getSortedKeys(info.Types) {
		expr := node.(ast.Expr)
		typeValue := info.Types[expr]
		fmt.Printf("%s - %s %T it's value: %v type: %s\n",
			fset.Position(expr.Pos()),
			fset.Position(expr.End()),
			expr,
			typeValue.Value,
			typeValue.Type,
		)
		if typeValue.Assignable() {
			fmt.Print("assignable ")
		}
		if typeValue.Addressable() {
			fmt.Print("addressable ")
		}
		if typeValue.IsNil() {
			fmt.Print("nil ")
		}
		if typeValue.HasOk() {
			fmt.Print("has ok ")
		}
		if typeValue.IsBuiltin() {
			fmt.Print("builtin ")
		}
		if typeValue.IsType() {
			fmt.Print("is type ")
		}
		if typeValue.IsValue() {
			fmt.Print("is value ")
		}
		if typeValue.IsVoid() {
			fmt.Print("void ")
		}
		fmt.Println()
	}
	// 打印defs
	fmt.Println("------------ def -----------")
	for _, node := range getSortedKeys(info.Defs) {
		ident := node.(*ast.Ident)
		object := info.Defs[ident]
		fmt.Printf("%s - %s %T",
			fset.Position(ident.Pos()),
			fset.Position(ident.End()),
			object,
		)
		if object != nil {
			fmt.Printf(" it's object: %s type: %s pos: %s",
				object,
				object.Type().String(),
				fset.Position(object.Pos()),
			)

		}
		fmt.Println()
	}
	// 打印Uses
	fmt.Println("------------ uses -----------")
	for _, node := range getSortedKeys(info.Uses) {
		ident := node.(*ast.Ident)
		object := info.Uses[ident]
		fmt.Printf("%s - %s %T",
			fset.Position(ident.Pos()),
			fset.Position(ident.End()),
			object,
		)
		if object != nil {
			fmt.Printf(" it's object: %s type: %s pos: %s",
				object,
				object.Type().String(),
				fset.Position(object.Pos()),
			)

		}
		fmt.Println()
	}
	// 打印Implicits
	fmt.Println("------------ implicits -----------")
	for _, node := range getSortedKeys(info.Implicits) {
		object := info.Implicits[node]
		fmt.Printf("%s - %s %T it's object: %s\n",
			fset.Position(node.Pos()),
			fset.Position(node.End()),
			node,
			object,
		)
	}
	// 打印Selections
	fmt.Println("------------ selections -----------")
	for _, node := range getSortedKeys(info.Selections) {
		sel := node.(*ast.SelectorExpr)
		typeSel := info.Selections[sel]
		fmt.Printf("%s - %s it's selection: %s\n",
			fset.Position(sel.Pos()),
			fset.Position(sel.End()),
			typeSel.String(),
		)
		fmt.Printf("receive: %s index: %v obj: %s\n", typeSel.Recv(), typeSel.Index(), typeSel.Obj())
	}
	// 打印Scopes
	fmt.Println("------------ scopes -----------")
	//打印package scope
	fmt.Printf("package level scope %s\n",
		pkg.Scope().String(),
	)
	// 打印宇宙级scope
	fmt.Printf("universe level scope %s\n",
		pkg.Scope().Parent().String(),
	)
	for _, node := range getSortedKeys(info.Scopes) {
		scope := info.Scopes[node]
		fmt.Printf("%s - %s %T it's scope: %s\n",
			fset.Position(node.Pos()),
			fset.Position(node.End()),
			node,
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
