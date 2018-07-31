package stubapi
//
//import (
//	"fmt"
//	"github.com/tahitianstud/mata-cli/internal/platform/rest"
//	"github.com/tahitianstud/mata-cli/internal/platform/log"
//)
//
//// StubAPI describe the stub api connection
//type StubAPI struct {
//	ConnectionString string
//	Ticket           LoginTicket
//}
//
//// login will login to the specified server using the api
//func (api *StubAPI) Login() error {
//
//	// save successful login ticket inside client
//	api.Ticket = validLoginTicket
//
//	return nil
//}
//
//// fetchSession will return a string representing the session
//func (api *StubAPI) FetchSession() string {
//	return api.Ticket.Serialized()
//}
//
//// restoreSession will a serialized ticket and restore it into the api session
//func (api *StubAPI) RestoreSession(session string) error {
//	return api.Ticket.DeSerialize(session)
//}
//
//// ListEnabledStreams will fetch the list of enabled streams from the GraylogServer
//func (api *StubAPI) ListStreams() (streamsList StreamsList, err error) {
//
//	listOfStreams := StreamsList{}
//
//	serverDefinition, err := ParseConnectionString(api.ConnectionString)
//	if err != nil {
//		return streamsList, err
//	}
//
//	// use ticket if already logged in
//	if api.Ticket.SessionID != "" {
//		serverDefinition.Username = api.Ticket.SessionID
//		serverDefinition.Password = "session"
//	}
//
//	enabledStreamsOperation := fmt.Sprintf("%s%s", serverDefinition.GetURL(), enabledStreamsAPI)
//
//	streamsListRequest := rest.Request{
//		URL:      enabledStreamsOperation,
//		Payload:  nil,
//		Username: serverDefinition.Username,
//		Password: serverDefinition.Password,
//	}
//
//	response, err := rest.Get(streamsListRequest, &listOfStreams)
//
//	if err != nil {
//		return listOfStreams, err
//	}
//
//	if response.StatusCode >= 400 {
//		return listOfStreams, fmt.Errorf("streams listing operation returned %s", response.Status)
//	}
//
//	log.DebugWith("List of streams response from Graylog server",
//		log.Data("response", listOfStreams))
//
//	return listOfStreams, nil
//}
//
//// search will call the graylog API to make a search
//func (api *StubAPI) Search(search Search) (result SearchResult, err error) {
//
//	serverDefinition, err := ParseConnectionString(api.ConnectionString)
//	if err != nil {
//		return result, err
//	}
//
//	// use ticket if already logged in
//	if api.Ticket.SessionID != "" {
//		serverDefinition.Username = api.Ticket.SessionID
//		serverDefinition.Password = "session"
//	}
//
//	searchOperation := fmt.Sprintf("%s%s", serverDefinition.GetURL(), searchUniversalAPI)
//
//	searchRequest := rest.Request{
//		URL:      searchOperation,
//		Payload:  search,
//		Username: serverDefinition.Username,
//		Password: serverDefinition.Password,
//	}
//
//	response, err := rest.Get(searchRequest, &result)
//
//	if err != nil {
//		return result, err
//	}
//
//	if response.StatusCode >= 400 {
//		return result, fmt.Errorf("search operation returned %s", response.Status)
//	}
//
//	log.DebugWith("search response from Graylog server",
//		log.Data("response", result))
//
//	return result, nil
//
//}
