package newlinebeforereturn

func compute() int { return 0 }

func examples() int {
	value := compute()
	return value // want "missing blank line before return statement"
}
