package graylog_test

import (
	"testing"
	"github.com/tahitianstud/mata-cli/internal/api/graylog"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/platform/rest"
	"github.com/tahitianstud/mata-cli/internal/platform/rest/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"github.com/tahitianstud/mata-cli/internal/platform/errors"
	"github.com/tahitianstud/mata-cli/internal/platform/json"
)

func TestAPI_Login(t *testing.T) {

	Convey("Given an API struct", t, func() {
		testApi := graylog.API{}

		mockHandler := mocks.Handler{}
		rest.Client = &mockHandler

		OKResponse := &http.Response{
			StatusCode: 200,
		}

		Convey("on positive return from post", func() {

			mockHandler.On("Post", mock.Anything, mock.Anything).Return(OKResponse, nil)

			Convey("trying to login with correct URL should work", func() {
				err := testApi.Login("http://server.test.com/api?u=user&&p=pass")

				So(err, ShouldBeNil)
			})

			Convey("trying to login with invalid URL should return error", func() {

				err := testApi.Login("ftp://server.test.com/api?p=pass")

				So(err, ShouldNotBeNil)
			})

		})

		Convey("on error returned from POST", func() {

			InternalServerError := &http.Response{
				StatusCode: 500,
			}

			mockHandler.On("Post", mock.Anything, mock.Anything).Return(InternalServerError, errors.New("my error"))

			Convey("trying to login with correct URL should return correct error", func() {
				err := testApi.Login("http://server.test.com/api?u=user&&p=pass")

				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "my error")
			})

		})

		Convey("on error 500 returned from POST", func() {

			InternalServerError := &http.Response{
				Status: "500 internal server error",
				StatusCode: 500,
			}

			mockHandler.On("Post", mock.Anything, mock.Anything).Return(InternalServerError, nil)

			Convey("trying to login with correct URL should return correct error", func() {
				err := testApi.Login("http://server.test.com/api?u=user&&p=pass")

				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "login operation returned 500 internal server error")
			})

		})

	})

}

func TestAPI_FetchSession(t *testing.T) {

	Convey("Given a filled out API with session", t, func() {

		testApi := graylog.API{

			Session: graylog.Session {

				ConnectionString: "http://server.test.com/api?u=user&&p=pass",

				Ticket: graylog.LoginTicket{
					SessionID: "1234567890",
					ValidUntil: "now",
				},

			},

		}

		Convey("fetching the session should be correct", func() {

			session := testApi.FetchSession()

			So(session.ConnectionString, ShouldEqual, `http://server.test.com/api?u=user&&p=pass`)
			So(session.Ticket.SessionID, ShouldEqual, `1234567890`)
			So(session.Ticket.ValidUntil, ShouldEqual, `now`)

		})

	})

}


func TestAPI_RestoreSession(t *testing.T) {

	Convey("Given a filled out API with session", t, func() {

		testApi := graylog.API{

			Session: graylog.Session {},

		}

		Convey("restoring the session should work", func() {

			var testSession graylog.Session
			json.DeSerializeString(`{"ConnectionString":"http://server.test.com/api?u=user&p=pass","Ticket":{"session_id":"1234567890","valid_until":"now"}}`, &testSession)

			testApi.RestoreSession(testSession)

			So(testApi.ConnectionString, ShouldEqual, "http://server.test.com/api?u=user&p=pass")
			So(testApi.Ticket.SessionID, ShouldEqual, "1234567890")
			So(testApi.Ticket.ValidUntil, ShouldEqual, "now")

		})

	})

}


