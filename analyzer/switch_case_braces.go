package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

// SwitchCaseBraces checks that each switch case body is wrapped in a block.
var SwitchCaseBraces = &analysis.Analyzer{
	Name: "switchcasebraces",
	Doc:  "checks that each switch case body is wrapped in braces { }",
	Run:  runSwitchCaseBraces,
}

func runSwitchCaseBraces(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		ast.Inspect(file, func(node ast.Node) bool {
			caseClause, isCaseClause := node.(*ast.CaseClause)
			if !isCaseClause {
				return true
			}

			if len(caseClause.Body) == 0 {
				return true
			}

			// A 'fallthrough' must be the last statement of the clause and is
			// illegal inside a nested block, so such a case cannot be wrapped.
			if endsWithFallthrough(caseClause.Body) {
				return true
			}

			if isSingleBlock(caseClause.Body) {
				return true
			}

			pass.Reportf(caseClause.Body[0].Pos(), "switch case body should be wrapped in braces { }")

			return true
		})
	})

	return nil, nil
}

func isSingleBlock(statements []ast.Stmt) bool {
	if len(statements) != 1 {
		return false
	}

	_, isBlockStatement := statements[0].(*ast.BlockStmt)

	return isBlockStatement
}

func endsWithFallthrough(statements []ast.Stmt) bool {
	if len(statements) == 0 {
		return false
	}

	lastStatement := statements[len(statements)-1]
	branchStatement, isBranchStatement := lastStatement.(*ast.BranchStmt)
	if !isBranchStatement {
		return false
	}

	return branchStatement.Tok == token.FALLTHROUGH
}
