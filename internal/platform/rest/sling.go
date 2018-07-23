package rest

import (
	"github.com/tahitianstud/mata-cli/internal/platform/rest/sling"
	"net/http"
)

type SlingClient struct {
	*sling.Sling
}

func GetSlingHandler() *SlingClient {
	return &SlingClient{
		Sling: &sling.Sling{},
	}
}

func (s *SlingClient) Get(request Request, target interface{}) (response *http.Response, err error) {
	// convert request to sling format (possible only if same field names and types)
	slingRequest := Convert(request)

	return s.Sling.Get(slingRequest).Update(target)
}

func (s *SlingClient) Post(request Request, target interface{}) (response *http.Response, err error) {
	// convert request to sling format (possible only if same field names and types)
	slingRequest := Convert(request)

	return s.Sling.Post(slingRequest).Update(target)
}

// Convert will convert between the definition of requests
func Convert(request Request) sling.Request {
	return sling.Request(request)
}

