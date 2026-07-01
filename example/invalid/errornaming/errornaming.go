// Package errornaming contains error variables with disallowed names.
package errornaming

func doThing() error {
	return nil
}

func example() {
	err := doThing()
	_ = err

	parseErr := doThing()
	_ = parseErr
}
