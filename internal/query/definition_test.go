package query_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/query"
)

func TestQueryModel(t *testing.T) {
	Convey("Given a Query model created with the constructor", t, func() {
		testQuery := query.Default()

		Convey("the Default Values should be correct", func() {
			So(testQuery.Terms, ShouldEqual, "*")
		})
	})
}
