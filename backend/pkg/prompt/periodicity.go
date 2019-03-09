package prompt

import (
	"monita/pkg/validator"

	"github.com/urfave/cli"
)

// Periodicity prompts for observable periodicity
func Periodicity(c *cli.Context) string {
	return prompt(c, "periodicity", validator.Periodicity)
}
