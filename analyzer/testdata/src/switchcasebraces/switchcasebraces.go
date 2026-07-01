package switchcasebraces

func expressionSwitch(value int) {
	switch value {
	case 1:
		{
			println("one")
		}
	case 2:
		println("two") // want "should be wrapped in braces"
	case 3:
		value++ // want "should be wrapped in braces"
		println(value)
	case 4: // empty body is skipped
	case 5:
		fallthrough // fallthrough cannot be wrapped, so it is skipped
	case 6:
		{
			println("six")
		}
	default:
		println("other") // want "should be wrapped in braces"
	}
}

func typeSwitch(value any) {
	switch value.(type) {
	case int:
		{
			println("int")
		}
	case string:
		println("string") // want "should be wrapped in braces"
	}
}
