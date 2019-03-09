package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "monita"
	app.HelpName = "monita"

	app.Usage = "cli for web pages monitoring"
	app.UsageText = "monita command [command options]"

	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		Login(),
		User(),
		Notifications(),
		Me(),
		Observable(),
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
