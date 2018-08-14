package api

import (
	"github.com/pkg/errors"
	"github.com/tahitianstud/mata-cli/internal/platform/io/file"
	"github.com/tahitianstud/mata-cli/internal"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
	"github.com/tahitianstud/mata-cli/internal/server"
	"github.com/tahitianstud/mata-cli/internal/platform/json"
	"github.com/tahitianstud/mata-cli/internal/stream"
	"github.com/tahitianstud/mata-cli/internal/search"
	"strings"
)

// FetchProvider will return the right implementation for the api.provider interface
func FetchProvider(api supportedAPI) (client provider) {

	switch api {
	case GRAYLOG:

		return fetchGraylogAPI()
	}

	return nil
}


// Connect will use an authentifier to log in successfully but
// will return an error if authentication failed
func Connect(authentifier authentifier, connectionString string) error {

	// use connection string to define the server

	serverDefinition := server.Definition{}
	server.UpdateDef(&serverDefinition, connectionString)

	// check if non-expired session file already exists and if so, restore the session into the api provider

	var savedSession session
	savedSessionStr := ""
	err := file.ReadFrom(internal.ConfigLocation, ".session", &savedSessionStr)
	log.DebugIf(errors.Wrap(err, "could not read session from disk"))

	err = json.DeSerializeString(savedSessionStr, &savedSession)
	log.DebugIf(errors.Wrap(err, "could not deserialize session from disk"))

	if err == nil && len(savedSessionStr) > 0 {
		err = authentifier.restoreSession(savedSession)
		log.WarningIf(errors.Wrap(err, "could not restore saved user session"))
	}

	// if both username and password are provided, redo login even if session still active

	if serverDefinition.Username != "" && serverDefinition.Password != "" {

		log.InfoWith("connecting to server",
			log.Data("URL", server.URL(serverDefinition.Connection)),
			log.Data("user", serverDefinition.Username),
		)

		err = authentifier.login(connectionString)
		if err != nil {
			return err
		}

		// save session
		session := authentifier.fetchSession()

		// remove password from connectionstring before saving to disk
		session.ConnectionString = session.ConnectionString[:strings.LastIndex(session.ConnectionString, "&p=")-1]

		// serialize to json string for saving
		serializedSession := json.SerializeToString(session)

		// write to .session file
		err = file.WriteInside(internal.ConfigLocation, ".session", []byte(serializedSession), true)
		log.WarningIf(errors.Wrap(err, "could not save session to disk"))

		// if no session active and no credentials provided, then return an error

	} else if authentifier.sessionActive() == false && (serverDefinition.Username == "" || serverDefinition.Password == "") {

		return errors.Wrap(err, "you have to provide username/password in order to login (your session may have expired)")

		// if a session is active, log a message informing the user so

	} else if authentifier.sessionActive() == true {

		log.InfoWith("already connected to server",
			log.Data("URL", savedSession.ConnectionString),
		)

	}

	return nil

}


// SessionExists determines if a session is already saved on disk
func SessionExists() bool {
	content := ""

	err := file.ReadFrom(internal.ConfigLocation, ".session", &content)

	if err != nil {
		return false
	}

	return true
}


// DeserializeSession will extract loginTicket data from a string
func DeserializeSession(data string, ticket interface{}) error {

	err := json.DeSerializeString(data, &ticket)
	if err != nil {
		return err
	}

	return nil
}


// listStreams allows you to search for enabled streams in the API
func ListStreams(searcher searcher) (stream.StreamsList, error) {

	return searcher.listStreams()

}

// Search will execute a search using the API provider
func Search(searcher searcher, searchType search.Type, queryDefinition search.Definition, result *search.Result) error {

	switch searchType {

	case search.ABSOLUTE:
		return searcher.searchAbsolute(queryDefinition, result)

	case search.RELATIVE:
		return searcher.search(queryDefinition, result)

	default:
		return errors.New("invalid search type")
	}

}