// Package ifinitstatement contains an if statement that uses an init statement.
package ifinitstatement

func example(byName map[string]int, step string) bool {
	if _, isPresent := byName[step]; !isPresent {
		return false
	}

	return true
}