func TestAPI_ListStreams(t *testing.T) {

	Convey("Given an API struct with mocked rest handler", t, func() {
		testApi := graylog.API{
			Session: graylog.Session {

				ConnectionString: "http://server.test.com/api?u=user&&p=pass",

			},
		}

		mockHandler := mocks.Handler{}
		rest.Client = &mockHandler

		OKResponse := &http.Response{
			StatusCode: 200,
		}

		mockHandler.On("Get", mock.Anything, mock.Anything).Return(OKResponse, nil)

		Convey("listing stream should be called correctly", func() {

			streamsList, err := testApi.ListStreams()

			So(err, ShouldBeNil)
			So(streamsList, ShouldNotBeNil)

		})
	})

	Convey("Given an API struct with mocked rest handler in error", t, func() {
		testApi := graylog.API{

			Session: graylog.Session {

				ConnectionString: "http://server.test.com/api?u=user&&p=pass",

			},

		}

		mockHandler := mocks.Handler{}
		rest.Client = &mockHandler

		BadResponse := &http.Response{
			StatusCode: 500,
		}

		mockHandler.On("Get", mock.Anything, mock.Anything).Return(BadResponse, errors.New("failed Get"))

		Convey("listing stream should return an error", func() {

			_, err := testApi.ListStreams()

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "failed Get")

		})
	})

	Convey("Given an API struct with mocked rest handler with bad response", t, func() {
		testApi := graylog.API{

			Session: graylog.Session {

				ConnectionString: "http://server.test.com/api?u=user&&p=pass",

			},
		}

		mockHandler := mocks.Handler{}
		rest.Client = &mockHandler

		BadResponse := &http.Response{
			StatusCode: 500,
			Status: "500",
		}

		mockHandler.On("Get", mock.Anything, mock.Anything).Return(BadResponse, nil)

		Convey("listing stream should return an error", func() {

			_, err := testApi.ListStreams()

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "streams listing operation returned 500")

		})
	})

	Convey("Given an API struct with mocked rest handler but bad connection string", t, func() {
		testApi := graylog.API{

			Session: graylog.Session {

				ConnectionString: "ftp://server.test.com/api?u=user&&p=pass",

			},

		}

		mockHandler := mocks.Handler{}
		rest.Client = &mockHandler

		OKResponse := &http.Response{
			StatusCode: 200,
		}

		mockHandler.On("Get", mock.Anything, mock.Anything).Return(OKResponse, nil)

		Convey("listing stream should return an error", func() {

			_, err := testApi.ListStreams()

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "connnection string is not valid for the API")

		})
	})

	Convey("Given an API struct with mocked rest handler and saved login ticket", t, func() {
		testApi := graylog.API{
			Session: graylog.Session {

				ConnectionString: "https://server.test.com/api?u=user&&p=pass",
				Ticket: graylog.LoginTicket{
					SessionID: "123456789",
					ValidUntil: "tomorrow",
				},

			},
		}

		mockHandler := mocks.Handler{}
		rest.Client = &mockHandler

		OKResponse := &http.Response{
			StatusCode: 200,
		}

		mockHandler.On("Get", mock.Anything, mock.Anything).Return(OKResponse, nil)

		Convey("listing stream should be called correctly", func() {

			streamsList, err := testApi.ListStreams()

			So(err, ShouldBeNil)
			So(streamsList, ShouldNotBeNil)

		})
	})

}


func TestAPI_Search(t *testing.T) {

	Convey("Given an API struct with mocked rest handler", t, func() {
		testApi := graylog.API{
			Session: graylog.Session {

				ConnectionString: "http://server.test.com/api?u=user&&p=pass",
				Ticket: graylog.LoginTicket{
					SessionID: "123456789",
					ValidUntil: "now",
				},

			},
		}

		mockHandler := mocks.Handler{}
		rest.Client = &mockHandler

		OKResponse := &http.Response{
			StatusCode: 200,
		}

		testQuery := graylog.Search{
			Query: "*",
		}

		mockHandler.On("Get", mock.Anything, mock.Anything).Return(OKResponse, nil)

		Convey("searching should be called correctly", func() {

			searchResult, err := testApi.Search(testQuery)

			So(err, ShouldBeNil)
			So(searchResult, ShouldNotBeNil)

		})
	})

	Convey("Given an API struct with mocked rest handler that fails", t, func() {
		testApi := graylog.API{
			Session: graylog.Session {

				ConnectionString: "http://server.test.com/api?u=user&&p=pass",
				Ticket: graylog.LoginTicket{
					SessionID: "123456789",
					ValidUntil: "now",
				},

			},
		}

		mockHandler := mocks.Handler{}
		rest.Client = &mockHandler

		BadResponse := &http.Response{
			StatusCode: 500,
			Status: "500",
		}

		testQuery := graylog.Search{
			Query: "*",
		}

		mockHandler.On("Get", mock.Anything, mock.Anything).Return(BadResponse, nil)

		Convey("searching should return an error", func() {

			_, err := testApi.Search(testQuery)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "search operation returned 500")

		})
	})

	Convey("Given an API struct with mocked rest handler that errors", t, func() {
		testApi := graylog.API{
			Session: graylog.Session {

				ConnectionString: "http://server.test.com/api?u=user&&p=pass",
				Ticket: graylog.LoginTicket{
					SessionID: "123456789",
					ValidUntil: "now",
				},

			},
		}

		mockHandler := mocks.Handler{}
		rest.Client = &mockHandler

		BadResponse := &http.Response{
			StatusCode: 500,
			Status: "500",
		}

		testQuery := graylog.Search{
			Query: "*",
		}

		mockHandler.On("Get", mock.Anything, mock.Anything).Return(BadResponse, errors.New("rest error"))

		Convey("searching should return an error", func() {

			_, err := testApi.Search(testQuery)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "rest error")

		})
	})

	Convey("Given an API struct with bad connection string with mocked rest handler", t, func() {
		testApi := graylog.API{
			Session: graylog.Session {

				ConnectionString: "ftp://server.test.com/api?u=user&&p=pass",

			},
		}

		mockHandler := mocks.Handler{}
		rest.Client = &mockHandler

		OKResponse := &http.Response{
			StatusCode: 200,
		}

		testQuery := graylog.Search{
			Query: "*",
		}

		mockHandler.On("Get", mock.Anything, mock.Anything).Return(OKResponse, nil)

		Convey("searching should be called correctly", func() {

			_, err := testApi.Search(testQuery)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "connnection string is not valid for the API")

		})
	})

}