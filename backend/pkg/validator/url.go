package validator

import (
	"errors"
	"net/http"
)

// URL returns error if the string is not an valid URL
func URL(str string) error {
	resp, err := http.Get(str)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return errors.New("invalid url")
	}

	return nil
}
