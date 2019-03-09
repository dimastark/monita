package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func prompt(c *cli.Context, param string, validator func(string) error) string {
	reader := bufio.NewReader(os.Stdin)

	s := c.String(param)

	var err error

	for err != nil || s == "" {
		if err != nil {
			color.Red("%s", err)
		}

		fmt.Printf("Enter %s: ", param)

		s, _ = reader.ReadString('\n')

		err = validator(strings.TrimSpace(s))
	}

	return strings.TrimSpace(s)
}
