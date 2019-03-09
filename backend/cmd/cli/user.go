package main

import (
	"monita/handlers/user"
	"monita/pkg/prompt"

	"github.com/urfave/cli"
)

// User cli subcommand
func User() cli.Command {
	return cli.Command{
		Name: "user",

		Usage: "manage monita users",

		Subcommands: []cli.Command{
			createUser(),
			deleteUser(),
		},
	}
}

func createUser() cli.Command {
	return cli.Command{
		Name: "create",

		Usage:     "create monita user",
		UsageText: "monita user create [--name -n <name>] [--email -e <email>] [--password -p <password>]",

		Flags: adminSubcommandFlags(
			cli.StringFlag{
				Name:  "name, n",
				Usage: "user name",
			},
			cli.StringFlag{
				Name:  "email, e",
				Usage: "user email",
			},
			cli.StringFlag{
				Name:  "password, p",
				Usage: "user password (insecure, use prompt)",
			},
		),

		Action: func(c *cli.Context) (err error) {
			if checkNotAdmin(c) {
				return
			}

			resp, err := newAdminClient(c).
				CreateUser(user.CreateUserRequest{
					Name:     prompt.Name(c),
					Email:    prompt.Email(c),
					Password: prompt.Password(c),
				})

			prettyPrintlnStruct(resp)

			return
		},
	}
}

func deleteUser() cli.Command {
	return cli.Command{
		Name: "delete",

		Usage:     "delete monita user",
		UsageText: "monita user delete [--name -n <name>]",

		Flags: adminSubcommandFlags(
			cli.StringFlag{
				Name:  "name, n",
				Usage: "user name",
			},
		),

		Action: func(c *cli.Context) (err error) {
			if checkNotAdmin(c) {
				return
			}

			return newAdminClient(c).
				DeleteUser(prompt.Name(c))
		},
	}
}
