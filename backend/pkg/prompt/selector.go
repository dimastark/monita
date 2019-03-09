package prompt

import (
	"monita/pkg/validator"

	"github.com/urfave/cli"
)

// Selector prompts for observable selector
func Selector(c *cli.Context) string {
	return prompt(c, "selector", validator.Selector)
}
