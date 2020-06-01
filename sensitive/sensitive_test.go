package sensitive

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFilter(t *testing.T) {
	Convey("test cal distance", t, func() {
		dirtyFilter := New("")

		Convey("test add word", func(t C) {
			dirtyFilter.AddWord("fuck")
			isIllegal, word := dirtyFilter.Validate("fuck")

			So(isIllegal, ShouldBeFalse)
			So(word, ShouldEqual, "fuck")
		})

		Convey("test del word", func(t C) {
			dirtyFilter.AddWord("fuck")
			isIllegal, _ := dirtyFilter.Validate("fuck")
			So(isIllegal, ShouldBeFalse)
			dirtyFilter.DelWord("fuck")
			isIllegal, _ = dirtyFilter.Validate("fuck")
			So(isIllegal, ShouldBeTrue)
		})

		Convey("test replace word", func(t C) {
			dirtyFilter.AddWord("fuck")
			isIllegal, _ := dirtyFilter.Validate("fuck")
			So(isIllegal, ShouldBeFalse)
			ret := dirtyFilter.Replace("fuck", rune('*'))
			So(ret, ShouldEqual, "****")
		})
	})
}
