package server_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/server"
)


func TestDefaults(t *testing.T) {

	Convey("Given a Server model created with the constructor", t, func() {
		testServer := server.Defaults()

		Convey("the Default Values should be correct", func() {
			So(testServer.Host, ShouldEqual, "server.domain.com")
			So(testServer.Port, ShouldEqual, "9000")
			So(testServer.Scheme, ShouldEqual, "http")
			So(testServer.Endpoint, ShouldEqual, "/api")
			So(testServer.Username, ShouldEqual, "user")
			So(testServer.Password, ShouldEqual, "pass")
		})
	})

}
