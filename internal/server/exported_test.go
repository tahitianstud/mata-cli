package server_test

import (
	"github.com/tahitianstud/mata-cli/internal/server"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpdate(t *testing.T) {

	Convey("Given a Server model created with the constructor", t, func() {
		testServer := server.Defaults()

		Convey("updating a connection using a given correct URL", func() {
			err := server.UpdateConn(&testServer.Connection, "https://mata.server.com:4393/something")
			So(err, ShouldBeNil)

			Convey("the results are correctly updated", func() {
				So(testServer.Host, ShouldEqual, "mata.server.com")
				So(testServer.Port, ShouldEqual, "4393")
				So(testServer.Scheme, ShouldEqual, "https")
				So(testServer.Endpoint, ShouldEqual, "/something")
			})
		})

		Convey("updating using a given correct URL on default HTTPS port 443", func() {
			err := server.UpdateConn(&testServer.Connection, "https://mata.server.com/something")
			So(err, ShouldBeNil)

			Convey("the results are correctly updated", func() {
				So(testServer.Host, ShouldEqual, "mata.server.com")
				So(testServer.Port, ShouldEqual, "443")
				So(testServer.Scheme, ShouldEqual, "https")
				So(testServer.Endpoint, ShouldEqual, "/something")
			})
		})

		Convey("updating using a given correct URL on default HTTP port 80", func() {
			err := server.UpdateConn(&testServer.Connection, "http://mata.server.com/something")
			So(err, ShouldBeNil)

			Convey("the results are correctly updated", func() {
				So(testServer.Host, ShouldEqual, "mata.server.com")
				So(testServer.Port, ShouldEqual, "80")
				So(testServer.Scheme, ShouldEqual, "http")
				So(testServer.Endpoint, ShouldEqual, "/something")
			})
		})

		Convey("updating using a given incorrect URL", func() {
			err := server.UpdateConn(&testServer.Connection, "ftp://mata.server.com:4393/something")

			Convey("and error should be produced", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("the connectionstring should be correct", func() {

			testConnectionString := server.ConnectionString(testServer)

			So(testConnectionString, ShouldEqual, "http://server.domain.com:9000/api?u=user&&p=pass")

		})

		Convey("a blank URL cannot be used to update the server definition", func() {
			err := server.UpdateConn(&testServer.Connection, "")

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "the URL cannot be blank")
		})

	})
}

func TestURL(t *testing.T) {

	Convey("Given a test HTTP Server model", t, func() {
		testServer := server.Defaults()
		testServer.Host = "graylog.test.com"

		Convey("in HTTP", func() {
			testServer.Scheme = "http"

			Convey("with non standard port", func() {
				testServer.Port = "9001"

				Convey("with provided endpoint", func() {
					testServer.Endpoint = "/endpoint"

					Convey("the URL should be correct", func() {
						So(server.URL(testServer.Connection), ShouldEqual, "http://graylog.test.com:9001/endpoint")
					})
				})

				Convey("with empty endpoint", func() {
					testServer.Endpoint = ""

					Convey("the URL should be correct", func() {
						So(server.URL(testServer.Connection), ShouldEqual, "http://graylog.test.com:9001")
					})
				})

				Convey("with default endpoint", func() {

					Convey("the URL should be correct", func() {
						So(server.URL(testServer.Connection), ShouldEqual, "http://graylog.test.com:9001/api")
					})
				})
			})

			Convey("with standard port", func() {
				testServer.Port = "80"

				Convey("with provided endpoint", func() {
					testServer.Endpoint = "/endpoint"

					Convey("the URL should be correct", func() {
						So(server.URL(testServer.Connection), ShouldEqual, "http://graylog.test.com/endpoint")
					})
				})

				Convey("with empty endpoint", func() {
					testServer.Endpoint = ""

					Convey("the URL should be correct", func() {
						So(server.URL(testServer.Connection), ShouldEqual, "http://graylog.test.com")
					})
				})

				Convey("with default endpoint", func() {

					Convey("the URL should be correct", func() {
						So(server.URL(testServer.Connection), ShouldEqual, "http://graylog.test.com/api")
					})
				})
			})

		})

		Convey("in HTTPS", func() {
			testServer.Scheme = "https"

			Convey("with non standard port", func() {
				testServer.Port = "9001"

				Convey("with provided endpoint", func() {
					testServer.Endpoint = "/endpoint"

					Convey("the URL should be correct", func() {
						So(server.URL(testServer.Connection), ShouldEqual, "https://graylog.test.com:9001/endpoint")
					})
				})

				Convey("with empty endpoint", func() {
					testServer.Endpoint = ""

					Convey("the URL should be correct", func() {
						So(server.URL(testServer.Connection), ShouldEqual, "https://graylog.test.com:9001")
					})
				})

				Convey("with default endpoint", func() {

					Convey("the URL should be correct", func() {
						So(server.URL(testServer.Connection), ShouldEqual, "https://graylog.test.com:9001/api")
					})
				})
			})

			Convey("with standard port", func() {
				testServer.Port = "443"

				Convey("with provided endpoint", func() {
					testServer.Endpoint = "/endpoint"

					Convey("the URL should be correct", func() {
						So(server.URL(testServer.Connection), ShouldEqual, "https://graylog.test.com/endpoint")
					})
				})

				Convey("with empty endpoint", func() {
					testServer.Endpoint = ""

					Convey("the URL should be correct", func() {
						So(server.URL(testServer.Connection), ShouldEqual, "https://graylog.test.com")
					})
				})

				Convey("with default endpoint", func() {

					Convey("the URL should be correct", func() {
						So(server.URL(testServer.Connection), ShouldEqual, "https://graylog.test.com/api")
					})
				})
			})

		})
	})

}

func TestConnectionString(t *testing.T) {

	Convey("Given test definition with default values", t, func() {

		testDefinition := server.Defaults()

		Convey("the connection string should be correct", func() {

			testConnectionString := server.ConnectionString(testDefinition)

			So(testConnectionString, ShouldEqual, "http://server.domain.com:9000/api?u=user&&p=pass")

		})

	})

}

func TestUpdateWithConnectionString(t *testing.T) {

	Convey("Given test definition with default values", t, func() {

		testDefinition := server.Defaults()

		Convey("and a correct connection string", func() {

			connectionString := "https://test.server.com:4833/api?u=user&&p=pass"

			Convey("calling UpdateDef should update correctly", func() {

				err := server.UpdateDef(&testDefinition, connectionString)

				So(err, ShouldBeNil)
				So(server.URL(testDefinition.Connection), ShouldEqual, "https://test.server.com:4833/api")
				So(testDefinition.User.Username, ShouldEqual, "user")
				So(testDefinition.User.Password, ShouldEqual, "pass")

			})
		})

		Convey("and a correct HTTP connection string without port", func() {

			connectionString := "http://test.server.com/api?u=user&&p=pass"

			Convey("calling UpdateDef should update correctly", func() {

				err := server.UpdateDef(&testDefinition, connectionString)

				So(err, ShouldBeNil)
				So(server.URL(testDefinition.Connection), ShouldEqual, "http://test.server.com/api")
				So(testDefinition.Port, ShouldEqual, "80")
				So(testDefinition.User.Username, ShouldEqual, "user")
				So(testDefinition.User.Password, ShouldEqual, "pass")

			})
		})

		Convey("and a correct HTTPs connection string without port", func() {

			connectionString := "https://test.server.com/api?u=user&&p=pass"

			Convey("calling UpdateDef should update correctly", func() {

				err := server.UpdateDef(&testDefinition, connectionString)

				So(err, ShouldBeNil)
				So(server.URL(testDefinition.Connection), ShouldEqual, "https://test.server.com/api")
				So(testDefinition.Port, ShouldEqual, "443")
				So(testDefinition.User.Username, ShouldEqual, "user")
				So(testDefinition.User.Password, ShouldEqual, "pass")

			})
		})

		Convey("and an incorrect connection string", func() {

			connectionString := "ftps://test.server.com:4833/api?u=user&&p=pass"

			Convey("calling UpdateDef should return an error", func() {

				err := server.UpdateDef(&testDefinition, connectionString)

				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "error while parsing connection string")

			})
		})


	})

}
