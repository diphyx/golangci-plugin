package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// forEachFile runs visit for every non-test file in the pass, skipping
// _test.go files so that fixtures and tests are never flagged.
func forEachFile(pass *analysis.Pass, visit func(file *ast.File)) {
	for _, file := range pass.Files {
		if isTestFile(pass.Fset, file) {
			continue
		}

		visit(file)
	}
}

// hasBlankLineBetween reports whether at least one blank line separates the
// end position from the start position.
func hasBlankLineBetween(pass *analysis.Pass, endPos, startPos token.Pos) bool {
	endLine := pass.Fset.Position(endPos).Line
	startLine := pass.Fset.Position(startPos).Line

	return startLine-endLine >= 2
}

func isTestFile(fileSet *token.FileSet, file *ast.File) bool {
	fileName := fileSet.Position(file.Pos()).Filename

	return strings.HasSuffix(fileName, "_test.go")
}
