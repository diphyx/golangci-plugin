package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

// NoAbbreviation checks that common abbreviations are not used.
var NoAbbreviation = &analysis.Analyzer{
	Name: "noabbreviation",
	Doc:  "checks that common abbreviations are not used (ctx, cmd, cfg, msg, req, resp, res)",
	Run:  runNoAbbreviation,
}

var forbiddenAbbreviations = map[string]string{
	"ctx":  "context",
	"cmd":  "command",
	"cfg":  "config",
	"msg":  "message",
	"req":  "request",
	"resp": "response",
	"res":  "response",
}

func runNoAbbreviation(pass *analysis.Pass) (any, error) {
	forEachFile(pass, func(file *ast.File) {
		ast.Inspect(file, func(node ast.Node) bool {
			switch typedNode := node.(type) {
			case *ast.AssignStmt:
				if typedNode.Tok != token.DEFINE {
					return true
				}

				for _, leftHandSide := range typedNode.Lhs {
					reportAbbreviation(pass, leftHandSide)
				}

			case *ast.FuncDecl:
				if typedNode.Type.Params != nil {
					for _, field := range typedNode.Type.Params.List {
						for _, name := range field.Names {
							reportAbbreviationName(pass, name)
						}
					}
				}

				if typedNode.Type.Results != nil {
					for _, field := range typedNode.Type.Results.List {
						for _, name := range field.Names {
							reportAbbreviationName(pass, name)
						}
					}
				}

			case *ast.RangeStmt:
				if typedNode.Tok != token.DEFINE {
					return true
				}

				if typedNode.Key != nil {
					reportAbbreviation(pass, typedNode.Key)
				}

				if typedNode.Value != nil {
					reportAbbreviation(pass, typedNode.Value)
				}

			case *ast.ValueSpec:
				for _, name := range typedNode.Names {
					reportAbbreviationName(pass, name)
				}
			}

			return true
		})
	})

	return nil, nil
}

func reportAbbreviation(pass *analysis.Pass, expression ast.Expr) {
	identifier, isIdentifier := expression.(*ast.Ident)
	if !isIdentifier {
		return
	}

	reportAbbreviationName(pass, identifier)
}

func reportAbbreviationName(pass *analysis.Pass, identifier *ast.Ident) {
	if identifier.Name == "_" {
		return
	}

	fullName, isForbidden := forbiddenAbbreviations[identifier.Name]
	if !isForbidden {
		return
	}

	pass.Reportf(identifier.Pos(), "use '%s' instead of '%s'", fullName, identifier.Name)
}
