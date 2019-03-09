package prompt

import (
	"fmt"
	"strings"
	"syscall"

	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

// Password prompts for user password
func Password(c *cli.Context) string {
	var err error
	var bytePassword []byte

	password := c.String("password")

	for err != nil || strings.TrimSpace(password) == "" {
		fmt.Print("Enter password:")

		bytePassword, err = terminal.ReadPassword(int(syscall.Stdin))

		password = string(bytePassword)
	}

	fmt.Println()

	return strings.TrimSpace(password)
}
