package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// FunctionalOptions checks that functions returning a configuration function follow the With{Option} pattern.
var FunctionalOptions = &analysis.Analyzer{
	Name: "functionaloptions",
	Doc:  "checks that exported functions returning func(*T) are named With{Option}",
	Run:  runFunctionalOptions,
}

func runFunctionalOptions(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		for _, declaration := range file.Decls {
			functionDeclaration, isFunctionDeclaration := declaration.(*ast.FuncDecl)
			if !isFunctionDeclaration || functionDeclaration.Recv != nil {
				continue
			}

			functionName := functionDeclaration.Name.Name
			if !ast.IsExported(functionName) {
				continue
			}

			if functionDeclaration.Type.Results == nil || len(functionDeclaration.Type.Results.List) != 1 {
				continue
			}

			_, isFunctionType := functionDeclaration.Type.Results.List[0].Type.(*ast.FuncType)
			if !isFunctionType {
				continue
			}

			if strings.HasPrefix(functionName, "With") {
				continue
			}

			pass.Reportf(functionDeclaration.Name.Pos(),
				"function '%s' returns a configuration function, should be named 'With{Option}'",
				functionName)
		}
	})

	return nil, nil
}
