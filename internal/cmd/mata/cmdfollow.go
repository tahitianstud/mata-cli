package mata

import (
	"github.com/urfave/cli"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
	"github.com/tahitianstud/mata-cli/internal/server"
	"github.com/tahitianstud/mata-cli/internal/api"
	"github.com/tahitianstud/mata-cli/internal/search"
	"fmt"
	"github.com/tahitianstud/mata-cli/internal/platform/date"
	"time"
)

// TODO
// usage example:

func followCommand() cli.Command {
	return cli.Command{
		Name:      "follow",
		ShortName: "t",
		Usage:     "Outputs logs in a follow-mode similar to follow",
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
		Action: doFollow,
	}
}

// doFollow will launch the follow operation
func doFollow(c *cli.Context) error {

	// SERVER DEFINITION

	var serverDefinition server.Definition

	err := assembleServerDefinition(c, &serverDefinition)
	if err != nil {
		return err
	}

	// SEARCH DEFINITION

	var searchDefinition search.Definition

	searchDefinition = search.Default()

	// stream information
	var streamID = c.String("stream")
	streamID, err = selectStream(streamID, serverDefinition)
	if err != nil {
		return err
	}
	searchDefinition.Stream = streamID

	log.InfoWith("searching for",
		log.Data("search", searchDefinition.Terms),
		log.Data("range", searchDefinition.Range),
		log.Data("stream", searchDefinition.Stream),
	)

	return follow(searchDefinition, serverDefinition)
}

func follow(searchDefinition search.Definition, apiServer server.Definition) error {

	graylogAPI := api.FetchProvider(api.GRAYLOG)

	err := api.Connect(graylogAPI, server.ConnectionString(apiServer))
	if err != nil {
		return err
	}

	fmt.Printf("\n----------------------------------------------------------------------------------------------------\n\n")

	latency := 1 * time.Second      // 1 second latency in order not to miss any message
	searchWindow := 1 * time.Second // 1 second search searchWindow

	// initial time range
	toTime := time.Now().Add(-latency)
	fromTime := toTime.Add(-searchWindow)

	// loop indefinitely while logging messages
	for {

		result := search.Result{}

		searchDefinition.From = date.DateToStr(fromTime.UTC(), "2006-01-02T15:04:05.000Z")
		searchDefinition.To = date.DateToStr(toTime.UTC(), "2006-01-02T15:04:05.000Z")

		err = api.Search(graylogAPI, search.ABSOLUTE, searchDefinition, &result)
		if err != nil {
			return err
		}

		// print out the result on screen
		for i := len(result.Messages) - 1; i >= 0; i-- {
			msg := result.Messages[i]
			m := msg["message"].(map[string]interface{})
			fmt.Printf("%s\n", m["message"])
		}

		time.Sleep(latency)

		// update range accordingly correctly
		fromTime, _ = date.StrToDate(result.To)
		toTime = fromTime.Add(searchWindow)

	}

	return nil

}
