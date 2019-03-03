package validator

import (
	"errors"
	"net/http"
)

// URL returns error if the string is not an valid URL
func URL(str string) error {
	res, err := http.Get(str)

	if err != nil {
		return errors.New("url is not a valid")
	}

	defer res.Body.Close()

	return nil
}
