package date_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/platform/date"
)

func TestStrToDate(t *testing.T) {

	Convey("Given a date string", t, func() {

		testDateStr := "2018-07-27T06:42:57.528+0000"

		Convey("converting should result in the right date", func() {

			convertedDate, err := date.StrToDate(testDateStr)

			year, month, day := convertedDate.Date()
			hour, min, sec := convertedDate.Clock()

			So(err, ShouldBeNil)
			So(year, ShouldEqual, 2018)
			So(month, ShouldEqual, 07)
			So(day, ShouldEqual, 27)
			So(hour, ShouldEqual, 06)
			So(min, ShouldEqual, 42)
			So(sec, ShouldEqual, 57)

		})

	})

	Convey("Given a bad date string", t, func() {

		testDateStr := "20183"

		Convey("converting should return an error", func() {

			_, err := date.StrToDate(testDateStr)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, `Could not find format for "20183"`)

		})

	})

}
