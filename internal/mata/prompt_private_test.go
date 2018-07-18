package mata

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	survey "gopkg.in/AlecAivazis/survey.v1"
)

func TestPrivateMethods(t *testing.T) {

	Convey("Given a default server config struct", t, func() {
		server := DefaultServer()

		Convey("the created questions array should be correct ", func() {
			questions, err := createQuestionsFrom(server)

			So(err, ShouldBeNil)

			So(questions, ShouldNotBeNil)

			So(len(questions), ShouldEqual, 6)

			So(questions[0].Name, ShouldEqual, "Scheme")

			schemePrompt := fmt.Sprintf("%v", questions[0].Prompt)
			expectedSchemePrompt := fmt.Sprintf("%v", &survey.Select{
				Message: "the Graylog server connection scheme",
				Options: []string{"http", "https"},
				Default: "http",
			})

			So(schemePrompt, ShouldEqual, expectedSchemePrompt)

			So(questions[5].Name, ShouldEqual, "Password")

			passwordPrompt := fmt.Sprintf("%v", questions[5].Prompt)
			expectedPasswordPrompt := fmt.Sprintf("%v", &survey.Password{
				Message: "the Graylog password",
			})
			So(passwordPrompt, ShouldEqual, expectedPasswordPrompt)
		})
	})
}
