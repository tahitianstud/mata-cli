package mata

import (
	"github.com/tahitianstud/mata-cli/internal/server"
	"github.com/tahitianstud/mata-cli/internal/api"
	"github.com/tahitianstud/mata-cli/internal/platform/io/survey"
	"github.com/urfave/cli"
	"github.com/tahitianstud/mata-cli/internal/platform/errors"
)

func assembleServerDefinition(c *cli.Context, serverDefinition *server.Definition) (err error) {

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
			err := server.UpdateConn(&connection, apiURL)

			if err != nil {
				return errors.Wrap(err, "cannot use URL")
			}
		}

	}

	// define the server
	serverDefinition.Connection = connection
	serverDefinition.User = user

	return nil
}
