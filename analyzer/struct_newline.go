package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// StructNewline checks for blank lines after embedded fields and before private fields.
var StructNewline = &analysis.Analyzer{
	Name: "structnewline",
	Doc:  "checks for a blank line after embedded fields and before the first private field",
	Run:  runStructNewline,
}

func runStructNewline(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		ast.Inspect(file, func(node ast.Node) bool {
			structType, isStructType := node.(*ast.StructType)
			if !isStructType || structType.Fields == nil {
				return true
			}

			fields := structType.Fields.List
			for index := 1; index < len(fields); index++ {
				previousField := fields[index-1]
				currentField := fields[index]

				if hasBlankLineBetween(pass, previousField.End(), currentField.Pos()) {
					continue
				}

				if isEmbeddedField(previousField) {
					pass.Reportf(currentField.Pos(), "missing blank line after embedded field")

					continue
				}

				if isExportedField(previousField) && isPrivateField(currentField) {
					pass.Reportf(currentField.Pos(), "missing blank line before private fields")
				}
			}

			return true
		})
	})

	return nil, nil
}

func isEmbeddedField(field *ast.Field) bool {
	return len(field.Names) == 0
}

func isExportedField(field *ast.Field) bool {
	if len(field.Names) == 0 {
		return false
	}

	return ast.IsExported(field.Names[0].Name)
}

func isPrivateField(field *ast.Field) bool {
	if len(field.Names) == 0 {
		return false
	}

	return !ast.IsExported(field.Names[0].Name)
}
