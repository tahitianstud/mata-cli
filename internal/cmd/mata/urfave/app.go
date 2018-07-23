package urfave

import (
	"github.com/urfave/cli"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
	"os"
	"github.com/tahitianstud/mata-cli/internal"
)

type CLIImplementation struct {
	application *cli.App
}

func CreateCLI() CLIImplementation {
	app := cli.NewApp()
	app.Name = internal.AppName
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

	// deal with global flags:
	//   - if debug activated, set loglevel to DEBUG, else set to INFO
	app.Before = func(c *cli.Context) error {
		if c.GlobalBool("debug") == true {
			log.SetLevelTo(log.DEBUG)
		} else {
			log.SetLevelTo(log.INFO)
		}

		return nil
	}

	app.Commands = []cli.Command{
		ConfigCommand(),
		SearchCommand(),
		SetupCommand(),
	}

	return CLIImplementation{
		application: app,
	}
}

func (c CLIImplementation) Run() {
	// Error Handling:
	// if an error bubbles its way up, then
	// exit application with Fatal error
	err := c.application.Run(os.Args)
	log.DieIf(err)
}