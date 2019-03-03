package validator

import (
	"errors"
)

// Password returns error if the string is not an valid password
func Password(str string) error {
	if len(str) < 3 {
		return errors.New("password should contains at least 3 characters")
	}

	return nil
}
