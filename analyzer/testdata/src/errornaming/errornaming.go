package errornaming

func doThing() error { return nil }

func examples() {
	parseError := doThing()
	_ = parseError

	err := doThing() // want "error variable 'err' should use '.functionName.Error' pattern"
	_ = err

	parseErr := doThing() // want "error variable 'parseErr' should end with 'Error' not 'Err'"
	_ = parseErr
}
