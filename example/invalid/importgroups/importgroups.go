// Package importgroups contains a standard import placed after a third-party one.
package importgroups

import (
	"fmt"

	"github.com/diphyx/golangci-plugin/example/thirdparty"

	"os"
)

func example() {
	_ = fmt.Sprint
	_ = thirdparty.Value
	_ = os.Args
}
