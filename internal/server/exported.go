package server

import (
	"regexp"
	"strings"
	"fmt"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
	"github.com/tahitianstud/mata-cli/internal/platform/errors"
)

// URL derives the URL to use to contact the server
// from a Connection struct
func URL(c Connection) string {

	var URL = ""

	if (c.Scheme == "https" && c.Port == "443") || (c.Scheme == "http" && c.Port == "80") {

		URL = fmt.Sprintf("%s://%s%s", c.Scheme, c.Host, c.Endpoint)

	} else {

		URL = fmt.Sprintf("%s://%s:%s%s", c.Scheme, c.Host, c.Port, c.Endpoint)

	}

	return URL
}

// UpdateConn will update the server connection information using an URL
func UpdateConn(c *Connection, URL string) error {

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

	c.Scheme = matches[0][1]
	c.Host = matches[0][2]
	c.Endpoint = matches[0][4]

	if c.Scheme == "http" && matches[0][3] == "" {

		c.Port = "80"

	} else if c.Scheme == "https" && matches[0][3] == "" {

		c.Port = "443"

	} else {

		c.Port = strings.Replace(matches[0][3], ":", "", -1)

	}

	return nil
}

// ConnectionString returns the connection string
// to use to connect to the server
func ConnectionString(d Definition) string {

	connectionString := fmt.Sprintf("%s?u=%s&&p=%s", URL(d.Connection), d.Username, d.Password)

	return connectionString

}

// UpdateDef updates the definition using a connection string
func UpdateDef(d *Definition, connectionString string) error {

	var validPattern = regexp.MustCompile(`(?m)^(http|https)://([a-z0-9\.]+)(:[0-9]+)?(/[0-9a-z]+)\?u=(.*)\&\&p=(.*)$`)

	matches := validPattern.FindAllStringSubmatch(connectionString, -1)

	if matches == nil {
		return errors.New("error while parsing connection string")
	}

	d.Scheme = matches[0][1]
	d.Host = matches[0][2]
	d.Endpoint = matches[0][4]

	if d.Scheme == "http" && matches[0][3] == "" {
		d.Port = "80"
	} else if d.Scheme == "https" && matches[0][3] == "" {
		d.Port = "443"
	} else {
		d.Port = strings.Replace(matches[0][3], ":", "", -1)
	}

	d.Username = matches[0][5]
	d.Password = matches[0][6]

	return nil

}
