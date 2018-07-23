package search_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tahitianstud/mata-cli/internal/search"
)

func TestDefault(t *testing.T) {
	Convey("Given a Query model created with the constructor", t, func() {
		testQuery := search.Default()

		Convey("the Default Values should be correct", func() {
			So(testQuery.Terms, ShouldEqual, "*")
		})
	})
}
