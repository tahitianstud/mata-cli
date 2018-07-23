package file_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"github.com/tahitianstud/mata-cli/internal/platform/io/file"
	"os"
	"path/filepath"
)

func TestDeleteIfExists(t *testing.T) {

	Convey("Given a temp file", t, func() {

		tempDirectory := "/tmp"
		tempFile, _ := ioutil.TempFile(tempDirectory, "testfile")
		filepath := tempFile.Name()

		Convey("deleting the file should work", func() {

			ok := file.DeleteIfExists(filepath)

			So(ok, ShouldBeTrue)

		})

		Convey("deleting the file forcefully", func() {

			os.RemoveAll(filepath)

			Convey("deleting the file a second time should return an error", func() {

				ok := file.DeleteIfExists(filepath)

				So(ok, ShouldBeFalse)

			})

		})


	})

}

func TestDeleteIfExistsIn(t *testing.T) {

	Convey("Given a temp file", t, func() {

		tempDirectory := "/tmp"
		tempFile, _ := ioutil.TempFile(tempDirectory, "testfile")
		temppath := tempFile.Name()
		filename := filepath.Base(temppath)
		basePath := filepath.Dir(temppath)

		Convey("deleting the file should work", func() {

			ok := file.DeleteIfExistsIn(basePath, filename)

			So(ok, ShouldBeTrue)

		})

	})

}
