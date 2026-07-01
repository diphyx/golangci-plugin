// Package errornaming shows an error variable using the '{functionName}Error' pattern.
package errornaming

func doThing() error {
	return nil
}

// Run demonstrates a correctly named error variable.
func Run() error {
	doThingError := doThing()

	return doThingError
}
