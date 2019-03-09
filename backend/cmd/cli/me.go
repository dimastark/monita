package main

import (
	"monita/handlers/user"

	"github.com/urfave/cli"
)

// Me cli subcommand
func Me() cli.Command {
	return cli.Command{
		Name: "me",

		Usage:     "change yours email or password",
		UsageText: "monita me [--email -e <email>] [--password -p <password>]",

		Flags: subcommandFlags(
			cli.StringFlag{
				Name:  "email, e",
				Usage: "new email",
			},
			cli.StringFlag{
				Name:  "password, p",
				Usage: "new password",
			},
		),

		Action: func(c *cli.Context) (err error) {
			if checkUnauthorized(c) {
				return
			}

			mc := newClient(c)

			if email := c.String("email"); email != "" {
				err = mc.ChangeEmail(user.ChangeEmailRequest{
					New: email,
				})
			}

			if err != nil {
				return
			}

			if password := c.String("password"); password != "" {
				err = mc.ChangePassword(user.ChangePasswordRequest{
					New: password,
				})
			}

			return
		},
	}
}
