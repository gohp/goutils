package gotime

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestFormatDatetime(t *testing.T) {
	Convey("test format datetime", t, func() {
		// gen datetime 2020-02-25 10:12:09.0
		t1 := time.Date(2020, 02, 25, 10, 12, 9, 0, time.Local)
		So(FormatDatetime(t1, "YY-MM"), ShouldEqual, "20-02")
		So(FormatDatetime(t1, "YYYY-MM"), ShouldEqual, "2020-02")
		So(FormatDatetime(t1, "YYYY-MM-DD"), ShouldEqual, "2020-02-25")
		So(FormatDatetime(t1, "YY/M/D"), ShouldEqual, "20/2/25")

		So(FormatDatetime(t1, "YYYY/MM/DD HH:mm:ss"), ShouldEqual, "2020/02/25 10:12:09")
	})
}