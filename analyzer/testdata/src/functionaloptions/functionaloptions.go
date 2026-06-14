package functionaloptions

type Command struct{}

// WithName sets the command name.
func WithName(name string) func(*Command) {
	return func(command *Command) {}
}

// SetName sets the command name.
func SetName(name string) func(*Command) { // want "function 'SetName' returns a configuration function, should be named 'With.Option.'"
	return func(command *Command) {}
}
