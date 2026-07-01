// Package switchcasebraces shows switch case bodies wrapped in braces.
package switchcasebraces

// Describe prints a label for the given value.
func Describe(value int) {
	switch value {
	case 1:
		{
			println("one")
		}
	default:
		{
			println("other")
		}
	}
}
