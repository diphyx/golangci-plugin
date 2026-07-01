// Package errorcreation shows the correct error constructor for each case.
package errorcreation

import (
	"errors"
	"fmt"
)

// Validate uses errors.New for a static message and fmt.Errorf for a formatted one.
func Validate(name string) error {
	if name == "" {
		return errors.New("missing name")
	}

	return fmt.Errorf("invalid name '%s'", name)
}
