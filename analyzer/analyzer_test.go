package analyzer_test

import (
	"testing"

	"github.com/diphyx/golangci-plugin/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestBoolNaming(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzer.BoolNaming, "boolnaming")
}
