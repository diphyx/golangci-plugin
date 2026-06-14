package analyzer

import "golang.org/x/tools/go/analysis"

// All returns every diphyx code style analyzer. It is the single source of
// truth shared by the golangci-lint plugin and the standalone command.
func All() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		BoolNaming,
		ConstNaming,
		ConstructorNaming,
		ErrorCreation,
		ErrorNaming,
	}
}
