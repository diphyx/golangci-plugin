package newlineafterdefer

func cleanup() {}

func work() {}

func examples() {
	defer cleanup()
	work() // want "missing blank line after defer statement"
}
