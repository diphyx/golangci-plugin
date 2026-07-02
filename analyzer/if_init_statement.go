package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// IfInitStatement checks that an if statement does not carry an init statement,
// so the declaration and the condition never share a single line.
var IfInitStatement = &analysis.Analyzer{
	Name: "ifinitstatement",
	Doc:  "checks that if statements do not use an init statement",
	Run:  runIfInitStatement,
}

func runIfInitStatement(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		ast.Inspect(file, func(node ast.Node) bool {
			ifStatement, isIfStatement := node.(*ast.IfStmt)
			if !isIfStatement {
				return true
			}

			if ifStatement.Init == nil {
				return true
			}

			pass.Reportf(ifStatement.Init.Pos(), "if statement should not use an init statement; declare it before the if")

			return true
		})
	})

	return nil, nil
}
