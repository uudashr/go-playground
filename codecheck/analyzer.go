package codecheck

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "codecheck",
	Doc:      "Code check",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

var (
	limit int // -limit flag
)

func init() {
	Analyzer.Flags.IntVar(&limit, "limit", 2, "Limit on struct literals without key")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		lit := n.(*ast.CompositeLit)

		ident, ok := lit.Type.(*ast.Ident)
		if !ok {
			return
		}

		if ident.Obj.Kind != ast.Typ {
			return
		}

		fmt.Printf("%#+v\n", n)
		fmt.Printf("Lit type: %#+v\n", lit.Type)
		fmt.Printf("Ident: %#+v\n", ident)
		fmt.Printf("Ident Obj Decl: %#+v\n", ident.Obj.Decl)
		typeSpec := ident.Obj.Decl.(*ast.TypeSpec)
		structType := typeSpec.Type.(*ast.StructType)
		fmt.Printf("Ident Struct Type: %#+v\n", structType)
		fmt.Printf("Struct field list: %#+v\n", structType.Fields.List)
		fields := structType.Fields.List
		for i, f := range fields {
			fmt.Printf("field[%d]: %#+v name: %#+v\n", i, f, f.Names[0])
		}

		var nonKeyCount int
		for _, e := range lit.Elts {
			if _, ok := e.(*ast.KeyValueExpr); !ok {
				nonKeyCount++
			}

			fmt.Printf("- Element %#+v\n", e)
		}

		// fmt.Println("Non Keys:", nonKeyCount)
		if nonKeyCount > limit {
			pass.Reportf(lit.Pos(), "found %d non keyed on struct literal (> %d)", nonKeyCount, limit)
		}
	})
	return nil, nil
}
