package main

import (
	"github.com/urfave/cli"
)

// Notifications cli subcommand
func Notifications() cli.Command {
	return cli.Command{
		Name: "notifications",

		Usage: "manage notifications",

		Subcommands: []cli.Command{
			mute(),
			unmute(),
		},
	}
}

func mute() cli.Command {
	return cli.Command{
		Name: "mute",

		Usage:     "mute notifications. Disables all, if observable is not provided",
		UsageText: "monita notifications mute [--observable -o <observableID>]",

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

			mc := newClient(c)

			if observableID := c.Uint("observable"); observableID != 0 {
				err = mc.MuteObservable(observableID)
			} else {
				err = mc.MuteNotifications()
			}

			return
		},
	}
}

func unmute() cli.Command {
	return cli.Command{
		Name: "unmute",

		Usage:     "unmute notifications. By default enables all, except observable specific",
		UsageText: "monita notifications unmute [--observable -o <observableID>]",

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

			mc := newClient(c)

			if observableID := c.Uint("observable"); observableID != 0 {
				err = mc.UnmuteObservable(observableID)
			} else {
				err = mc.UnmuteNotifications()
			}

			return
		},
	}
}
