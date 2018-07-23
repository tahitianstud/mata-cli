package sling

import (
	"net/http"
	"github.com/dghubble/sling"
)

type Slinger interface {
	Get(pathURL string) *sling.Sling
	Post(pathURL string) *sling.Sling
	Set(key, value string) *sling.Sling
	ReceiveSuccess(interface{}) (*http.Response, error)
	BodyJSON(bodyJSON interface{}) *sling.Sling
}