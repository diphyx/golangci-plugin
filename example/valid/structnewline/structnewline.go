// Package structnewline shows structs with the required blank lines.
package structnewline

// Config holds settings with a blank line before its private fields.
type Config struct {
	Name string

	count int
}

type embedded struct{}

// Wrapper embeds a type and separates it with a blank line.
type Wrapper struct {
	embedded

	Name string
}
