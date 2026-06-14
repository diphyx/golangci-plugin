package golangciplugin

import (
	"github.com/diphyx/golangci-plugin/analyzer"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("diphyx", NewDiphyxPlugin)
}

// DiphyxPlugin is the golangci-lint module plugin for diphyx code style rules.
type DiphyxPlugin struct{}

// NewDiphyxPlugin creates a new diphyx linter plugin.
func NewDiphyxPlugin(settings any) (register.LinterPlugin, error) {
	diphyxPlugin := &DiphyxPlugin{}

	return diphyxPlugin, nil
}

// BuildAnalyzers returns all diphyx code style analyzers.
func (diphyxPlugin *DiphyxPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return analyzer.All(), nil
}

// GetLoadMode returns the load mode for the plugin.
func (diphyxPlugin *DiphyxPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
