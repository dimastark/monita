package validator

import (
	"errors"
	"regexp"
)

var nameRe = regexp.MustCompile("[a-z0-9_-]{3,}")

// Name returns error if the string is not an valid user name
func Name(str string) error {
	if !nameRe.MatchString(str) {
		return errors.New("name should match [a-z0-9_-]{3,}")
	}

	return nil
}
