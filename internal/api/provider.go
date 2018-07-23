package api

import (
	"github.com/tahitianstud/mata-cli/internal/api/graylog"
)

// SupportedAPI describes the types of API that are supported (for now Graylog only)
type SupportedAPI int

const (
	GRAYLOG SupportedAPI = iota
)

// Connect returns a client API to the underlying server
func Connect(connectionString string, api ...SupportedAPI) APIProvider {

	if len(api) <= 0 {

		return graylog.Connect(connectionString)

	}

	return nil

}

// APIProvider defines what you can do with an API
type APIProvider interface {

	Login(connectionString string) (string, error)

	ListEnabledStreams() (string, error)

}
