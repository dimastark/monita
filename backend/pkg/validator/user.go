package validator

import (
	"errors"

	"monita/storage/user"
)

// LoginExist returns error if user with login does not exist
func LoginExist(login string) error {
	if user.GetByLogin(login) == nil {
		return errors.New("user does not exist")
	}

	return nil
}

// EmailNotExist returns error if user with email exist
func EmailNotExist(email string) error {
	if user.GetByEmail(email) != nil {
		return errors.New("email already exists")
	}

	return nil
}

// NameOrEmailNotExist returns error if user with name or email exist
func NameOrEmailNotExist(name, email string) error {
	if user.GetByNameOrEmail(name, email) != nil {
		return errors.New("name or email already exists")
	}

	return nil
}
