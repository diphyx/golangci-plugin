// Package newlineafterdefer omits the blank line after a defer statement.
package newlineafterdefer

func release() {
}

func work() {
	defer release()
	release()
}
