package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// LoggingLevel checks that only Debug and Error log levels are used.
var LoggingLevel = &analysis.Analyzer{
	Name: "logginglevel",
	Doc:  "checks that only Debug and Error log levels are used (no Info, Warn, Notice, Trace, Fatal, Panic)",
	Run:  runLoggingLevel,
}

var forbiddenLogLevels = map[string]bool{
	"Info":   true,
	"Warn":   true,
	"Notice": true,
	"Trace":  true,
	"Fatal":  true,
	"Panic":  true,
}

func runLoggingLevel(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		ast.Inspect(file, func(node ast.Node) bool {
			selectorExpression, isSelectorExpression := node.(*ast.SelectorExpr)
			if !isSelectorExpression {
				return true
			}

			identifier, isIdentifier := selectorExpression.X.(*ast.Ident)
			if !isIdentifier || identifier.Name != "log" {
				return true
			}

			levelName := selectorExpression.Sel.Name
			if !forbiddenLogLevels[levelName] {
				return true
			}

			pass.Reportf(selectorExpression.Sel.Pos(),
				"log level '%s' is not allowed; use only Debug or Error",
				levelName)

			return true
		})
	})

	return nil, nil
}
