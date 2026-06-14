package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// NewlineBeforeReturn checks for blank line before return statements.
var NewlineBeforeReturn = &analysis.Analyzer{
	Name: "newlinebeforereturn",
	Doc:  "checks for blank line before return statements",
	Run:  runNewlineBeforeReturn,
}

func runNewlineBeforeReturn(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		ast.Inspect(file, func(node ast.Node) bool {
			blockStatement, isBlockStatement := node.(*ast.BlockStmt)
			if !isBlockStatement {
				return true
			}

			for index, statement := range blockStatement.List {
				returnStatement, isReturnStatement := statement.(*ast.ReturnStmt)
				if !isReturnStatement {
					continue
				}

				if index == 0 {
					continue
				}

				previousStatement := blockStatement.List[index-1]
				if hasBlankLineBetween(pass, previousStatement.End(), returnStatement.Pos()) {
					continue
				}

				if isSimpleIfReturn(previousStatement) {
					continue
				}

				if isDefer(previousStatement) {
					continue
				}

				if isCaseOrComm(previousStatement) {
					continue
				}

				pass.Reportf(returnStatement.Pos(), "missing blank line before return statement")
			}

			return true
		})
	})

	return nil, nil
}

func isSimpleIfReturn(statement ast.Stmt) bool {
	ifStatement, isIfStatement := statement.(*ast.IfStmt)
	if !isIfStatement || ifStatement.Body == nil {
		return false
	}

	if len(ifStatement.Body.List) != 1 {
		return false
	}

	_, isReturn := ifStatement.Body.List[0].(*ast.ReturnStmt)

	return isReturn
}

func isDefer(statement ast.Stmt) bool {
	_, isDeferStatement := statement.(*ast.DeferStmt)

	return isDeferStatement
}

func isCaseOrComm(statement ast.Stmt) bool {
	_, isCaseClause := statement.(*ast.CaseClause)
	if isCaseClause {
		return true
	}

	_, isCommClause := statement.(*ast.CommClause)

	return isCommClause
}
