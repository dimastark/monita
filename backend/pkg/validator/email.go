package validator

import (
	"errors"
	"strings"
)

// Email returns error if the string is not an valid email
func Email(str string) error {
	if !strings.Contains(str, "@") {
		return errors.New("email should contains @")
	}

	return nil
}
