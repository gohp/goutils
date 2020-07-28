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

func TestToInt64(t *testing.T) {
	Convey("Test ToInt64", t, func() {
		Convey("Test int ToInt64", func() {
			var i1 int = 1
			var i2 int8 = 1
			var i3 int16 = 1
			var i4 int32 = 1
			var i5 int64 = 1
			i1r, err := ToInt64(i1)
			So(i1r, ShouldEqual, i5)
			So(err, ShouldBeNil)

			i2r, err := ToInt64(i2)
			So(i2r, ShouldEqual, i5)
			So(err, ShouldBeNil)

			i3r, err := ToInt64(i3)
			So(i3r, ShouldEqual, i5)
			So(err, ShouldBeNil)

			i4r, err := ToInt64(i4)
			So(i4r, ShouldEqual, i5)
			So(err, ShouldBeNil)

			i5r, err := ToInt64(i5)
			So(i5r, ShouldEqual, i5)
			So(err, ShouldBeNil)
		})

		Convey("Test uint ToInt64", func() {
			var i int64 = 1

			var u1 uint = 1
			var u2 uint8 = 1
			var u3 uint16 = 1
			var u4 uint32 = 1
			var u5 uint64 = 1
			u1r, err := ToInt64(u1)
			So(u1r, ShouldEqual, i)
			So(err, ShouldBeNil)

			u2r, err := ToInt64(u2)
			So(u2r, ShouldEqual, i)
			So(err, ShouldBeNil)

			u3r, err := ToInt64(u3)
			So(u3r, ShouldEqual, i)
			So(err, ShouldBeNil)

			u4r, err := ToInt64(u4)
			So(u4r, ShouldEqual, i)
			So(err, ShouldBeNil)

			u5r, err := ToInt64(u5)
			So(u5r, ShouldEqual, i)
			So(err, ShouldBeNil)
		})

		Convey("Test non-numeric ToInt64", func() {
			var u1 string = "1"
			var u2 bool = true
			_, err := ToInt64(u1)
			So(err, ShouldNotBeNil)

			_, err = ToInt64(u2)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestAnyToString(t *testing.T) {
	Convey("Test AnyToString", t, func() {
		res, err := AnyToString(1)
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "1")

		res, err = AnyToString(int64(1))
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "1")

		res, err = AnyToString(-1)
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "-1")

		res, err = AnyToString([]byte("1"))
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "1")

		res, err = AnyToString(true)
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "true")

		res, err = AnyToString(false)
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "false")
	})
}

func TestFormatByteSize(t *testing.T) {
	Convey("Test FormatByteSize", t, func() {
		So(FormatByteSize(1), ShouldEqual, "1.00B")
		So(FormatByteSize(1024), ShouldEqual, "1.00KB")
		So(FormatByteSize(1024*1024), ShouldEqual, "1.00MB")
		So(FormatByteSize(1024*1024*1024), ShouldEqual, "1.00GB")
		So(FormatByteSize(1024*1024*1024*1024), ShouldEqual, "1.00TB")
		So(FormatByteSize(1024*1024*1024*1024*1024), ShouldEqual, "1.00EB")
	})
}
