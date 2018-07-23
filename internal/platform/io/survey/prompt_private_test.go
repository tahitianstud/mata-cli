package survey

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Definition describes the data used to define a server
type Definition struct {
	Scheme string `description:"the server connection scheme" type:"select" choice:"http|https"`

	Host string `description:"the server hostname"`

	Port string `description:"the server port"`

	Endpoint string `description:"the server API endpoint"`

	Username string `description:"the username"`

	Password string `description:"the password" type:"password"`
}

// Defaults return an instance of Definition with sensible default values
func Defaults() Definition {
	return Definition{

		Host:     "",

		Port:     "9000",

		Endpoint: "/api",

		Username: "",

		Password: "",

		Scheme:   "http",

	}
}

func TestPrivateMethods(t *testing.T) {

	Convey("Given a default server config struct", t, func() {
		server := Defaults()

		Convey("the created questions array should be correct ", func() {
			questions, err := createQuestionsFrom(server)

			So(err, ShouldBeNil)

			So(questions, ShouldNotBeNil)

			So(len(questions), ShouldEqual, 6)

			So(questions[0].Name, ShouldEqual, "Scheme")

			schemePrompt := fmt.Sprintf("%v", questions[0].Prompt)
			expectedSchemePrompt := fmt.Sprintf("%v", &survey.Select{
				Message: "the server connection scheme",
				Options: []string{"http", "https"},
				Default: "http",
			})

			So(schemePrompt, ShouldEqual, expectedSchemePrompt)

			So(questions[5].Name, ShouldEqual, "Password")

			passwordPrompt := fmt.Sprintf("%v", questions[5].Prompt)
			expectedPasswordPrompt := fmt.Sprintf("%v", &survey.Password{
				Message: "the password",
			})
			So(passwordPrompt, ShouldEqual, expectedPasswordPrompt)
		})
	})
}
