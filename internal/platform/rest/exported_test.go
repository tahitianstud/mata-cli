package rest_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/tahitianstud/mata-cli/internal/platform/rest"
	"net/http"
	"github.com/tahitianstud/mata-cli/internal/platform/rest/mocks"
	"github.com/stretchr/testify/mock"
)

func TestExportedFunctions(t *testing.T) {

	Convey("Given mocks of the exported interface", t, func() {

		OKResponse := &http.Response{
			StatusCode: 200,
		}

		Convey("Mocking the Get call", func() {

			mockHandler := mocks.Handler{}
			rest.Client = &mockHandler

			mockHandler.On("Get", mock.AnythingOfType("Request"), mock.Anything).Return(OKResponse, nil)

			Convey("with test values", func() {

				testRequest := rest.Request{
					URL: "http://test.server.com",
					Headers: http.Header{
						"test": []string{"testheader"},
					},
					Payload: "",
				}

				var target interface{}

				Convey("calling functions should be correctly routed", func() {

					response, err := rest.Get(testRequest, &target)
					So(response, ShouldEqual, OKResponse)
					So(err, ShouldBeNil)

					mockHandler.AssertExpectations(t)

				})
			})

		})

		Convey("Mocking the Post call", func() {

			mockHandler := mocks.Handler{}
			rest.Client = &mockHandler

			mockHandler.On("Post", mock.Anything, mock.Anything).Return(OKResponse, nil)

			Convey("with test values", func() {

				var target interface{}
				testRequest := rest.Request{
					URL: "http://test.server.com",
					Headers: http.Header{
						"test": []string{"testheader"},
					},
					Payload: "",
				}

				Convey("calling functions should be correctly routed", func() {

					response, err := rest.Post(testRequest, &target)
					So(response, ShouldEqual, OKResponse)
					So(err, ShouldBeNil)

					mockHandler.AssertExpectations(t)


				})
			})

		})

	})

}
