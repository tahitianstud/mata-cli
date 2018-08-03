package mata

import (
	"github.com/tahitianstud/mata-cli/internal/server"
	"github.com/tahitianstud/mata-cli/internal/api"
	"github.com/tahitianstud/mata-cli/internal/platform/io/survey"
	"github.com/urfave/cli"
	"github.com/tahitianstud/mata-cli/internal/platform/errors"
	"github.com/tahitianstud/mata-cli/internal/search"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
)

type serverDefiner func(c *cli.Context) (serverDefinition server.Definition, err error)
type searchDefiner func(c *cli.Context, srvDef server.Definition) (s search.Definition, err error)

func serverDefinitionPrompt(c *cli.Context) (serverDefinition server.Definition, err error) {

	var connection server.Connection
	var user server.User

	// user information
	user.Username = c.String("username")
	user.Password = c.String("password")

	// if no session exists, connection, username and password are mandatory
	if !api.SessionExists() {

		// user
		if len(user.Username) == 0 {
			user.Username = survey.AskForInput("Username")
		}
		if len(user.Password) == 0 {
			user.Password = survey.AskForPassword("Password")
		}

		// connection information
		connection = server.DefaultConnection()

		apiURL := c.String("api-url")
		if len(apiURL) == 0 {
			survey.AskQuestionsForConfig(&connection)
		} else {
			err = server.UpdateConn(&connection, apiURL)

			if err != nil {
				return serverDefinition, errors.Wrap(err, "cannot use URL")
			}
		}

	}

	// define the server
	serverDefinition.Connection = connection
	serverDefinition.User = user

	return serverDefinition, nil
}

func searchDefinitionPrompt(c *cli.Context, srvDef server.Definition) (s search.Definition, err error) {

	// search information
	s = search.Default()
	terms := ""

	arguments := c.Args()
	if len(arguments) == 1 {
		terms = arguments.Get(0)
	}

	if terms != "" {
		s.Terms = terms
	}

	// stream information
	var streamID = c.String("stream")
	streamID, err = selectStream(streamID, srvDef)
	if err != nil {
		return s, err
	}
	s.Stream = streamID

	log.InfoWith("searching for",
		log.Data("search", s.Terms),
		log.Data("range", s.Range),
		log.Data("stream", s.Stream),
	)

	return s, nil

}