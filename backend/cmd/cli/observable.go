package main

import (
	"fmt"

	"monita/handlers/observable"
	"monita/pkg/prompt"

	"github.com/urfave/cli"
)

// Observable cli subcommand
func Observable() cli.Command {
	return cli.Command{
		Name: "observable",

		Usage: "manage and handle observables",

		Subcommands: []cli.Command{
			list(),
			create(),
			delete(),
			load(),
		},
	}
}

func list() cli.Command {
	return cli.Command{
		Name: "list",

		Usage:     "list all observables",
		UsageText: "monita observable list",

		Flags: subcommandFlags(),

		Action: func(c *cli.Context) (err error) {
			if checkUnauthorized(c) {
				return
			}

			u, err := newClient(c).CurrentUser()

			if err != nil {
				return
			}

			for _, o := range u.Observables {
				o.Print()
				fmt.Println()
			}

			return nil
		},
	}
}

func create() cli.Command {
	return cli.Command{
		Name: "create",

		Usage:     "create observable",
		UsageText: "monita user create [--name -n <name>] [--periodicity -p <periodicity>] [--url -u <url>] [--selector -s <selector>]",

		Flags: subcommandFlags(
			cli.StringFlag{
				Name:  "name, n",
				Usage: "observable name",
			},
			cli.StringFlag{
				Name:  "perioricity, p",
				Usage: "observable perioricity",
			},
			cli.StringFlag{
				Name:  "url, u",
				Usage: "observable url",
			},
			cli.StringFlag{
				Name:  "selector, s",
				Usage: "observable selector",
			},
		),

		Action: func(c *cli.Context) (err error) {
			if checkUnauthorized(c) {
				return
			}

			_, err = newClient(c).
				CreateObservable(observable.CreateObservableRequest{
					Name:        prompt.Name(c),
					Periodicity: prompt.Periodicity(c),
					URL:         prompt.URL(c),
					Selector:    prompt.Selector(c),
				})

			return
		},
	}
}

func delete() cli.Command {
	return cli.Command{
		Name: "delete",

		Usage:     "delete observable",
		UsageText: "monita observable delete [--observable -o <observableID>]",

		Flags: subcommandFlags(
			cli.UintFlag{
				Name:  "observable, o",
				Usage: "observable ID",
			},
		),

		Action: func(c *cli.Context) (err error) {
			if checkUnauthorized(c) {
				return
			}

			return newClient(c).
				DeleteObservable(c.Uint("observable"))
		},
	}
}

func load() cli.Command {
	return cli.Command{
		Name: "load",

		Usage:     "load observable updates",
		UsageText: "monita observable load [--resolve]",

		Flags: subcommandFlags(
			cli.BoolFlag{
				Name:  "resolve, r",
				Usage: "mark updates as resolved",
			},
		),

		Action: func(c *cli.Context) (err error) {
			if checkUnauthorized(c) {
				return
			}

			mc := newClient(c)

			u, err := mc.CurrentUser()

			if err != nil {
				return
			}

			for _, o := range u.Observables {
				if o.Periodicity != "manually" {
					continue
				}

				resp, err := mc.LoadObservable(o.ID)

				if err != nil {
					continue
				}

				if resp.LastData != resp.NewData {
					resp.Print()
					fmt.Println()

					if c.Bool("handle") {
						mc.HandleObservable(o.ID)
					}
				}
			}

			return
		},
	}
}
