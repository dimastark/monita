package main

import (
	"fmt"

	"monita/pkg/client"

	"github.com/urfave/cli"
)

func newClient(c *cli.Context) *client.MonitaClient {
	return client.NewClient(client.Options{
		BaseURL: c.String("server"),
		Token:   c.String("token"),
	})
}

func newAdminClient(c *cli.Context) *client.MonitaClient {
	return newClient(c).
		SetAdminSecret(c.String("secret"))
}

func subcommandFlags(flags ...cli.Flag) []cli.Flag {
	commonFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "server",
			Value:  APIBaseURL,
			EnvVar: "MONITA_SERVER",
		},
		cli.StringFlag{
			Name:     "token",
			Hidden:   true,
			FilePath: TokenFile,
		},
	}

	return append(commonFlags, flags...)
}

func adminSubcommandFlags(flags ...cli.Flag) []cli.Flag {
	commonFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "secret",
			EnvVar: "ADMIN_SECRET",
		},
	}

	return append(subcommandFlags(flags...), commonFlags...)
}

func checkUnauthorized(c *cli.Context) bool {
	if c.String("token") == "" {
		fmt.Println("You should be authorized. Please run `monita login`.")

		return true
	}

	return false
}

func checkNotAdmin(c *cli.Context) bool {
	if c.String("secret") == "" {
		fmt.Println("You should be admin. Please set ADMIN_SECRET environment variable.")

		return true
	}

	return false
}
