package prompt

import (
	"github.com/urfave/cli"
)

// Credentials prompts for user login and password
func Credentials(c *cli.Context) (string, string) {
	return Login(c), Password(c)
}
