package mata

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQueryModel(t *testing.T) {
	Convey("Given a Query model created with the constructor", t, func() {
		testQuery := DefaultQuery()

		Convey("the Default Values should be correct", func() {
			So(testQuery.Terms, ShouldEqual, "*")
		})
	})
}
