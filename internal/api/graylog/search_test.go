package graylog_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/api/graylog"
)

func TestQueryModel(t *testing.T) {
	Convey("Given a search model created with the constructor", t, func() {
		testQuery := graylog.DefaultSearch()

		Convey("the Default Values should be correct", func() {
			So(testQuery.Query, ShouldEqual, "*")
		})
	})
}
