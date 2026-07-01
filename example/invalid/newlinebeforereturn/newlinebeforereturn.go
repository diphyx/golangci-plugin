// Package newlinebeforereturn omits the blank line before a return statement.
package newlinebeforereturn

func compute() int {
	return 0
}

func total() int {
	value := compute()
	return value
}
