// Package ifinitstatement shows if statements without an init statement.
package ifinitstatement

// Lookup reports whether the given step exists in the map.
func Lookup(byName map[string]int, step string) bool {
	_, isPresent := byName[step]
	if !isPresent {
		return false
	}

	return true
}
