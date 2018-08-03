// Package graylog holds the data structure and implementation of the
// Graylog API that can be interfaced with.
package graylog

import (
	"os"
	"fmt"
	"github.com/tahitianstud/mata-cli/internal/platform/rest"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
)

// API describe the api connection to Graylog
type API struct {
	Session
}

type Session struct {
	ConnectionString string
	Ticket           LoginTicket
}

// login will login to the specified server using the api
func (api *API) Login(connectionString string) error {

	api.ConnectionString = connectionString

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	var serverDefinition Server

	err = ParseConnectionString(&serverDefinition, api.ConnectionString)
	if err != nil {
		return err
	}

	validLoginTicket := LoginTicket{}
	loginPayload := loginParams{
		Username: serverDefinition.Username,
		Password: serverDefinition.Password,
		Host:     hostname,
	}
	loginOperation := fmt.Sprintf("%s%s", serverDefinition.GetURL(), loginAPI)

	loginRequest := rest.Request{
		URL:      loginOperation,
		Payload:  loginPayload,
		Username: serverDefinition.Username,
		Password: serverDefinition.Password,
	}
	response, err := rest.Post(loginRequest, &validLoginTicket)

	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("login operation returned %s", response.Status)
	}

	log.DebugWith("login response from Graylog server",
		log.Data("response", response))

	// save successful login ticket inside client
	api.Ticket = validLoginTicket

	return nil
}

// fetchSession will return a string representing the Session
func (api *API) FetchSession() Session {

	return api.Session

}

// restoreSession restores a Session
func (api *API) RestoreSession(session Session) {

	api.Session = session

}

// ListEnabledStreams will fetch the list of enabled streams from the GraylogServer
func (api *API) ListStreams() (streamsList StreamsList, err error) {

	listOfStreams := StreamsList{}

	var serverDefinition Server

	err = ParseConnectionString(&serverDefinition, api.ConnectionString)
	if err != nil {
		return streamsList, err
	}

	// use ticket if already logged in
	if api.Ticket.SessionID != "" {
		serverDefinition.Username = api.Ticket.SessionID
		serverDefinition.Password = "Session"
	}

	enabledStreamsOperation := fmt.Sprintf("%s%s", serverDefinition.GetURL(), enabledStreamsAPI)

	streamsListRequest := rest.Request{
		URL:      enabledStreamsOperation,
		Payload:  nil,
		Username: serverDefinition.Username,
		Password: serverDefinition.Password,
	}

	response, err := rest.Get(streamsListRequest, &listOfStreams)

	if err != nil {
		return listOfStreams, err
	}

	if response.StatusCode >= 400 {
		return listOfStreams, fmt.Errorf("streams listing operation returned %s", response.Status)
	}

	log.DebugWith("List of streams response from Graylog server",
		log.Data("response", listOfStreams))

	return listOfStreams, nil
}

// search will call the graylog API to make a search
func (api *API) Search(query Search) (result SearchResult, err error) {

	var serverDefinition Server

	err = ParseConnectionString(&serverDefinition, api.ConnectionString)
	if err != nil {
		return result, err
	}

	// use ticket if already logged in
	if api.Ticket.SessionID != "" {
		serverDefinition.Username = api.Ticket.SessionID
		serverDefinition.Password = "Session"
	}

	searchOperation := fmt.Sprintf("%s%s", serverDefinition.GetURL(), relativeSearchUniversalAPI)

	searchRequest := rest.Request{
		URL:      searchOperation,
		Payload:  query,
		Username: serverDefinition.Username,
		Password: serverDefinition.Password,
	}

	response, err := rest.Get(searchRequest, &result)

	if err != nil {
		return result, err
	}

	if response.StatusCode >= 400 {
		return result, fmt.Errorf("search operation returned %s", response.Status)
	}

	log.DebugWith("search response from Graylog server",
		log.Data("response", result))

	return result, nil

}

// search will call the graylog API to make a search
func (api *API) SearchAbsolute(query Search) (result SearchResult, err error) {

	var serverDefinition Server

	err = ParseConnectionString(&serverDefinition, api.ConnectionString)
	if err != nil {
		return result, err
	}

	// use ticket if already logged in
	if api.Ticket.SessionID != "" {
		serverDefinition.Username = api.Ticket.SessionID
		serverDefinition.Password = "Session"
	}

	searchOperation := fmt.Sprintf("%s%s", serverDefinition.GetURL(), absoluteSearchUniversalAPI)

	searchRequest := rest.Request{
		URL:      searchOperation,
		Payload:  query,
		Username: serverDefinition.Username,
		Password: serverDefinition.Password,
	}

	response, err := rest.Get(searchRequest, &result)

	if err != nil {
		return result, err
	}

	if response.StatusCode >= 400 {
		return result, fmt.Errorf("search operation returned %s", response.Status)
	}

	log.DebugWith("search response from Graylog server",
		log.Data("response", result))

	return result, nil

}
