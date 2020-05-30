package convert

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInt64ToString(t *testing.T) {
	Convey("Test Int64ToString", t, func() {
		So("1", ShouldEqual, Int64ToString(1))
	})
}

func TestIntToString(t *testing.T) {
	Convey("Test TestIntToString", t, func() {
		So("1", ShouldEqual, IntToString(1))
	})
}

func TestStringToInt(t *testing.T) {
	Convey("Test TestStringToInt", t, func() {
		i, err := StringToInt("1")
		So(err, ShouldBeNil)
		So(1, ShouldEqual, i)
	})
}

func TestStringToInt64(t *testing.T) {
	Convey("Test TestStringToInt64", t, func() {
		i, err := StringToInt64("123")
		So(err, ShouldBeNil)
		So(int64(123), ShouldEqual, i)
	})
}
