package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// ErrorNaming checks that error variables use '{functionName}Error' pattern.
var ErrorNaming = &analysis.Analyzer{
	Name: "errornaming",
	Doc:  "checks that error variables use '{functionName}Error' pattern instead of 'err' or 'e'",
	Run:  runErrorNaming,
}

func runErrorNaming(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		ifInitAssigns := collectIfInitAssigns(file)

		ast.Inspect(file, func(node ast.Node) bool {
			assignStatement, isAssignStatement := node.(*ast.AssignStmt)
			if !isAssignStatement || assignStatement.Tok != token.DEFINE {
				return true
			}

			if ifInitAssigns[assignStatement] {
				return true
			}

			for _, leftHandSide := range assignStatement.Lhs {
				identifier, isIdentifier := leftHandSide.(*ast.Ident)
				if !isIdentifier {
					continue
				}

				if identifier.Name == "err" || identifier.Name == "e" {
					pass.Reportf(identifier.Pos(), "error variable '%s' should use '{functionName}Error' pattern", identifier.Name)
				} else if strings.HasSuffix(identifier.Name, "Err") && !strings.HasSuffix(identifier.Name, "Error") {
					pass.Reportf(identifier.Pos(), "error variable '%s' should end with 'Error' not 'Err'", identifier.Name)
				}
			}

			return true
		})
	})

	return nil, nil
}

func collectIfInitAssigns(file *ast.File) map[*ast.AssignStmt]bool {
	result := make(map[*ast.AssignStmt]bool)

	ast.Inspect(file, func(node ast.Node) bool {
		ifStatement, isIfStatement := node.(*ast.IfStmt)
		if !isIfStatement {
			return true
		}

		if ifStatement.Init == nil {
			return true
		}

		assignStatement, isAssignStatement := ifStatement.Init.(*ast.AssignStmt)
		if isAssignStatement {
			result[assignStatement] = true
		}

		return true
	})

	return result
}
