package mata

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	modelUtils "gopkg.in/jeevatkm/go-model.v1"
)

func TestConfigModel(t *testing.T) {
	Convey("Given a Config created with the constructor", t, func() {
		testConfig := NewConfig()

		Convey("the Debug Flag should be false", func() {
			So(testConfig.ShowDebugOutput, ShouldEqual, false)
		})

		Convey("the tag for the Debug field should be correct", func() {
			debugTag, _ := modelUtils.Tag(testConfig, "ShowDebugOutput")
			So(debugTag.Get("description"), ShouldEqual, "Enable debug-level output")
		})
	})
}
