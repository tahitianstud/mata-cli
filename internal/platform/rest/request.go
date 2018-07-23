package rest

import (
	"net/http"
)

// Request represents the data needed to pass to methods
type Request struct {
	URL string
	Headers http.Header
	Payload interface{}
	Username string
	Password string
}