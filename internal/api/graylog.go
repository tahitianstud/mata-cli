package api

import (
	"time"
	"fmt"
	"github.com/tahitianstud/mata-cli/internal/api/graylog"
	"github.com/tahitianstud/mata-cli/internal/platform/date"
	"github.com/tahitianstud/mata-cli/internal/platform/errors"
	"github.com/tahitianstud/mata-cli/internal/search"
	"github.com/tahitianstud/mata-cli/internal/stream"
	"strconv"
)

type graylogAPI struct {
	API *graylog.API
}

func fetchGraylogAPI() *graylogAPI {
	return &graylogAPI{
		API: &graylog.API{},
	}
}

func (g *graylogAPI) login(connectionString string) error {
	return g.API.Login(connectionString)
}

func (g *graylogAPI) fetchSession() (session session) {

	graylogSession := g.API.FetchSession()

	session.ConnectionString = graylogSession.ConnectionString
	session.SessionID = graylogSession.Ticket.SessionID
	session.ValidUntil = graylogSession.Ticket.ValidUntil

	return session

}

func (g *graylogAPI) restoreSession(session session) error {

	graylogSession := graylog.Session{}

	// check expiration date
	expiryDate, err := date.StrToDate(session.ValidUntil)
	if err != nil {
		return err
	}

	now := time.Now()
	if expiryDate.Before(now) {
		return errors.New("cannot restore expired session")
	}

	graylogSession.ConnectionString = session.ConnectionString
	graylogSession.Ticket.SessionID = session.SessionID
	graylogSession.Ticket.ValidUntil = session.ValidUntil

	g.API.RestoreSession(graylogSession)

	return nil

}

func (g *graylogAPI) sessionActive() bool {

	sessionID := g.API.Ticket.SessionID
	validUntil := g.API.Ticket.ValidUntil

	// format expected: 2018-07-27T06:42:57.528+0000
	expirationDate, err := date.StrToDate(validUntil)
	if err != nil {
		return false
	}

	now := time.Now()

	return len(sessionID) > 0 && len(validUntil) > 0 && expirationDate.After(now)
}

func (g *graylogAPI) listStreams() (stream.StreamsList, error) {
	var graylogStreamsList graylog.StreamsList

	graylogStreamsList, err := g.API.ListStreams()

	if err != nil {
		return stream.StreamsList{}, err
	}

	streamsList := convertStreams(graylogStreamsList)

	return streamsList, nil
}

func convertStreams(graylogList graylog.StreamsList) (streams stream.StreamsList) {

	streams.NumberOfStreams = graylogList.NumberOfStreams
	arrayOfStreams := graylogList.Streams

	for _, graylogStream := range arrayOfStreams {
		stream := stream.Stream{
			Title:       graylogStream.Title,
			Description: graylogStream.Description,
			ID:          graylogStream.ID,
		}

		streams.Streams = append(streams.Streams, stream)
	}

	return streams
}

// search will convert the exported function into a graylog flavored call
func (g *graylogAPI) search(query search.Definition, result *search.Result) error {

	searchDefinition := graylog.Search{
		Query: query.Terms,
		Range: query.Range,
		Sort:  "timestamp:desc",
	}

	if query.Stream != "" {
		searchDefinition.Filter = fmt.Sprintf("streams:%s", query.Stream)
	}

	graylogResult, err := g.API.Search(searchDefinition)

	if err != nil {
		return err
	}

	// convert messages into messages
	for _, message := range graylogResult.Messages {

		//jsonStr, err := json.Marshal(message)

		if err != nil {
			return err
		}

		//log.InfoWith("debug", log.Data("test", jsonStr))

		messageDetails := message.(map[string]interface{})
		result.Messages = append(result.Messages, messageDetails)
	}

	return nil

}

// searchAbsolute will convert the exported function into a graylog flavored call
func (g *graylogAPI) searchAbsolute(query search.Definition, result *search.Result) error {

	searchDefinition := graylog.Search{
		Query: query.Terms,
		Range: "",
		From:  query.From,
		To:    query.To,
		Sort:  "timestamp:desc",
	}

	if query.Stream != "" {
		searchDefinition.Filter = fmt.Sprintf("streams:%s", query.Stream)
	}

	graylogResult, err := g.API.SearchAbsolute(searchDefinition)

	if err != nil {
		return err
	}

	// convert messages into messages
	for _, message := range graylogResult.Messages {

		//jsonStr, err := json.Marshal(message)

		if err != nil {
			return err
		}

		//log.InfoWith("debug", log.Data("test", jsonStr))

		messageDetails := message.(map[string]interface{})

		result.Messages = append(result.Messages, messageDetails)
	}

	// add all the others fields into result
	result.From = graylogResult.From
	result.To = graylogResult.To
	result.Query = graylogResult.Query
	result.Time = graylogResult.Time
	result.Fields = graylogResult.Fields
	result.TotalResults = strconv.Itoa(graylogResult.TotalResults)

	return nil

}
