package prompt

import (
	"monita/pkg/validator"

	"github.com/urfave/cli"
)

// Login prompts for user login
func Login(c *cli.Context) string {
	return prompt(c, "login", validator.NotEmpty)
}
