package json_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/platform/json"
)

type MyStruct struct {
	Str string
	Number int
	Password string `context:"sensitive"`
}

func TestJsonConversions(t *testing.T) {

	Convey("Given a struct with sensitive information", t, func() {

		sensitiveStruct := MyStruct{
			 Str: "value",
			 Number: 10,
			 Password: "sensitiveStuff",
		}

		Convey("converting to JSON should not print out the password", func() {

			jsonString, err := json.ToStringWithContext(&sensitiveStruct)

			So(err, ShouldBeNil)
			So(jsonString, ShouldNotContainSubstring, "sensitiveStuff")
			So(jsonString, ShouldContainSubstring, "**********")

		})
	})

}