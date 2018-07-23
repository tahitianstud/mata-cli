package server

import (
	"fmt"
	"regexp"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
	"strings"
)

// Definition describes the data used to define a server
type Definition struct {
	Scheme string `description:"the server connection scheme" type:"select" choice:"http|https"`

	Host string `description:"the server hostname"`

	Port string `description:"the server port"`

	Endpoint string `description:"the server API endpoint"`

	Username string `description:"the username"`

	Password string `description:"the password" type:"password"`
}

// Defaults return an instance of Definition with sensible default values
func Defaults() Definition {
	return Definition{

		Host:     "",

		Port:     "9000",

		Endpoint: "/api",

		Username: "",

		Password: "",

		Scheme:   "http",

	}
}

// GetURL derives the URL to use to contact the server
// from the complete server definition
func (d *Definition) GetURL() string {

	var URL string

	if (d.Scheme == "https" && d.Port == "443") || (d.Scheme == "http" && d.Port == "80") {

		URL = fmt.Sprintf("%s://%s%s", d.Scheme, d.Host, d.Endpoint)

	} else {

		URL = fmt.Sprintf("%s://%s:%s%s", d.Scheme, d.Host, d.Port, d.Endpoint)

	}

	return URL

}

// SetURL will update the Server definition using an URL
func (d *Definition) SetURL(URL string) (err error) {

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

	d.Scheme = matches[0][1]
	d.Host = matches[0][2]
	d.Endpoint = matches[0][4]

	if d.Scheme == "http" && matches[0][3] == "" {

		d.Port = "80"

	} else if d.Scheme == "https" && matches[0][3] == "" {

		d.Port = "443"

	} else {

		d.Port = strings.Replace(matches[0][3], ":", "", -1)

		if err != nil {

			return fmt.Errorf("cannot parse the port inside the URL %s", URL)

		}

	}

	return nil
}

// GetConnectionString returns the connection string
// to use to connect to the server
func (d *Definition) GetConnectionString() string {

	connectionString := fmt.Sprintf("%s?u=%s&&p=%s", d.GetURL(), d.Username, d.Password)

	return connectionString

}
