package prompt

import (
	"monita/pkg/validator"

	"github.com/urfave/cli"
)

// Email prompts for user email
func Email(c *cli.Context) string {
	return prompt(c, "email", validator.Email)
}
