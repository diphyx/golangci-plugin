package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// ExportedDoc checks that exported functions and methods have a documentation comment.
var ExportedDoc = &analysis.Analyzer{
	Name: "exporteddoc",
	Doc:  "checks that exported functions and method receivers have a documentation comment",
	Run:  runExportedDoc,
}

func runExportedDoc(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		for _, declaration := range file.Decls {
			functionDeclaration, isFunctionDeclaration := declaration.(*ast.FuncDecl)
			if !isFunctionDeclaration {
				continue
			}

			functionName := functionDeclaration.Name.Name
			if !ast.IsExported(functionName) {
				continue
			}

			if functionDeclaration.Doc != nil {
				continue
			}

			if functionDeclaration.Recv != nil {
				pass.Reportf(functionDeclaration.Name.Pos(),
					"exported method '%s' should have a documentation comment",
					functionName)

				continue
			}

			pass.Reportf(functionDeclaration.Name.Pos(),
				"exported function '%s' should have a documentation comment",
				functionName)
		}
	})

	return nil, nil
}
