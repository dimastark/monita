package validator

import (
	"errors"
)

var periodicities = map[string]bool{
	"manually":  true,
	"regularly": true,
	"daily":     true,
	"weekly":    true,
}

// Periodicity returns error if string is not an valid periodicity
func Periodicity(str string) error {
	if _, ok := periodicities[str]; !ok {
		return errors.New("periodicity should be one of {manually, regularly, daily, weekly}")
	}

	return nil
}
