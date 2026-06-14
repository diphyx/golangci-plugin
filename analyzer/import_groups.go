package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// ImportGroups checks that standard library imports precede third-party imports
// and that blank imports appear last.
var ImportGroups = &analysis.Analyzer{
	Name: "importgroups",
	Doc:  "checks that standard library imports come before third-party imports and blank imports are last",
	Run:  runImportGroups,
}

func runImportGroups(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		var seenThirdParty bool
		var firstBlank *ast.ImportSpec

		for _, importSpec := range file.Imports {
			path := strings.Trim(importSpec.Path.Value, `"`)

			if importSpec.Name != nil && importSpec.Name.Name == "_" {
				if firstBlank == nil {
					firstBlank = importSpec
				}

				continue
			}

			if firstBlank != nil {
				pass.Reportf(firstBlank.Pos(),
					"blank import '%s' should be at the end of the imports",
					strings.Trim(firstBlank.Path.Value, `"`))

				firstBlank = nil
			}

			if isStandardLibraryImport(path) {
				if seenThirdParty {
					pass.Reportf(importSpec.Pos(),
						"standard library import '%s' should come before third-party imports",
						path)
				}

				continue
			}

			seenThirdParty = true
		}
	})

	return nil, nil
}

func isStandardLibraryImport(path string) bool {
	firstSegment := path

	slashIndex := strings.Index(path, "/")
	if slashIndex >= 0 {
		firstSegment = path[:slashIndex]
	}

	return !strings.Contains(firstSegment, ".")
}
