package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// BoolNaming checks that boolean variables have a proper prefix.
var BoolNaming = &analysis.Analyzer{
	Name: "boolnaming",
	Doc:  "checks that boolean variables have a prefix: is, has, can, should, enable, or in",
	Run:  runBoolNaming,
}

var booleanPrefixes = []string{"is", "has", "can", "should", "enable", "in"}

var booleanExceptions = map[string]bool{
	"ok":    true,
	"found": true,
	"done":  true,
	"valid": true,
}

func runBoolNaming(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		ast.Inspect(file, func(node ast.Node) bool {
			assignStatement, isAssignStatement := node.(*ast.AssignStmt)
			if !isAssignStatement || assignStatement.Tok != token.DEFINE {
				return true
			}

			if len(assignStatement.Lhs) != 1 || len(assignStatement.Rhs) != 1 {
				return true
			}

			identifier, isIdentifier := assignStatement.Lhs[0].(*ast.Ident)
			if !isIdentifier || identifier.Name == "_" {
				return true
			}

			if !isBooleanLiteral(assignStatement.Rhs[0]) {
				return true
			}

			if booleanExceptions[identifier.Name] {
				return true
			}

			if hasBooleanPrefix(identifier.Name) {
				return true
			}

			pass.Reportf(identifier.Pos(),
				"boolean variable '%s' should have a prefix: is, has, can, should, enable, or in",
				identifier.Name)

			return true
		})
	})

	return nil, nil
}

func isBooleanLiteral(expression ast.Expr) bool {
	identifier, isIdentifier := expression.(*ast.Ident)
	if !isIdentifier {
		return false
	}

	return identifier.Name == "true" || identifier.Name == "false"
}

func hasBooleanPrefix(name string) bool {
	for _, prefix := range booleanPrefixes {
		if strings.HasPrefix(name, prefix) && len(name) > len(prefix) {
			nextCharacter := rune(name[len(prefix)])
			if nextCharacter >= 'A' && nextCharacter <= 'Z' {
				return true
			}
		}
	}

	return false
}
