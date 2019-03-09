package prompt

import (
	"monita/pkg/validator"

	"github.com/urfave/cli"
)

// URL prompts for observable url
func URL(c *cli.Context) string {
	return prompt(c, "url", validator.URL)
}
