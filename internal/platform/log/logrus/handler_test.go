package logrus_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/tahitianstud/mata-cli/internal/platform/log/logrus"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
)

func TestLogHandler(t *testing.T) {

	Convey("Given a logrus handler", t, func() {

		logger := logrus.GetLogrusHandler()
		logger.SetLevelTo("DEBUG")

		Convey("and a buffer to check the output of the command", func() {

			var testBuffer bytes.Buffer
			logger.Output = &testBuffer

			Convey("Calling WriteIf with a non-nil error", func() {

				testBuffer.Reset()
				logger.WriteIf(fmt.Errorf("debug test"), "DEBUG")

				Convey("Should output the correct debug message", func() {

					output := testBuffer.String()

					So(output, ShouldContainSubstring, "level=debug")
					So(output, ShouldContainSubstring, `msg="debug test"`)

				})

				testBuffer.Reset()
				logger.WriteIf(fmt.Errorf("info test"), "INFO")

				Convey("Should output the correct info message", func() {

					So(testBuffer.String(), ShouldContainSubstring, "level=info")
					So(testBuffer.String(), ShouldContainSubstring, `msg="info test"`)

				})

				testBuffer.Reset()
				logger.WriteIf(fmt.Errorf("warn test"), "WARN")

				Convey("Should output the correct warning message", func() {

					So(testBuffer.String(), ShouldContainSubstring, "level=warn")
					So(testBuffer.String(), ShouldContainSubstring, `msg="warn test"`)

				})

				testBuffer.Reset()
				logger.WriteIf(fmt.Errorf("error test"), "ERROR")

				Convey("Should output the correct error message", func() {

					So(testBuffer.String(), ShouldContainSubstring, "level=error")
					So(testBuffer.String(), ShouldContainSubstring, `msg="error test"`)

				})

			})

			Convey("Calling WriteIf with a nil error", func() {

				testBuffer.Reset()

				logger.WriteIf(nil, "ERROR")

				Convey("Should output nothing", func() {

					So(testBuffer.String(), ShouldBeEmpty)

				})

			})

			Convey("Calling WriteIf with a non-nil error but non-standard level", func() {

				testBuffer.Reset()

				logger.WriteIf(errors.New("Very bad error"), "FUBAR")

				Convey("Should output nothing", func() {

					So(testBuffer.String(), ShouldBeEmpty)

				})

			})


			Convey("Calling WriteAt with a message but without metadata", func() {

				testBuffer.Reset()

				logger.WriteAt("DEBUG", "my super message")

				Convey("Should output the correct DEBUG message in the buffer", func() {

					So(testBuffer.String(), ShouldContainSubstring, "level=debug")
					So(testBuffer.String(), ShouldContainSubstring, `msg="my super message"`)

				})

			})

			Convey("Calling WriteAt with a warn message but without metadata", func() {

				testBuffer.Reset()

				logger.WriteAt("WARN", "my super message")

				Convey("Should output the correct WARN message in the buffer", func() {

					So(testBuffer.String(), ShouldContainSubstring, "level=warn")
					So(testBuffer.String(), ShouldContainSubstring, `msg="my super message"`)

				})

			})

			Convey("Calling WriteAt with an error message but without metadata", func() {

				testBuffer.Reset()

				logger.WriteAt("ERROR", "my super message")

				Convey("Should output the correct ERROR message in the buffer", func() {

					So(testBuffer.String(), ShouldContainSubstring, "level=error")
					So(testBuffer.String(), ShouldContainSubstring, `msg="my super message"`)

				})

			})

			Convey("Calling WriteAt with a message with metadata", func() {

				testBuffer.Reset()

				logger.WriteAt("INFO", "my super message",
					"superman|was here", "frodo|got a ring", "link|went to the past")

				Convey("Should output the correct INFO message in the buffer", func() {

					So(testBuffer.String(), ShouldContainSubstring, "level=info")
					So(testBuffer.String(), ShouldContainSubstring, `msg="my super message"`)
					So(testBuffer.String(), ShouldContainSubstring, `superman="was here"`)
					So(testBuffer.String(), ShouldContainSubstring, `frodo="got a ring"`)
					So(testBuffer.String(), ShouldContainSubstring, `link="went to the past"`)

				})

			})

			Convey("Calling WriteAt with a message with some bad metadata", func() {

				testBuffer.Reset()

				logger.WriteAt("INFO", "my super message",
					"batman|was here", "wonder woman,not so much")

				Convey("Should output the correct INFO message in the buffer", func() {

					So(testBuffer.String(), ShouldContainSubstring, "level=info")
					So(testBuffer.String(), ShouldContainSubstring, `batman="was here"`)
					So(testBuffer.String(), ShouldNotContainSubstring, `wonder woman="not so much"`)

				})

			})

		})

	})
}