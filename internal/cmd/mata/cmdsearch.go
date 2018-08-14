package mata

import (
	"github.com/urfave/cli"
	"github.com/tahitianstud/mata-cli/internal/server"
	"github.com/tahitianstud/mata-cli/internal/api"
	"github.com/tahitianstud/mata-cli/internal/search"
	"fmt"
)

// usage example:
//   mata search "*" (if already logged in)
//   mata search --api-url matadev.cps.pf:12229 -u user -p pass
//   mata search --api-url matadev.cps.pf:12229 -u user -p pass --stream "0000000000001" "*"

func searchCommand() cli.Command {
	return cli.Command{

		Name: "search",

		ShortName: "s",

		Usage: "Searches logs for a particular search",

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

		Action: func(c *cli.Context) error {

			err := doSearchUsing(serverDefinitionPrompt, searchDefinitionPrompt, c)

			if err != nil {
				return err
			}

			return nil

		},
	}
}

func doSearchUsing(defineServerWith serverDefiner, defineSearchWith searchDefiner, c *cli.Context) error {

	serverDef, err := defineServerWith(c)
	if err != nil {
		return err
	}

	searchDef, err := defineSearchWith(c, serverDef)
	if err != nil {
		return err
	}

	graylogAPI := api.FetchProvider(api.GRAYLOG)

	// login or reuse the existing session if present
	err = api.Connect(graylogAPI, server.ConnectionString(serverDef))
	if err != nil {
		return err
	}

	// execute a relative search using the search definition provided
	result := search.Result{}
	err = api.Search(graylogAPI, search.RELATIVE, searchDef, &result)
	if err != nil {
		return err
	}

	// print out the result on screen
	for i := len(result.Messages) - 1; i >= 0; i-- {
		msg := result.Messages[i]
		m := msg["message"].(map[string]interface{})
		fmt.Printf("%s\n", m["message"])
	}

	return nil
}
