package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// ConstructorNaming checks that constructor functions follow the New{TypeName} pattern.
var ConstructorNaming = &analysis.Analyzer{
	Name: "constructornaming",
	Doc:  "checks that exported functions returning *TypeName are named New{TypeName}",
	Run:  runConstructorNaming,
}

func runConstructorNaming(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		for _, declaration := range file.Decls {
			functionDeclaration, isFunctionDeclaration := declaration.(*ast.FuncDecl)
			if !isFunctionDeclaration {
				continue
			}

			if functionDeclaration.Recv != nil {
				continue
			}

			if functionDeclaration.Type.Results == nil {
				continue
			}

			functionName := functionDeclaration.Name.Name
			if !ast.IsExported(functionName) {
				continue
			}

			returnTypeName := extractReturnPointerType(functionDeclaration.Type.Results)
			if returnTypeName == "" {
				continue
			}

			if strings.HasPrefix(functionName, "New") {
				continue
			}

			if hasActionPrefix(functionName) {
				continue
			}

			expectedName := "New" + returnTypeName
			pass.Reportf(functionDeclaration.Name.Pos(),
				"function '%s' returns *%s, should be named '%s'",
				functionName, returnTypeName, expectedName)
		}
	})

	return nil, nil
}

var actionPrefixes = []string{
	"Parse", "Collect", "Read", "Load", "Ensure", "Get", "Find",
	"Create", "Build", "Make", "Open", "Connect", "Start", "Init",
	"Fetch", "Extract", "Resolve", "Decode", "Unmarshal", "From",
	"Pointer", "Clone", "Copy", "Wrap",
}

func hasActionPrefix(name string) bool {
	for _, prefix := range actionPrefixes {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}

	return false
}

func extractReturnPointerType(results *ast.FieldList) string {
	if len(results.List) == 0 {
		return ""
	}

	firstResult := results.List[0]
	starExpression, isStarExpression := firstResult.Type.(*ast.StarExpr)
	if !isStarExpression {
		return ""
	}

	identifier, isIdentifier := starExpression.X.(*ast.Ident)
	if !isIdentifier {
		return ""
	}

	if !ast.IsExported(identifier.Name) {
		return ""
	}

	return identifier.Name
}
