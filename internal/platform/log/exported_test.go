package log_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/platform/log/mocks"
	"github.com/tahitianstud/mata-cli/internal/platform/log"
	"errors"
)

func TestExported(t *testing.T) {
	Convey("Given a mock of the log interface", t, func() {

		mockLog := mocks.Logger{}
		log.Log = &mockLog

		Convey("with mocked call to SetLevelTo", func() {

			mockLog.On("SetLevelTo", "DEBUG").Return("DEBUG")
			mockLog.On("SetLevelTo", "INFO").Return("INFO")
			mockLog.On("SetLevelTo", "WARN").Return("WARN")
			mockLog.On("SetLevelTo", "ERROR").Return("ERROR")
			mockLog.On("SetLevelTo", "FATAL").Return("FATAL")

			Convey("calling the exported SetLevelTo should call the mocked call", func() {

				log.SetLevelTo(log.DEBUG)
				log.SetLevelTo(log.INFO)
				log.SetLevelTo(log.WARNING)
				log.SetLevelTo(log.ERROR)
				log.SetLevelTo(log.FATAL)

				So(mockLog.AssertExpectations(t), ShouldBeTrue)

			})

		})


		Convey("with mocked call to DieIf", func() {

			mockLog.On("DieIf", errors.New("a new error")).Return("DIED")

			Convey("calling the exported DieIf should call the mocked call", func() {

				log.DieIf(errors.New("a new error"))

				So(mockLog.AssertExpectations(t), ShouldBeTrue)

			})

		})


		Convey("with mocked call to DebugIf", func() {

			mockLog.On("WriteIf", errors.New("a new error"), "DEBUG").Return("DEBUG")

			Convey("calling the exported DebugIf should call the mocked call", func() {

				log.DebugIf(errors.New("a new error"))

				So(mockLog.AssertExpectations(t), ShouldBeTrue)

			})

		})


		Convey("with mocked call to InfoIf", func() {

			mockLog.On("WriteIf", errors.New("a new error"), "INFO").Return("INFO")

			Convey("calling the exported InfoIf should call the mocked call", func() {

				log.InfoIf(errors.New("a new error"))

				So(mockLog.AssertExpectations(t), ShouldBeTrue)

			})

		})

		Convey("with mocked call to WarningIf", func() {

			mockLog.On("WriteIf", errors.New("a new error"), "WARN").Return("WARN")

			Convey("calling the exported WarningIf should call the mocked call", func() {

				log.WarningIf(errors.New("a new error"))

				So(mockLog.AssertExpectations(t), ShouldBeTrue)

			})

		})

		Convey("with mocked call to ErrorIf", func() {

			mockLog.On("WriteIf", errors.New("a new error"), "ERROR").Return("ERROR")

			Convey("calling the exported ErrorIf should call the mocked call", func() {

				log.ErrorIf(errors.New("a new error"))

				So(mockLog.AssertExpectations(t), ShouldBeTrue)

			})

		})

		Convey("with mocked call to DebugWith", func() {

			mockLog.On("WriteAt", "DEBUG", "a message").Return("WROTE DEBUG")

			Convey("calling the exported DebugWith should call the mocked call", func() {

				log.DebugWith("a message")

				So(mockLog.AssertExpectations(t), ShouldBeTrue)

			})

		})

		Convey("with mocked call to InfoWith", func() {

			mockLog.On("WriteAt", "INFO", "a message").Return("WROTE INFO")

			Convey("calling the exported InfoWith should call the mocked call", func() {

				log.InfoWith("a message")

				So(mockLog.AssertExpectations(t), ShouldBeTrue)

			})

		})

		Convey("with mocked call to WarnWith", func() {

			mockLog.On("WriteAt", "WARN", "a message").Return("WROTE WARNING")

			Convey("calling the exported WarnWith should call the mocked call", func() {

				log.WarnWith("a message")

				So(mockLog.AssertExpectations(t), ShouldBeTrue)

			})

		})

		Convey("with mocked call to ErrorWith", func() {

			mockLog.On("WriteAt", "ERROR", "a message").Return("WROTE ERROR")

			Convey("calling the exported ErrorWith should call the mocked call", func() {

				log.ErrorWith("a message")

				So(mockLog.AssertExpectations(t), ShouldBeTrue)

			})

		})

		Convey("with call to Data", func() {

			data := log.Data("key", "a message value")

			Convey("data should be of the format key|value", func() {
				So(data, ShouldEqual, "key|a message value")
			})

		})

	})
}



