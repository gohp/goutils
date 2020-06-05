package gotime

import (
	j "encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestFormatDatetime(t *testing.T) {
	Convey("test format datetime", t, func() {
		// gen datetime 2020-02-25 10:12:09.0
		Convey("raw time.Time", func(t C) {
			t1 := time.Date(2020, 02, 25, 10, 12, 9, 0, time.Local)
			So(FormatDatetime(t1, "YY-MM"), ShouldEqual, "20-02")
			So(FormatDatetime(t1, "YYYY-MM"), ShouldEqual, "2020-02")
			So(FormatDatetime(t1, "YYYY-MM-DD"), ShouldEqual, "2020-02-25")
			So(FormatDatetime(t1, "YY/M/D"), ShouldEqual, "20/2/25")

			So(FormatDatetime(t1, "YYYY/MM/DD HH:mm:ss"), ShouldEqual, "2020/02/25 10:12:09")
		})

		Convey("gotime.Time", func(t C) {
			t1 := Time(time.Date(2020, 02, 25, 10, 12, 9, 0, time.Local))
			So(t1.FormatX("YY-MM"), ShouldEqual, "20-02")
			So(t1.FormatX("YYYY-MM"), ShouldEqual, "2020-02")
			So(t1.FormatX("YYYY-MM-DD"), ShouldEqual, "2020-02-25")
			So(t1.FormatX("YY/M/D"), ShouldEqual, "20/2/25")

			So(t1.FormatX("YYYY/MM/DD HH:mm:ss"), ShouldEqual, "2020/02/25 10:12:09")
		})
	})
}

func TestGoTime(t *testing.T) {
	Convey("test gotime struct", t, func() {
		type aa struct {
			StartTime Time `json:"start_time"`
		}

		Convey("nil time json", func(t C) {
			a := &aa{}
			data0, _ := j.Marshal(a)
			So(string(data0), ShouldEqual, ``)
		})

		Convey("time json", func(t C) {
			a := &aa{}
			a.StartTime = Time(time.Date(2020, 02, 25, 10, 12, 9, 0, time.Local))
			data, _ := j.Marshal(a)
			So(string(data), ShouldEqual, `{"start_time":"2020-02-25 10:12:09"}`)
		})

		Convey("time json unmarshal", func(t C) {
			a := aa{}
			data := []byte(`{"start_time":"2020-02-25 10:12:09"}`)
			err := j.Unmarshal(data, &a)
			So(err, ShouldBeNil)
			So(a.StartTime.FormatX(TT), ShouldEqual, "2020-02-25 10:12:09")
		})

	})
}
