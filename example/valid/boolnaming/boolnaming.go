// Package boolnaming shows boolean variables with an approved prefix.
package boolnaming

// Evaluate demonstrates correctly named boolean variables.
func Evaluate() bool {
	isReady := true
	hasItems := false

	if isReady && hasItems {
		return true
	}

	return false
}
