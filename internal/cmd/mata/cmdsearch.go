package mata

import (
	"github.com/urfave/cli"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
	"github.com/tahitianstud/mata-cli/internal/server"
	"github.com/tahitianstud/mata-cli/internal/api"
	"github.com/tahitianstud/mata-cli/internal/search"
	"github.com/tahitianstud/mata-cli/internal/platform/io/survey"
	"fmt"
	"strings"
	"github.com/tahitianstud/mata-cli/internal/platform/errors"
	"strconv"
)

// usage example:
//   mata search -i
//   mata search --api-url matadev.cps.pf:12229 -u toto -p pass
//   mata search --api-url matadev.cps.pf:12229 -u toto -p pass --stream "FAREPREPROD" "*"
//   mata search --target farepreprod "*"

// searchCommand will define the search command
func searchCommand() cli.Command {
	return cli.Command{
		Name:      "search",
		ShortName: "s",
		Usage:     "Searches logs for a particular search",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "api-url",
				Usage: "the Graylog API address (for example: https://graylog.test.com:9000/api)",
			},
			cli.StringFlag{
				Name:  "password, p",
				Usage: "the password to use",
				Value: "",
			},
			cli.StringFlag{
				Name:  "stream",
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

	var serverDefinition server.Definition

	err := assembleServerDefinition(c, &serverDefinition)
	if err != nil {
		return err
	}

	// search information
	query := search.Default()
	terms := ""

	arguments := c.Args()
	if len(arguments) == 1 {
		terms = arguments.Get(0)
	}

	if terms != "" {
		query.Terms = terms
	}

	// stream information
	var streamID = c.String("stream")
	streamID, err = selectStream(streamID, serverDefinition)
	if err != nil {
		return err
	}
	query.Stream = streamID

	log.InfoWith("searching for",
		log.Data("search", query.Terms),
		log.Data("range", query.Range),
		log.Data("stream", query.Stream),
	)

	err = executeSearch(query, serverDefinition)
	if err != nil {
		return err
	}

	return nil
}

func executeSearch(queryDefinition search.Definition, apiServer server.Definition) error {

	graylogAPI := api.FetchProvider(api.GRAYLOG)

	// login or reuse the existing session if present
	err := api.Connect(graylogAPI, server.ConnectionString(apiServer))
	if err != nil {
		return err
	}

	// execute a relative search using the search definition provided
	result := search.Result{}
	err = api.Search(graylogAPI, search.RELATIVE, queryDefinition, &result)
	if err != nil {
		return err
	}

	// print out the result on screen
	for i := len(result.Messages)-1; i >= 0; i-- {
		msg := result.Messages[i]
		m := msg["message"].(map[string]interface{})
		fmt.Printf("%s\n", m["message"])
	}

	return nil

}

// selectStream will prompt the user for the stream he wants to search on
// if streamID is empty
func selectStream(streamID string, apiServer server.Definition) (string, error) {

	if len(streamID) > 0 {
		return streamID, nil
	}

	graylogAPI := api.FetchProvider(api.GRAYLOG)

	err := api.Connect(graylogAPI, server.ConnectionString(apiServer))
	if err != nil {
		return "", err
	}

	// connect to API and get list of streams on which we can execute the search
	streamsResult, err := api.ListStreams(graylogAPI)
	if err != nil {
		return "", err
	}

	if streamsResult.NumberOfStreams <= 0 {
		return "", errors.New("You do not have access to any search stream")
	}

	log.InfoWith("Fetched streams from server",
		log.Data("numberOfStreams", strconv.Itoa(streamsResult.NumberOfStreams)),
	)

	var streamOptions = make([]string, len(streamsResult.Streams))
	for i, stream := range streamsResult.Streams {
		streamOptions[i] = fmt.Sprintf("%s | %s (%s)", stream.ID, stream.Title, stream.Description)
	}

	streamChoice := survey.AskForSelection("Choose a stream to search in", streamOptions)

	// parse the choice to get only the stream ID
	splittedStringChoice := strings.Split(streamChoice, "|")

	streamID = strings.Trim(splittedStringChoice[0], " ")


	log.InfoWith("Selected stream to search on",
		log.Data("stream", streamID),
	)


	return streamID, nil

}
