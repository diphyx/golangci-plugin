// Package functionaloptions shows a functional option named With{Option}.
package functionaloptions

// Command is a sample command.
type Command struct {
	name string
}

// WithName sets the command name.
func WithName(name string) func(*Command) {
	return func(command *Command) {
		command.name = name
	}
}
