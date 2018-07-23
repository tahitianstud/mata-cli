package mata

import (
	"github.com/urfave/cli"
	"github.com/tahitianstud/mata-cli/internal/platform/io/survey"
	"github.com/tahitianstud/mata-cli/internal/api"
	"github.com/tahitianstud/mata-cli/internal/server"
	"github.com/pkg/errors"
	"github.com/tahitianstud/mata-cli/internal/platform/io/file"
	"github.com/tahitianstud/mata-cli/internal"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
	"github.com/tahitianstud/mata-cli/internal/platform/json"
)

// loginCommand will define the login command
func loginCommand() cli.Command {
	return cli.Command{
		Name:  "login",
		Usage: "Logs in to a specified server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "username, u",
				Usage: "the username to use",
			},
			cli.StringFlag{
				Name:  "password, p",
				Usage: "the password to use",
			},
		},
		ArgsUsage: "<SERVER_URL> (e.g https://server.domain.com:3489/api)",
		Action:    doLogin,
	}
}

func doLogin(c *cli.Context) error {

	args := c.Args()
	var URL string
	if len(args) < 1 {
		URL = survey.AskForInput("URL", "e.g. https://server.domain.com:4383/api")
	} else {
		URL = args.Get(0)
	}

	username := c.String("username")
	if len(username) <= 0 {
		username = survey.AskForInput("Username")
	}

	password := c.String("password")
	if len(password) <= 0 {
		password = survey.AskForPassword("Password")
	}

	serverDefinition := server.Definition{
		User: server.User{
			Username: username,
			Password: password,
		},
	}

	err := server.UpdateConn(&serverDefinition.Connection, URL)
	if err != nil {
		return errors.Wrap(err, "invalid server URL (should be something like https://server.domain.com:3944/api")
	}

	// fetch a Graylog API provider in order to execute the login command
	graylogAPI := api.FetchProvider(api.GRAYLOG)

	// execute a Connect API call that will login and save a session on disk
	err = api.Connect(graylogAPI, server.ConnectionString(serverDefinition))
	if err != nil {
		return errors.Wrap(err, "could not connect and login to server")
	}

	// check if session has been correctly saved to disk
	sessionOnDisk := ""
	err = file.ReadFrom(internal.ConfigLocation, ".session", &sessionOnDisk)
	if err != nil || len(sessionOnDisk) <= 0 {
		return errors.Wrap(err, "login failed: no session saved on disk")
	}

	var ticket interface{}
	err = api.DeserializeSession(sessionOnDisk, &ticket)
	if err != nil {
		return errors.Wrap(err, "login failed: invalid session information saved to disk")
	}

	log.InfoWith("login successful", log.Data("ticket", json.SerializeToString(ticket)))

	return nil
}
