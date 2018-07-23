package graylog

import (
	//"encoding/base64"
	"fmt"
	"os"

	"github.com/dghubble/sling"
	"encoding/base64"
	"encoding/json"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
)

// API describe the api connection to Graylog
type API struct {
	ConnectionString string
	Ticket           LoginTicket
	SessionString    string
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Connect will return the necessary api connection to Graylog
func Connect(connectionString string) *API {

	api := &API{
		ConnectionString: connectionString,
	}

	// TODO: check if non-expired session file already exists and skip if so

	loginTicketString, err := api.Login(connectionString)
	if err != nil {
		log.ErrorWith("could not login to Graylog API",
			log.Data("error", err))
	}

	loginTicket := DeSerializeLogin(loginTicketString)

	log.DebugWith("received valid loginTicket",
		log.Data("sessionID", loginTicket.SessionID))

	// TODO: write loginTicket to config file

	// store ticket inside API instance
	api.Ticket = loginTicket
	ticketSessionID := loginTicket.SessionID
	if ticketSessionID != "" {
		api.SessionString = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", ticketSessionID, "session"))))
	}

	return api
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Login will connect the api to the specified Server
func (api API) Login(connectionString string) (string, error) {

	api.ConnectionString = connectionString

	hostname, err := os.Hostname()
	if err != nil {
		return InvalidLoginTicket.Serialized(), err
	}

	serverDefinition, err := parseConnectionString(connectionString)
	if err != nil {
		return InvalidLoginTicket.Serialized(), err
	}

	validLoginTicket := LoginTicket{}
	loginParams := loginParams{
		Username: serverDefinition.Username,
		Password: serverDefinition.Password,
		Host:     hostname,
	}
	loginOperation := fmt.Sprintf("%s%s", serverDefinition.GetURL(), loginAPI)
	response, err := sling.New().
		Post(loginOperation).
		BodyJSON(loginParams).
		ReceiveSuccess(&validLoginTicket)
	if err != nil {
		return InvalidLoginTicket.Serialized(), err
	} else if response.StatusCode >= 400 {
		return InvalidLoginTicket.Serialized(), fmt.Errorf("login operation returned %s", response.Status)
	}

	log.DebugWith("Login response from Graylog server",
		log.Data("response", response))

	return validLoginTicket.Serialized(), nil
}

// ListEnabledStreams will fetch the list of enabled streams from the GraylogServer
func (api API) ListEnabledStreams() (list string, err error) {

	listOfStreams := StreamsList{}

	serverDefinition, err := parseConnectionString(api.ConnectionString)
	if err != nil {
		return "", err
	}

	enabledStreamsOperation := fmt.Sprintf("%s%s", serverDefinition.GetURL(), enabledStreamsAPI)
	response, err := sling.New().
		Set("Authorization", api.SessionString).
		Get(enabledStreamsOperation).
		ReceiveSuccess(&listOfStreams)
	if err != nil {
		return "", err
	} else if response.StatusCode >= 400 {

		return "", fmt.Errorf("streams listing operation returned %s", response.Status)
	}


	jsonListOfStreams, err := json.Marshal(listOfStreams)

	if err != nil {
		return "", err
	}

	log.DebugWith("List of streams response from Graylog server",
		log.Data("response", jsonListOfStreams))


	return string(jsonListOfStreams[:]),nil
}
