package set

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/6
 */

func Test(t *testing.T) {
	Convey("Test Set", t, func() {
		So(nil, ShouldEqual, nil)

		Convey("Set Add Value", func() {
			s := NewSet()
			So(s.Add(1), ShouldBeTrue)
			So(s.Add(1), ShouldBeFalse)
			So(s.Add(2), ShouldBeTrue)
		})

		Convey("Set Del Value", func() {
			s := NewSet()
			s.Add(1)
			So(s.Add(1), ShouldBeFalse)
			s.Del(1)
			So(s.Add(1), ShouldBeTrue)
		})

		Convey("Set Flush", func() {
			s := NewSet()
			s.Add(1)
			s.Add(2)
			s.Flush()
			So(s.Add(1), ShouldBeTrue)
			So(s.Add(2), ShouldBeTrue)
		})

		Convey("Seontains Value", func() {
			s := NewSet()
			s.Add(1)
			s.Add(2)
			s.Add("3")

			So(s.Contains(1), ShouldBeTrue)
			So(s.Contains(2), ShouldBeTrue)
			So(s.Contains(1, 2, "3"), ShouldBeTrue)
			So(s.Contains(3), ShouldBeFalse)
		})

		Convey("Set Len", func() {
			s := NewSet()
			s.Add(1)
			s.Add(2)
			s.Add("3")

			So(s.Len(), ShouldEqual, 3)
			s.Flush()
			So(s.Len(), ShouldEqual, 0)
		})

		Convey("Set Union", func() {
			s := NewSet()
			s2 := NewSet()
			s.Add(1)
			s.Add(2)
			s.Add("3")
			s2.Add(4)

			s3 := s.Union(s2)
			//t.Log(s3.ToString())
			So(s3.Contains(1, 2, "3", 4), ShouldBeTrue)
		})

		Convey("Set Difference", func() {
			s := NewSet()
			s2 := NewSet()
			s.Add(1)
			s.Add(2)
			s.Add("3")

			s2.Add(4)
			s2.Add("3")

			s3 := s.Difference(s2)
			s4 := s2.Difference(s)
			//t.Log(s3.ToString())
			So(s3.Contains(1, 2, "3", 4), ShouldBeFalse)
			So(s3.Contains(1, 2), ShouldBeTrue)

			So(s4.Contains(1, 2, "3", 4), ShouldBeFalse)
			So(s4.Contains(4), ShouldBeTrue)
		})

		Convey("Set Intersect", func() {
			s := NewSet()
			s2 := NewSet()
			s.Add(1)
			s.Add(2)
			s.Add("3")

			s2.Add(4)
			s2.Add("3")

			s3 := s.Intersect(s2)
			So(s3.Contains(1, 2, "3", 4), ShouldBeFalse)
			So(s3.Contains("3"), ShouldBeTrue)
		})
	})
}
