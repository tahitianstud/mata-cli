package graylog_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/api/graylog"
)

func TestLoginStructs(t *testing.T) {
	Convey("Given an Invalid ticket", t, func() {

		testInvalidTicket := graylog.InvalidLoginTicket

		Convey("values should be correct", func() {

			So(testInvalidTicket.SessionID, ShouldEqual, "N/A")
			So(testInvalidTicket.ValidUntil, ShouldEqual, "N/A")

		})

	})

	Convey("Given a nil ticket", t, func() {

		testTicket := &graylog.LoginTicket{}

		Convey("values should be correct", func() {

			So(testTicket.SessionID, ShouldEqual, "")
			So(testTicket.ValidUntil, ShouldEqual, "")

		})

	})
}