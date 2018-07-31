package sling

import (
	"github.com/dghubble/sling"
	"net/http"
)

type Method int

const (
	GET Method = iota
	POST
	HEAD
	PUT
)

type Sling struct {
	Client Slinger
}

func (s *Sling) Update(target interface{}) (response *http.Response, err error) {

	response, err = s.Client.ReceiveSuccess(target)

	return response, err

}

func (s *Sling) Get(request Request) *Sling {

	var slingClient = sling.New()

	slingClient = slingClient.Get(request.URL)

	slingClient = addHeaders(request, slingClient)

	slingClient = addBasicAuth(request, slingClient)

	slingClient = addQuerystring(request, slingClient)

	s.Client = slingClient

	return s
}

func (s *Sling) Post(request Request) *Sling {

	var slingClient = sling.New()

	slingClient = slingClient.Post(request.URL)

	slingClient = addHeaders(request, slingClient)

	slingClient = addBasicAuth(request, slingClient)

	slingClient = addPayload(request, slingClient)

	s.Client = slingClient

	return s
}

func addPayload(request Request, slingClient *sling.Sling) *sling.Sling {
	if request.Payload != nil {
		slingClient = slingClient.BodyJSON(request.Payload)
	}
	return slingClient
}

func addQuerystring(request Request, slingClient *sling.Sling) *sling.Sling {
	if request.Payload != nil {
		slingClient = slingClient.QueryStruct(request.Payload)
	}
	return slingClient
}

func addHeaders(request Request, slingClient *sling.Sling) *sling.Sling {
	if request.Headers != nil {
		for name, values := range request.Headers {
			for _, value := range values {
				slingClient = slingClient.Set(name, value)
			}
		}
	}
	return slingClient
}

func addBasicAuth(request Request, slingClient *sling.Sling) *sling.Sling {
	// add basic authorization
	if request.Username != "" && request.Password != "" {
		slingClient = slingClient.SetBasicAuth(request.Username, request.Password)
	}
	return slingClient
}