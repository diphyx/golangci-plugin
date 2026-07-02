package ifinitstatement

func doWork() error {
	return nil
}

func check() {
	callError := doWork()
	if callError != nil { // declared before the if is fine
		println("handled")
	}

	if callError := doWork(); callError != nil { // want "if statement should not use an init statement"
		println("error check in the init statement is flagged")
	}
}

func lookup(byName map[string]int, step string) {
	_, isPresent := byName[step]
	if !isPresent {
		println("declared before the if is fine")
	}

	if _, isMissing := byName[step]; !isMissing { // want "if statement should not use an init statement"
		println("comma-ok in the init statement is flagged")
	}

	if value := byName[step]; value > 0 { // want "if statement should not use an init statement"
		println(value)
	}

	if value := byName[step]; value > 0 { // want "if statement should not use an init statement"
		println("first")
	} else if _, isOther := byName["other"]; isOther { // want "if statement should not use an init statement"
		println("else-if init statement is flagged too")
	}
}
