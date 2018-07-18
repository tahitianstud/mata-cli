package urfave_test

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"github.com/tink-ab/tempfile"
	"github.com/tahitianstud/mata-cli/internal/mata"
	"github.com/tahitianstud/mata-cli/internal/cmd/mata/urfave"
)

func TestConfigAccessor(t *testing.T) {

	Convey("Given a test configuration file", t, func() {
		testConfiguration := `
output:
  debug: true
  verbose: false
`
		tempFile, _ := tempfile.TempFile("/tmp", "test", ".yml")
		tempFile.WriteString(testConfiguration)
		testfile := tempFile.Name()
		defer os.Remove(testfile)

		Convey("And reading the test file as a config", func() {
			testConfig, err := urfave.ReadConfigFromFile(testfile)
			So(err, ShouldBeNil)

			Convey("The Debug Flag should be correctly set", func() {
				So(testConfig.ShowDebugOutput, ShouldEqual, true)
			})

			Convey("The Verbose Flag should be correctly set", func() {
				So(testConfig.VerboseOutput, ShouldEqual, false)
			})
		})

		tempFile.Close()
	})

	Convey("Given another incomplete configuration file", t, func() {
		testConfiguration := `
output:
  verbose: true
`
		tempFile, _ := tempfile.TempFile("/tmp", "test", ".yml")
		tempFile.WriteString(testConfiguration)
		testfile := tempFile.Name()
		defer os.Remove(testfile)

		Convey("and reading the test file as a config", func() {
			testConfig, err := urfave.ReadConfigFromFile(testfile)
			So(err, ShouldBeNil)

			Convey("the Debug Flag should be correctly set", func() {
				So(testConfig.ShowDebugOutput, ShouldEqual, false)
			})

			Convey("the Verbose Flag should be correctly set", func() {
				So(testConfig.VerboseOutput, ShouldEqual, true)
			})
		})

		tempFile.Close()
	})

	Convey("Given a config struct", t, func() {

		testConfig := mata.NewConfig()
		testConfig.VerboseOutput = false
		testConfig.ShowDebugOutput = true

		Convey("and a new configuration context", func() {
			viper.New()

			Convey("and a target configuration file", func() {

				testfile := "/tmp/testMe.yml"

				err := urfave.WriteConfigToFile(testConfig, testfile)
				So(err, ShouldBeNil)

				defer os.Remove(testfile)

				Convey("the config file's contents should be correct", func() {
					bytes, err := ioutil.ReadFile(testfile)
					So(err, ShouldBeNil)
					So(string(bytes), ShouldEqual, `output:
  debug: true
  verbose: false
`)
				})
			})

			Convey("and a target configuration file in a non-existing place", func() {
				testfile := "/tmp/non/existing/directories/here/test4389349893.yml"
				defer os.Remove(testfile)

				err := urfave.WriteConfigToFile(testConfig, testfile)
				So(err, ShouldBeNil)

				Convey("the config file's contents should be correct", func() {
					bytes, err := ioutil.ReadFile(testfile)
					So(err, ShouldBeNil)
					So(string(bytes), ShouldEqual, `output:
  debug: true
  verbose: false
`)
				})

			})
		})
	})
}
