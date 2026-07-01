// Package newlineafterdefer shows a blank line after a defer statement.
package newlineafterdefer

func release() {
}

// Work runs with a deferred cleanup followed by a blank line.
func Work() {
	defer release()

	release()
}
