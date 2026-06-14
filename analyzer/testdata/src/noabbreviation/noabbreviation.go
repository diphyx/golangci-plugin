package noabbreviation

func examples() {
	ctx := 0 // want "use 'context' instead of 'ctx'"
	_ = ctx

	request := 0
	_ = request
}
