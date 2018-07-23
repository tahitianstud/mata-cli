package graylog

import (
	"fmt"
	"regexp"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
	"strings"
	"github.com/tahitianstud/mata-cli/internal/platform/errors"
)

// Server describe the data model for a Graylog Server
type Server struct {
	Scheme string `description:"the Graylog server connection scheme" type:"select" choice:"http|https"`

	Host string `description:"the Graylog server hostname"`

	Port string `description:"the Graylog server port"`

	Endpoint string `description:"the Graylog server GraylogAPI endpoint"`

	Username string `description:"the Graylog username"`

	Password string `description:"the Graylog password" type:"password"`
}

// Defaults return a Server struct with sensible default values
func Defaults() *Server {
	return &Server{
		Host:     "",
		Port:     "9000",
		Endpoint: "/api",
		Username: "",
		Password: "",
		Scheme:   "http",
	}
}

func ParseConnectionString(s *Server, connectionString string) error {

	if len(connectionString) <= 0 {
		return errors.New("connnection string cannot be blank")
	}

	var validPattern = regexp.MustCompile(`(?m)^(http|https)://([a-z0-9\.]+)(:[0-9]+)?(/[0-9a-z]+)(\?u=([0-9a-zA-Z]+)(\&\&p=(.*))?)?$`)

	matches := validPattern.FindAllStringSubmatch(connectionString, -1)

	if matches == nil {
		return errors.New("connnection string is not valid for the API")
	}

	s.Scheme = matches[0][1]
	s.Host = matches[0][2]
	s.Endpoint = matches[0][4]

	if s.Scheme == "http" && matches[0][3] == "" {
		s.Port = "80"
	} else if s.Scheme == "https" && matches[0][3] == "" {
		s.Port = "443"
	} else {
		s.Port = strings.Replace(matches[0][3], ":", "", -1)
	}

	s.Username = matches[0][6]
	s.Password = matches[0][8]

	return nil
}

// URL will return the assembled URL for the given server definition
func (s *Server) GetURL() (URL string) {

	if (s.Scheme == "https" && s.Port == "443") || (s.Scheme == "http" && s.Port == "80") {

		URL = fmt.Sprintf("%s://%s%s", s.Scheme, s.Host, s.Endpoint)

	} else {

		URL = fmt.Sprintf("%s://%s:%s%s", s.Scheme, s.Host, s.Port, s.Endpoint)

	}

	return URL

}

// UpdateWithURL will update the Server definition using an URL
func (s *Server) UpdateWithURL(URL string) (err error) {

	// i.e. https://graylog.test.com:3949/api

	if len(URL) <= 0 {
		return fmt.Errorf("the URL cannot be blank")
	}

	// validate the URL first
	var validURLPattern = regexp.MustCompile(`(?m)^(http|https)://([a-z0-9\.]+)(:[0-9]+)?(/[0-9a-z]+)$`)
	matches := validURLPattern.FindAllStringSubmatch(URL, -1)

	if matches == nil {
		return fmt.Errorf("the URL '%s' is not valid", URL)
	}

	log.DebugWith(
		"URL matches with regular expression",
		log.Data("matches", matches))

	s.Scheme = matches[0][1]
	s.Host = matches[0][2]
	s.Endpoint = matches[0][4]

	if s.Scheme == "http" && matches[0][3] == "" {
		s.Port = "80"
	} else if s.Scheme == "https" && matches[0][3] == "" {
		s.Port = "443"
	} else {
		s.Port = strings.Replace(matches[0][3], ":", "", -1)
	}

	return nil
}

func (s *Server) GetConnectionString() string {
	connectionString := fmt.Sprintf("%s?u=%s&&p=%s", s.GetURL(), s.Username, s.Password)

	return connectionString
}
