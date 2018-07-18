package mata

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/dghubble/sling"
	"github.com/sirupsen/logrus"
)

// GraylogAPI describe the api connection to Graylog
type GraylogAPI struct {
	ServerConnection Server
	Ticket           LoginTicket
	SessionString    string
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Connect will return the necessary api connection to Graylog
func Connect(server Server) *GraylogAPI {

	api := &GraylogAPI{
		ServerConnection: server,
	}

	// TODO: check if non-expired session file already exists and skip if so

	loginTicket, err := api.Login(server)
	if err != nil {
		logrus.
			WithError(err).
			Fatalf("could not login to Graylog API")
	}

	logrus.
		WithField("sessionID", loginTicket.SessionID).
		Debugln("received valid loginTicket")

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

const (
	loginAPI = "/system/sessions"
)

// InvalidLoginTicket represents a ticket from un unsuccessful login attempt
var InvalidLoginTicket = LoginTicket{
	SessionID:  "N/A",
	ValidUntil: "N/A",
}

// LoginTicket describes the login ticket from a successful login call
type LoginTicket struct {
	SessionID  string `json:"session_id"`
	ValidUntil string `json:"valid_until"`
}

// LoginParams describes the params to send to the login API
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

// Login will connect the api to the specified Server
func (api *GraylogAPI) Login(server Server) (LoginTicket, error) {

	api.ServerConnection = server

	hostname, err := os.Hostname()
	if err != nil {
		return InvalidLoginTicket, err
	}

	validLoginTicket := LoginTicket{}
	loginParams := LoginParams{
		Username: api.ServerConnection.Username,
		Password: api.ServerConnection.Password,
		Host:     hostname,
	}
	loginOperation := fmt.Sprintf("%s%s", api.ServerConnection.GetURL(), loginAPI)
	response, err := sling.New().
		Post(loginOperation).
		BodyJSON(loginParams).
		ReceiveSuccess(&validLoginTicket)
	if err != nil {
		return InvalidLoginTicket, err
	} else if response.StatusCode >= 400 {
		return InvalidLoginTicket, fmt.Errorf("Login operation returned %s", response.Status)
	}

	logrus.
		WithField("response", response).
		Debugln("Login response from Graylog server")

	return validLoginTicket, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// StreamsList describes the result of the listing of streams
type StreamsList struct {
	NumberOfStreams int      `json:"total"`
	Streams         []Stream `json:"streams"`
}

// Stream describes a Graylog stream
type Stream struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

const (
	streamsAPI        = "/streams"
	enabledStreamsAPI = "/streams/enabled"
)

// ListEnabledStreams will fetch the list of enabled streams from the GraylogServer
func (api *GraylogAPI) ListEnabledStreams() (list StreamsList, err error) {

	listOfStreams := StreamsList{}

	enabledStreamsOperation := fmt.Sprintf("%s%s", api.ServerConnection.GetURL(), enabledStreamsAPI)
	response, err := sling.New().
		Set("Authorization", api.SessionString).
		Get(enabledStreamsOperation).
		ReceiveSuccess(&listOfStreams)
	if err != nil {
		return StreamsList{}, err
	} else if response.StatusCode >= 400 {
		return StreamsList{}, fmt.Errorf("Streams listing operation returned %s", response.Status)
	}

	logrus.
		WithField("response", response).
		Debugln("List of streams response from Graylog server")

	return listOfStreams, nil
}
