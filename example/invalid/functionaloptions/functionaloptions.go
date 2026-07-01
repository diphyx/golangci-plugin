// Package functionaloptions contains an option constructor with a non-With name.
package functionaloptions

// Command is a sample command.
type Command struct{}

// SetName returns a configuration function but is not named WithName.
func SetName(name string) func(*Command) {
	return func(command *Command) {
	}
}
