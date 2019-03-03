package validator

import (
	"errors"
)

// NotEmpty returns error if string is empty
func NotEmpty(str string) error {
	if str == "" {
		return errors.New("string is empty")
	}

	return nil
}
