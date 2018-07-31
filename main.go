package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/sirupsen/logrus"
	"github.com/tahitianstud/mata-cli/commands"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "mata"
	app.Usage = "convenient Graylog output at the cli"
	app.Description = "cli utility used to output logs from a Graylog server"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug, d",
			Usage:  "output debug messages",
			EnvVar: "DEBUG",
		},
		cli.StringFlag{
			Name:   "target",
			Usage:  "target specific settings",
			EnvVar: "TARGET",
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.GlobalBool("debug") == true {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}

		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:      "config",
			Usage:     "Configures global settings",
			Action:    commands.Config,
			UsageText: fmt.Sprintf("%s config", app.Name),
		},
		{
			Name:      "search",
			ShortName: "s",
			Usage:     "Searches logs for a particular query",
		},
		{
			Name:      "setup",
			Usage:     "Sets up the application's preferences",
			UsageText: fmt.Sprintf("%s --target <settings> setup", app.Name),
			Action:    commands.Setup,
		},
		{
			Name:      "trace",
			ShortName: "t",
			Usage:     "Outputs logs in follow fashion",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
