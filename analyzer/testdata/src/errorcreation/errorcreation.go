package errorcreation

import (
	"errors"
	"fmt"
)

func staticMessage() error {
	return fmt.Errorf("service not found") // want "use errors.New for static error messages"
}

func formattedConstant() error {
	return errors.New("service %s not found") // want "use fmt.Errorf for formatted error messages"
}

func correct(name string) error {
	if name == "" {
		return errors.New("missing name")
	}

	return fmt.Errorf("service '%s' not found", name)
}
