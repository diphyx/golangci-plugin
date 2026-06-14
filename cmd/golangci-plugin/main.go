// Command golangci-plugin runs the diphyx code style analyzers as a
// standalone vet-style tool, independent of golangci-lint.
package main

import (
	"github.com/diphyx/golangci-plugin/analyzer"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(analyzer.All()...)
}
