// Package switchcasebraces contains a switch case body not wrapped in braces.
package switchcasebraces

func example(value int) {
	switch value {
	case 1:
		println("one")
	default:
		println("other")
	}
}
