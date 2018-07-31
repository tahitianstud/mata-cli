package file_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/platform/io/file"
	"os"
	"io/ioutil"
	"fmt"
)

func TestRead(t *testing.T) {
	Convey("Given a non-existing file", t, func() {

		testFile := "/tmp/filethatdoesnotexist"
		testResult := "N/A"

		Convey("reading should return an error", func() {

			err := file.Read(testFile, &testResult)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "no such file or directory")
			So(testResult, ShouldEqual, "N/A")

		})

	})

	Convey("Given a non-accessible file", t, func() {

		testFile := "/tmp/nonaccessiblefile"
		ioutil.WriteFile(testFile, []byte(""), 0100)
		testResult := "N/A"

		Convey("reading should return an error", func() {

			err := file.Read(testFile, &testResult)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "permission denied")
			So(testResult, ShouldEqual, "N/A")

		})

		os.RemoveAll(testFile)

	})

	Convey("Given an accessible file", t, func() {

		testFile := "/tmp/accessiblefile"
		ioutil.WriteFile(testFile, []byte("file contents"), 0400)
		testResult := "N/A"

		Convey("reading should return an error", func() {

			err := file.Read(testFile, &testResult)

			So(err, ShouldBeNil)
			So(testResult, ShouldEqual, "file contents")

		})

		os.RemoveAll(testFile)

	})
}

func TestReadFrom(t *testing.T) {

	Convey("Given a non-existing file", t, func() {

		testLocation := "/tmp/non-existing"
		testFile := "filethatdoesnotexist"
		testResult := "N/A"

		Convey("reading from location should return an error", func() {

			err := file.ReadFrom(testLocation, testFile, &testResult)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "no such file or directory")
			So(testResult, ShouldEqual, "N/A")

		})

	})

	Convey("Given a non-accessible file", t, func() {

		testLocation := "/tmp/"
		testFile := "nonaccessiblefile"
		ioutil.WriteFile(fmt.Sprintf("%s%s", testLocation, testFile), []byte(""), 0100)
		testResult := "N/A"

		Convey("reading from location should return an error", func() {

			err := file.ReadFrom(testLocation, testFile, &testResult)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "permission denied")
			So(testResult, ShouldEqual, "N/A")

		})

		os.RemoveAll(fmt.Sprintf("%s%s", testLocation, testFile))

	})

	Convey("Given an accessible file", t, func() {

		testLocation := "/tmp/"
		testFile := "accessiblefile"
		ioutil.WriteFile(fmt.Sprintf("%s%s", testLocation, testFile), []byte("file contents"), 0700)
		testResult := "N/A"

		Convey("reading should return an error", func() {

			err := file.ReadFrom(testLocation, testFile, &testResult)

			So(err, ShouldBeNil)
			So(testResult, ShouldEqual, "file contents")

		})

		os.RemoveAll(fmt.Sprintf("%s%s", testLocation, testFile))

	})

}
