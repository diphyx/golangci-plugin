package analyzer

import (
	"go/ast"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

// ReceiverNaming checks that method receivers use camelCase of the full type name.
var ReceiverNaming = &analysis.Analyzer{
	Name: "receivernaming",
	Doc:  "checks that method receivers use camelCase of the full type name",
	Run:  runReceiverNaming,
}

func runReceiverNaming(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		for _, declaration := range file.Decls {
			functionDeclaration, isFunctionDeclaration := declaration.(*ast.FuncDecl)
			if !isFunctionDeclaration || functionDeclaration.Recv == nil {
				continue
			}

			if len(functionDeclaration.Recv.List) == 0 {
				continue
			}

			receiverField := functionDeclaration.Recv.List[0]
			if len(receiverField.Names) == 0 {
				continue
			}

			receiverName := receiverField.Names[0].Name
			typeName := extractTypeName(receiverField.Type)
			if typeName == "" {
				continue
			}

			expectedName := toLowerFirst(typeName)
			if receiverName != expectedName && len(receiverName) <= 2 {
				pass.Reportf(receiverField.Names[0].Pos(),
					"receiver '%s' for type '%s' should be '%s' (camelCase of full type name)",
					receiverName, typeName, expectedName)
			}
		}
	})

	return nil, nil
}

func extractTypeName(expression ast.Expr) string {
	switch typedExpression := expression.(type) {
	case *ast.StarExpr:
		return extractTypeName(typedExpression.X)
	case *ast.Ident:
		return typedExpression.Name
	case *ast.IndexExpr:
		return extractTypeName(typedExpression.X)
	case *ast.IndexListExpr:
		return extractTypeName(typedExpression.X)
	}

	return ""
}

func toLowerFirst(name string) string {
	if len(name) == 0 {
		return name
	}

	runes := []rune(name)
	runes[0] = unicode.ToLower(runes[0])

	return string(runes)
}
