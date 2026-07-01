// Package errorcreation contains misused error constructors.
package errorcreation

import (
	"errors"
	"fmt"
)

func staticMessage() error {
	return fmt.Errorf("service not found")
}

func formattedConstant() error {
	return errors.New("service %s not found")
}
