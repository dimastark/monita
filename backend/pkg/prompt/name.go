package prompt

import (
	"monita/pkg/validator"

	"github.com/urfave/cli"
)

// Name prompts for user name
func Name(c *cli.Context) string {
	return prompt(c, "name", validator.Name)
}
