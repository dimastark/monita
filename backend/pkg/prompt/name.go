package prompt

import (
	"monita/pkg/validator"

	"github.com/urfave/cli"
)

// UserName prompts for user name
func UserName(c *cli.Context) string {
	return prompt(c, "name", validator.Name)
}

// ObservableName prompts for observable name
func ObservableName(c *cli.Context) string {
	return prompt(c, "name", validator.NotEmpty)
}
