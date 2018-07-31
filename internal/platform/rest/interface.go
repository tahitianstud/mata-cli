package rest

import (
	"net/http"
)

// Handler defines the operations necessary for a Rest Client
type Handler interface {
	Get(request Request, target interface{}) (response *http.Response, err error)

	Post(request Request, target interface{}) (response *http.Response, err error)
}