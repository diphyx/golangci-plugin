// Package structnewline contains structs missing the required blank lines.
package structnewline

type badFields struct {
	Name  string
	count int
}

type embedded struct{}

type badEmbed struct {
	embedded
	Name string
}
