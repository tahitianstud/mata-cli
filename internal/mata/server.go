package mata

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

// Server describe the data model for a Graylog Server
type Server struct {
	Scheme   string `description:"the Graylog server connection scheme" type:"select" choice:"http|https"`
	Host     string `description:"the Graylog server hostname"`
	Port     string `description:"the Graylog server port"`
	Endpoint string `description:"the Graylog server API endpoint"`
	Username string `description:"the Graylog username"`
	Password string `description:"the Graylog password" type:"password"`
}

// ServerService determines what you can do with Server structs
type ServerService interface {
	Validate(server Server) bool
}

// DefaultServer creates a Server definition with default values
func DefaultServer() Server {

	return Server{
		Host:     "",
		Port:     "9000",
		Endpoint: "/api",
		Username: "",
		Password: "",
		Scheme:   "http",
	}
}

// GetURL will return the assembled URL for the given server definition
func (s Server) GetURL() (URL string) {
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

	logrus.
		WithField("matches", matches).
		Debugln("URL matches with regular expression")

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

// UpdateWithUsername will update the server information with the username
func (s *Server) UpdateWithUsername(username string) error {
	if len(username) <= 0 {
		return fmt.Errorf("the username cannot be blank")
	}

	s.Username = username

	return nil
}

// UpdateWithPassword will update the server information with the password
func (s *Server) UpdateWithPassword(password string) error {
	if len(password) <= 0 {
		return fmt.Errorf("the password cannot be blank")
	}

	s.Password = password

	return nil
}
