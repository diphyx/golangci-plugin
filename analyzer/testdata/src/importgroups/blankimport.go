package importgroups

import (
	_ "example.com/foo" // want "blank import 'example.com/foo' should be at the end of the imports"

	"strings"
)

var _ = strings.TrimSpace
