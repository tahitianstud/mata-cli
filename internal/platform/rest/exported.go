package rest

import (
	"net/http"
)

var Client Handler = GetSlingHandler()

func Get(request Request, target interface{}) (response *http.Response, err error) {
	return Client.Get(request, target)
}

func Post(request Request, target interface{}) (response *http.Response, err error) {
	return Client.Post(request, target)
}
