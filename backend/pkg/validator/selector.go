package validator

import (
	"github.com/andybalholm/cascadia"
)

// Selector returns error if string is not an valid selector
func Selector(str string) error {
	if _, err := cascadia.Compile(str); err != nil {
		return err
	}

	return nil
}
