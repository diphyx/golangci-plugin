package analyzer_test

import (
	"testing"

	"github.com/diphyx/golangci-plugin/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestBoolNaming(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzer.BoolNaming, "boolnaming")
}

func TestConstNaming(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzer.ConstNaming, "constnaming")
}

func TestConstructorNaming(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzer.ConstructorNaming, "constructornaming")
}

func TestErrorCreation(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzer.ErrorCreation, "errorcreation")
}
