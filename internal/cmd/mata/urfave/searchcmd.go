package urfave

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/urfave/cli"
	"github.com/tahitianstud/mata-cli/internal/mata"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
)

// usage example:
//   mata search -i
//   mata search --api-url matadev.cps.pf:12229 -u toto -p pass
//   mata search --api-url matadev.cps.pf:12229 -u toto -p pass --stream "FAREPREPROD" "*"
//   mata search --target farepreprod "*"

// SearchCommand will define the search command
func SearchCommand() cli.Command {
	return cli.Command{
		Name:      "search",
		ShortName: "s",
		Usage:     "Searches logs for a particular query",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "api-url",
				Usage: "the Graylog API address (for example: https://graylog.test.com:9000/api)",
			},
			cli.BoolFlag{
				Name:  "interactive, i",
				Usage: "ask user questions to create connection and search query",
			},
			cli.StringFlag{
				Name:  "password, p",
				Usage: "the password to use",
				Value: "",
			},
			cli.StringFlag{
				Name:  "stream-id",
				Usage: "the stream to search on",
			},
			cli.StringFlag{
				Name:  "username, u",
				Usage: "the username to use",
			},
		},
		Action: doSearch,
	}
}

// doSearch will launch the search operation
func doSearch(c *cli.Context) error {
	var interactiveMode = c.Bool("interactive")

	if interactiveMode == true {
		err := doInteractiveSearch(c)
		if err != nil {
			return err
		}
	} else {
		err := doNonInteractiveSearch(c)
		if err != nil {
			return err
		}
	}

	return nil
}

// doInteractiveSearch will prompt the user for information to use for the search
func doInteractiveSearch(c *cli.Context) error {

	server := mata.DefaultServer()

	mata.AskQuestionsForConfig(&server)

	log.InfoWith("connecting to server",
		log.Data("URL", server.GetURL()),
		log.Data("user", server.Username),
	)

	query := mata.DefaultQuery()

	mata.AskQuestionsForConfig(&query)

	var streamID = c.String("stream-id")
	streamID, err := selectStream(streamID, server)

	if err != nil {
		return err
	}

	log.InfoWith("searching for",
		log.Data("query", query.Terms),
	)

	return nil
}

// doNonInteractiveSearch will execute a search query on the Graylog server
// using the provided server and query information
func doNonInteractiveSearch(c *cli.Context) (err error) {

	// deal with cli options
	var apiURL = c.String("api-url")
	var password = c.String("password")
	var streamID = c.String("stream-id")
	var username = c.String("username")

	server := mata.DefaultServer()

	err = server.UpdateWithURL(apiURL)
	if err != nil {
		return err
	}

	err = server.UpdateWithUsername(username)
	if err != nil {
		return err
	}

	if len(password) <= 0 {
		password = mata.AskForPassword("Please enter your password")
	}
	err = server.UpdateWithPassword(password)
	if err != nil {
		return err
	}


	log.InfoWith("connecting to server",
		log.Data("URL", server.GetURL()),
		log.Data("user", server.Username),
	)


	query := mata.DefaultQuery()

	streamID, err = selectStream(streamID, server)
	if err != nil {
		return err
	}


	log.InfoWith("search for",
		log.Data("query", query.Terms),
	)


	return err
}

// selectStream will prompt the user for the stream he wants to search on
// if streamID is empty
func selectStream(streamID string, server mata.Server) (string, error) {

	if len(streamID) > 0 {

		log.InfoWith("Execute search on",
			log.Data("stream", streamID),
		)

		return streamID, nil
	}

	// connect to API and get list of streams on which we can execute the query
	streamsResult, err := mata.Connect(server).ListEnabledStreams()
	if err != nil {
		return "", err
	}

	if streamsResult.NumberOfStreams <= 0 {
		return "", errors.New("You do not have access to any search stream")
	}


	log.InfoWith("Fetched streams from server",
		log.Data("numberOfStreams", streamsResult.NumberOfStreams),
	)


	var streamOptions = make([]string, len(streamsResult.Streams))
	for i, stream := range streamsResult.Streams {
		streamOptions[i] = fmt.Sprintf("%s | %s (%s)", stream.ID, stream.Title, stream.Description)
	}

	streamChoice := mata.AskForSelection("Choose a stream to search in", streamOptions)

	// parse the choice to get only the stream ID
	splittedStringChoice := strings.Split(streamChoice, "|")

	streamID = strings.Trim(splittedStringChoice[0], " ")


	log.InfoWith("Selected stream to search on",
		log.Data("stream", streamID),
	)


	return streamID, nil

}
