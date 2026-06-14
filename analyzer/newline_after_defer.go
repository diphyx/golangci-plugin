package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// NewlineAfterDefer checks for blank line after defer statements.
var NewlineAfterDefer = &analysis.Analyzer{
	Name: "newlineafterdefer",
	Doc:  "checks for blank line after defer statements",
	Run:  runNewlineAfterDefer,
}

func runNewlineAfterDefer(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		ast.Inspect(file, func(node ast.Node) bool {
			blockStatement, isBlockStatement := node.(*ast.BlockStmt)
			if !isBlockStatement {
				return true
			}

			for index, statement := range blockStatement.List {
				deferStatement, isDeferStatement := statement.(*ast.DeferStmt)
				if !isDeferStatement {
					continue
				}

				if isDeferClosure(deferStatement) {
					continue
				}

				if index >= len(blockStatement.List)-1 {
					continue
				}

				nextStatement := blockStatement.List[index+1]
				if hasBlankLineBetween(pass, deferStatement.End(), nextStatement.Pos()) {
					continue
				}

				_, isNextDefer := nextStatement.(*ast.DeferStmt)
				if isNextDefer {
					continue
				}

				pass.Reportf(nextStatement.Pos(), "missing blank line after defer statement")
			}

			return true
		})
	})

	return nil, nil
}

func isDeferClosure(deferStatement *ast.DeferStmt) bool {
	_, isFunctionLiteral := deferStatement.Call.Fun.(*ast.FuncLit)

	return isFunctionLiteral
}
