// Package api holds the data structures and functions
// that can be used to interact with a logging server API.
//
// It also contains implementations to specific logging server like Graylog.
package api

import (
	"github.com/tahitianstud/mata-cli/internal/search"
	"github.com/tahitianstud/mata-cli/internal/stream"
)

// supportedAPI describes the types of API that are supported (for now Graylog only)
type supportedAPI int

const (
	GRAYLOG supportedAPI = iota
)

// provider defines what you should be able to do with an API implementation
type provider interface {
	authentifier

	searcher

}

type authentifier interface {

	// LOGIN-related API

	login(connectionString string) error

	fetchSession() session

	restoreSession(session session) error

	sessionActive() bool

}

type searcher interface {

	// SEARCH-related API

	listStreams() (stream.StreamsList, error)

	search(query search.Definition, result *search.Result) error

	searchAbsolute(query search.Definition, result *search.Result) error

}
