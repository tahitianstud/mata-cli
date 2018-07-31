package file_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/platform/io/file"
	"os"
)

func TestWrite(t *testing.T) {

	Convey("Given a non-existing test file location", t, func() {

		testFileLocation := "/tmp/non-existing/file"
		os.Remove(testFileLocation)

		Convey("writing should create the file", func() {

			err := file.Write(testFileLocation, []byte("data"))

			So(err, ShouldBeNil)

			info, err := os.Stat(testFileLocation)

			So(err, ShouldBeNil)
			So(info.IsDir(), ShouldBeFalse)
			So(info.Mode(), ShouldEqual, 0700)

		})

	})

	Convey("Given an existing test file location", t, func() {

		testFileLocation := "/tmp/non-existing/file"
		os.Create(testFileLocation)

		Convey("writing w/ overwrite true should create the file", func() {

			err := file.Write(testFileLocation, []byte("data"), true)

			So(err, ShouldBeNil)

			info, err := os.Stat(testFileLocation)

			So(err, ShouldBeNil)
			So(info.IsDir(), ShouldBeFalse)
			So(info.Mode(), ShouldEqual, 0700)

		})

		Convey("writing w/ overwrite false should error out", func() {

			err := file.Write(testFileLocation, []byte("data"), false)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "cannot write: file already exist and overwrite is false")

		})


	})

	Convey("Given a non-existing test file location in a restricted area", t, func() {

		os.Remove("/tmp/restricted")
		os.Mkdir("/tmp/restricted", 0400)

		testFileLocation := "/tmp/restricted/file"

		Convey("writing should fail", func() {

			err := file.Write(testFileLocation, []byte("data"))

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "cannot write: file location is not writeable")

		})

	})

}

func TestWriteInside(t *testing.T) {

	fileLocation := "/tmp/test"
	os.RemoveAll(fileLocation)

	Convey("Given non-existing test location and filename", t, func() {

		filename := "testerFile"

		Convey("write should work as promised", func() {

			err := file.WriteInside(fileLocation, filename, []byte("data"), false)

			So(err, ShouldBeNil)

		})

	})

	Convey("Given existing test location with terminating / and filename", t, func() {

		fileLocation2 := "/tmp/test/"
		os.Remove(fileLocation2)

		filename := "testerFile"

		Convey("write should work as promised", func() {

			err := file.WriteInside(fileLocation2, filename, []byte("data"), true)

			So(err, ShouldBeNil)

		})


	})

}
