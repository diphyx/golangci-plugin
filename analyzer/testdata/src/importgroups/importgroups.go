package importgroups

import (
	"fmt"

	"example.com/foo"

	"os" // want "standard library import 'os' should come before third-party imports"
)

var _ = fmt.Sprint
var _ = foo.Value
var _ = os.Args
