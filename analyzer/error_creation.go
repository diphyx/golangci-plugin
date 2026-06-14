package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// ErrorCreation checks that errors.New is used for static messages and fmt.Errorf for formatted ones.
var ErrorCreation = &analysis.Analyzer{
	Name: "errorcreation",
	Doc:  "checks that errors.New is used for static messages and fmt.Errorf for formatted messages",
	Run:  runErrorCreation,
}

func runErrorCreation(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		ast.Inspect(file, func(node ast.Node) bool {
			callExpression, isCallExpression := node.(*ast.CallExpr)
			if !isCallExpression {
				return true
			}

			packageName, functionName, isPackageCall := packageCall(callExpression.Fun)
			if !isPackageCall {
				return true
			}

			if packageName == "fmt" && functionName == "Errorf" && len(callExpression.Args) == 1 {
				pass.Reportf(callExpression.Pos(), "use errors.New for static error messages")

				return true
			}

			if packageName == "errors" && functionName == "New" && hasFormatVerb(callExpression.Args) {
				pass.Reportf(callExpression.Pos(), "use fmt.Errorf for formatted error messages")
			}

			return true
		})
	})

	return nil, nil
}

func packageCall(expression ast.Expr) (string, string, bool) {
	selectorExpression, isSelectorExpression := expression.(*ast.SelectorExpr)
	if !isSelectorExpression {
		return "", "", false
	}

	identifier, isIdentifier := selectorExpression.X.(*ast.Ident)
	if !isIdentifier {
		return "", "", false
	}

	return identifier.Name, selectorExpression.Sel.Name, true
}

func hasFormatVerb(arguments []ast.Expr) bool {
	if len(arguments) != 1 {
		return false
	}

	basicLiteral, isBasicLiteral := arguments[0].(*ast.BasicLit)
	if !isBasicLiteral || basicLiteral.Kind != token.STRING {
		return false
	}

	return strings.Contains(basicLiteral.Value, "%")
}
