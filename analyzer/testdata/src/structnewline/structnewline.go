package structnewline

type Good struct {
	Name string

	count int
}

type Bad struct {
	Name  string
	count int // want "missing blank line before private fields"
}

type embedded struct{}

type BadEmbed struct {
	embedded
	Name string // want "missing blank line after embedded field"
}
