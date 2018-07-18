package mata

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServerModel(t *testing.T) {
	Convey("Given a Server model created with the constructor", t, func() {
		testServer := DefaultServer()

		Convey("the Default Values should be correct", func() {
			So(testServer.Host, ShouldBeBlank)
			So(testServer.Port, ShouldEqual, "9000")
			So(testServer.Scheme, ShouldEqual, "http")
			So(testServer.Endpoint, ShouldEqual, "/api")
			So(testServer.Username, ShouldBeBlank)
			So(testServer.Password, ShouldBeBlank)
		})

		Convey("updating using a given correct URL", func() {
			err := testServer.UpdateWithURL("https://mata.server.com:4393/something")
			So(err, ShouldBeNil)

			Convey("the results are correctly updated", func() {
				So(testServer.Host, ShouldEqual, "mata.server.com")
				So(testServer.Port, ShouldEqual, "4393")
				So(testServer.Scheme, ShouldEqual, "https")
				So(testServer.Endpoint, ShouldEqual, "/something")
			})
		})

		Convey("updating using a given correct URL on default HTTPS port 443", func() {
			err := testServer.UpdateWithURL("https://mata.server.com/something")
			So(err, ShouldBeNil)

			Convey("the results are correctly updated", func() {
				So(testServer.Host, ShouldEqual, "mata.server.com")
				So(testServer.Port, ShouldEqual, "443")
				So(testServer.Scheme, ShouldEqual, "https")
				So(testServer.Endpoint, ShouldEqual, "/something")
			})
		})

		Convey("updating using a given correct URL on default HTTP port 80", func() {
			err := testServer.UpdateWithURL("http://mata.server.com/something")
			So(err, ShouldBeNil)

			Convey("the results are correctly updated", func() {
				So(testServer.Host, ShouldEqual, "mata.server.com")
				So(testServer.Port, ShouldEqual, "80")
				So(testServer.Scheme, ShouldEqual, "http")
				So(testServer.Endpoint, ShouldEqual, "/something")
			})
		})

		Convey("updating using a given incorrect URL", func() {
			err := testServer.UpdateWithURL("ftp://mata.server.com:4393/something")

			Convey("and error should be produced", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a test HTTP Server model", t, func() {
		testServer := DefaultServer()
		testServer.Host = "graylog.test.com"

		Convey("in HTTP", func() {
			testServer.Scheme = "http"

			Convey("with non standard port", func() {
				testServer.Port = "9001"

				Convey("with provided endpoint", func() {
					testServer.Endpoint = "/endpoint"

					Convey("the URL should be correct", func() {
						So(testServer.GetURL(), ShouldEqual, "http://graylog.test.com:9001/endpoint")
					})
				})

				Convey("with empty endpoint", func() {
					testServer.Endpoint = ""

					Convey("the URL should be correct", func() {
						So(testServer.GetURL(), ShouldEqual, "http://graylog.test.com:9001")
					})
				})

				Convey("with default endpoint", func() {

					Convey("the URL should be correct", func() {
						So(testServer.GetURL(), ShouldEqual, "http://graylog.test.com:9001/api")
					})
				})
			})

			Convey("with standard port", func() {
				testServer.Port = "80"

				Convey("with provided endpoint", func() {
					testServer.Endpoint = "/endpoint"

					Convey("the URL should be correct", func() {
						So(testServer.GetURL(), ShouldEqual, "http://graylog.test.com/endpoint")
					})
				})

				Convey("with empty endpoint", func() {
					testServer.Endpoint = ""

					Convey("the URL should be correct", func() {
						So(testServer.GetURL(), ShouldEqual, "http://graylog.test.com")
					})
				})

				Convey("with default endpoint", func() {

					Convey("the URL should be correct", func() {
						So(testServer.GetURL(), ShouldEqual, "http://graylog.test.com/api")
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
						So(testServer.GetURL(), ShouldEqual, "https://graylog.test.com:9001/endpoint")
					})
				})

				Convey("with empty endpoint", func() {
					testServer.Endpoint = ""

					Convey("the URL should be correct", func() {
						So(testServer.GetURL(), ShouldEqual, "https://graylog.test.com:9001")
					})
				})

				Convey("with default endpoint", func() {

					Convey("the URL should be correct", func() {
						So(testServer.GetURL(), ShouldEqual, "https://graylog.test.com:9001/api")
					})
				})
			})

			Convey("with standard port", func() {
				testServer.Port = "443"

				Convey("with provided endpoint", func() {
					testServer.Endpoint = "/endpoint"

					Convey("the URL should be correct", func() {
						So(testServer.GetURL(), ShouldEqual, "https://graylog.test.com/endpoint")
					})
				})

				Convey("with empty endpoint", func() {
					testServer.Endpoint = ""

					Convey("the URL should be correct", func() {
						So(testServer.GetURL(), ShouldEqual, "https://graylog.test.com")
					})
				})

				Convey("with default endpoint", func() {

					Convey("the URL should be correct", func() {
						So(testServer.GetURL(), ShouldEqual, "https://graylog.test.com/api")
					})
				})
			})

		})
	})

}
