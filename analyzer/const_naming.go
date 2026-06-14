package analyzer

import (
	"go/ast"
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

// ConstNaming checks that constants use ALL_CAPS_WITH_UNDERSCORES naming.
var ConstNaming = &analysis.Analyzer{
	Name: "constnaming",
	Doc:  "checks that constants use ALL_CAPS_WITH_UNDERSCORES naming",
	Run:  runConstNaming,
}

func runConstNaming(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		for _, declaration := range file.Decls {
			generalDeclaration, isGeneralDeclaration := declaration.(*ast.GenDecl)
			if !isGeneralDeclaration || generalDeclaration.Tok != token.CONST {
				continue
			}

			for _, spec := range generalDeclaration.Specs {
				valueSpec, isValueSpec := spec.(*ast.ValueSpec)
				if !isValueSpec {
					continue
				}

				for _, name := range valueSpec.Names {
					if name.Name == "_" {
						continue
					}

					if !isAllCaps(name.Name) {
						pass.Reportf(name.Pos(), "constant '%s' should be ALL_CAPS_WITH_UNDERSCORES", name.Name)
					}
				}
			}
		}
	})

	return nil, nil
}

func isAllCaps(name string) bool {
	if len(name) == 0 || strings.HasPrefix(name, "_") {
		return false
	}

	for _, character := range name {
		if character == '_' || unicode.IsUpper(character) || unicode.IsDigit(character) {
			continue
		}

		return false
	}

	return true
}
