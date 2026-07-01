// Package importgroups shows standard imports before third-party imports.
package importgroups

import (
	"fmt"

	"github.com/diphyx/golangci-plugin/example/thirdparty"
)

// Describe uses a standard and a third-party import in the correct order.
func Describe() string {
	return fmt.Sprintf("value=%d", thirdparty.Value)
}
