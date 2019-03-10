package main

import (
	"io/ioutil"

	"monita/handlers/user"
	"monita/pkg/prompt"

	"github.com/urfave/cli"
)

// Login cli subcommand
func Login() cli.Command {
	return cli.Command{
		Name: "login",

		Usage:     "logins to the monita",
		UsageText: "monita login [--login -u <user>] [--password -p <password>]",

		Flags: subcommandFlags(
			cli.StringFlag{
				Name:  "login, u",
				Usage: "user login (name or email)",
			},
			cli.StringFlag{
				Name:  "password, p",
				Usage: "user password (insecure, use prompt)",
			},
		),

		Action: func(c *cli.Context) (err error) {
			login, password := prompt.Credentials(c)

			resp, err := newClient(c).Login(user.LoginRequest{
				Login:    login,
				Password: password,
			})

			if err != nil {
				return
			}

			return ioutil.WriteFile(TokenFile, []byte(resp.Token), 0600)
		},
	}
}
