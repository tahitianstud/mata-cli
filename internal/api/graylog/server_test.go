package graylog_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/api/graylog"
)

func TestParseConnectionString(t *testing.T) {

	Convey("Given a HTTP connection string with login", t, func() {
		testConnectionStr := "http://mata.domain.com:4383/api?u=user&&p=test"

		Convey("converting should return the right struct", func() {

			var testDefinition graylog.Server

			err := graylog.ParseConnectionString(&testDefinition, testConnectionStr)

			So(err, ShouldBeNil)
			So(testDefinition.GetURL(), ShouldEqual, "http://mata.domain.com:4383/api")


		})
	})

	Convey("Given a HTTP connection string with login", t, func() {
		testConnectionStr := "http://mata.domain.com/api?u=user&&p=test"

		Convey("converting should return the right struct", func() {

			var testDefinition graylog.Server

			err := graylog.ParseConnectionString(&testDefinition, testConnectionStr)

			So(err, ShouldBeNil)
			So(testDefinition.GetURL(), ShouldEqual, "http://mata.domain.com/api")


		})
	})

	Convey("Given a HTTPs connection string with login", t, func() {
		testConnectionStr := "https://mata.domain.com/api?u=user&&p=test"

		Convey("converting should return the right struct", func() {

			var testDefinition graylog.Server

			err := graylog.ParseConnectionString(&testDefinition, testConnectionStr)

			So(err, ShouldBeNil)
			So(testDefinition.GetURL(), ShouldEqual, "https://mata.domain.com/api")


		})
	})
}

func TestServer_UpdateWithURL(t *testing.T) {
	Convey("Given a correct HTTP connection string", t, func() {
		testConnectionStr := "http://mata.domain.com:4383/api"

		Convey("updating a test definition should return the right struct", func() {

			testServerDefinition := graylog.Defaults()

			err := testServerDefinition.UpdateWithURL(testConnectionStr)

			So(err, ShouldBeNil)
			So(testServerDefinition.GetURL(), ShouldEqual, "http://mata.domain.com:4383/api")

		})
	})

	Convey("Given a correct HTTP connection string w/o port", t, func() {
		testConnectionStr := "http://mata.domain.com/api"

		Convey("updating a test definition should return the right struct", func() {

			testServerDefinition := graylog.Defaults()

			err := testServerDefinition.UpdateWithURL(testConnectionStr)

			So(err, ShouldBeNil)
			So(testServerDefinition.GetURL(), ShouldEqual, "http://mata.domain.com/api")

		})
	})

	Convey("Given a correct HTTPsa connection string w/o port", t, func() {
		testConnectionStr := "https://mata.domain.com/api"

		Convey("updating a test definition should return the right struct", func() {

			testServerDefinition := graylog.Defaults()

			err := testServerDefinition.UpdateWithURL(testConnectionStr)

			So(err, ShouldBeNil)
			So(testServerDefinition.GetURL(), ShouldEqual, "https://mata.domain.com/api")

		})
	})

	Convey("Given an incorrect HTTP connection string", t, func() {

		testConnectionStr := "ftp://mata.domain.com:4383/api"

		Convey("updating a test definition should return the right struct", func() {

			testServerDefinition := graylog.Defaults()

			err := testServerDefinition.UpdateWithURL(testConnectionStr)

			So(err.Error(), ShouldEqual, "the URL 'ftp://mata.domain.com:4383/api' is not valid")

		})
	})

	Convey("Given a blank HTTP connection string", t, func() {

		Convey("updating a test definition should return the right struct", func() {

			testServerDefinition := graylog.Defaults()

			err := testServerDefinition.UpdateWithURL("")

			So(err.Error(), ShouldEqual, "the URL cannot be blank")

		})
	})
}

func TestServer_GetConnectionString(t *testing.T) {
	Convey("Given test server definition", t, func() {

		testServer := graylog.Server{
			Scheme: "https",
			Username: "test",
			Password: "test",
			Host: "app.domain.com",
			Port: "3493",
			Endpoint: "/api",
		}

		Convey("the corresponding ConnectionString should be correct", func() {

			connectionString := testServer.GetConnectionString()

			So(connectionString, ShouldEqual, "https://app.domain.com:3493/api?u=test&&p=test")

		})


	})

	Convey("Given test server definition w/o port", t, func() {

		testServer := graylog.Server{
			Scheme: "https",
			Username: "test",
			Password: "test",
			Host: "app.domain.com",
			Port: "443",
			Endpoint: "/api",
		}

		Convey("the corresponding ConnectionString should be correct", func() {

			connectionString := testServer.GetConnectionString()

			So(connectionString, ShouldEqual, "https://app.domain.com/api?u=test&&p=test")

		})


	})
}