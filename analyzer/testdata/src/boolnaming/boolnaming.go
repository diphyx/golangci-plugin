package boolnaming

func examples() {
	isReady := true
	_ = isReady

	hasItems := false
	_ = hasItems

	ok := true
	_ = ok

	ready := false // want "boolean variable 'ready' should have a prefix: is, has, can, should, enable, or in"
	_ = ready
}
