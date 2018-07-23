package sling_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/platform/rest/sling"
	"net/http"
	"github.com/tahitianstud/mata-cli/internal/platform/rest/sling/mocks"
	"github.com/stretchr/testify/mock"
)

func TestSlingOperations(t *testing.T) {

	Convey("Given test data", t, func() {

		URL := "http://test.server.com/api"
		var headers http.Header
		payload := ""

		testRequest := sling.Request{
			URL: URL,
			Headers: headers,
			Payload: payload,
		}

		Convey("creating a sling client for a GET request", func() {

			testClient := &sling.Sling{}
			testClient = testClient.Get(testRequest)

			Convey("resulting client should be correct", func() {

				So(testClient, ShouldNotBeNil)

				So(testClient.Client, ShouldNotBeNil)
			})

		})

		Convey("creating a sling client for a POST request", func() {

			testClient := &sling.Sling{}
			testClient = testClient.Post(testRequest)

			Convey("resulting client should be correct", func() {

				So(testClient, ShouldNotBeNil)

				So(testClient.Client, ShouldNotBeNil)
			})

		})

		Convey("creating a sling client for a POST request with headers", func() {

			headers = http.Header{
				"test": []string{"testVal"},
			}
			testRequest.Headers = headers

			testClient := &sling.Sling{}
			testClient = testClient.Post(testRequest)

			Convey("resulting client should be correct", func() {

				So(testClient, ShouldNotBeNil)

				So(testClient.Client, ShouldNotBeNil)
			})

		})

		Convey("creating a sling client for a GET request with headers", func() {

			headers = http.Header{
				"test": []string{"testVal"},
			}
			testRequest.Headers = headers

			testClient := &sling.Sling{}
			testClient = testClient.Get(testRequest)

			Convey("resulting client should be correct", func() {

				So(testClient, ShouldNotBeNil)

				So(testClient.Client, ShouldNotBeNil)
			})

		})

		Convey("creating a sling client for a GET request with basic auth", func() {

			headers = http.Header{
				"test": []string{"testVal"},
			}
			testRequest.Headers = headers
			testRequest.Username = "user"
			testRequest.Password = "pass"

			testClient := &sling.Sling{}
			testClient = testClient.Get(testRequest)

			Convey("resulting client should be correct", func() {

				So(testClient, ShouldNotBeNil)

				So(testClient.Client, ShouldNotBeNil)

			})

		})

		Convey("creating a client with a mock", func() {

			var testTarget = ""

			testClient := &sling.Sling{}
			mockSlinger := &mocks.Slinger{}
			testClient.Client = mockSlinger

			mockSlinger.On("ReceiveSuccess", mock.Anything).Run(func(args mock.Arguments) {
				arg := args.Get(0).(*string)
				*arg = "updated value !"
			}).Return(&http.Response{StatusCode:200}, nil)

			Convey("asking for update will call the correct code", func() {

				response, err := testClient.Update(&testTarget)

				So(response.StatusCode, ShouldEqual, 200)
				So(err, ShouldBeNil)
				So(testTarget, ShouldEqual, "updated value !")

			})

		})



	})

}