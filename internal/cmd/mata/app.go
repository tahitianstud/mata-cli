// Package mata holds the definition of the mata cli application
// and its various commands.
package mata

import (
	"github.com/urfave/cli"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
	"os"
)

type cliApp struct {
	application *cli.App
}

func createCLI(version string) cliRunner {
	app := cli.NewApp()

	// APP DESCRIPTION

	app.Name = "mata"
	app.Usage = "convenient logging output at the cli"
	app.Description = "cli utility used to output logs from a logging server"
	app.Version = version

	// GLOBAL FLAGS

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug, d",
			Usage:  "output debug messages",
			EnvVar: "DEBUG",
		},
	}
	app.Before = func(c *cli.Context) error {

		// if debug activated, set loglevel to DEBUG, else set to INFO
		if c.GlobalBool("debug") == true {
			log.SetLevelTo(log.DEBUG)
		} else {
			log.SetLevelTo(log.INFO)
		}

		return nil
	}

	// SUB-COMMANDS

	app.Commands = []cli.Command{
		loginCommand(),
		logoutCommand(),
		searchCommand(),
		followCommand(),
	}

	return cliApp{
		application: app,
	}
}

func (c cliApp) run() {

	// ERROR HANDLING

	// if an error bubbles its way up, then exit application with Fatal error
	err := c.application.Run(os.Args)

	log.DieIf(err)
}
