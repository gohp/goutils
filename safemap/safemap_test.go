package safemap

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSafeMap(t *testing.T) {
	Convey("Test Safe Map", t, func() {

		var sm *SafeMap
		sm = NewSafeMap()

		Convey("Set key", func(t C) {
			r1 := sm.Set("key1", "122")
			r2 := sm.Set(2, "123")
			t.So(r1, ShouldBeTrue)
			t.So(r2, ShouldBeTrue)
		})

		Convey("Get key", func(t C) {
			_ = sm.Set("key1", "122")
			v1 := sm.Get("key1")
			v2 := sm.Get("key2")
			t.So(v1, ShouldEqual, "122")
			t.So(v2, ShouldBeNil)
		})

		Convey("Delete key", func(t C) {
			_ = sm.Set("key1", "122")
			t.So(sm.Get("key1"), ShouldEqual, "122")

			sm.Delete("key1")
			t.So(sm.Get("key1"), ShouldBeNil)
		})

		Convey("Exist key", func(t C) {
			_ = sm.Set("key1", "122")
			t.So(sm.Exist("key1"), ShouldBeTrue)

			sm.Delete("key1")
			t.So(sm.Exist("key1"), ShouldBeFalse)
		})

		Convey("Len key", func(t C) {
			_ = sm.Set("key1", "121")
			_ = sm.Set("key2", "122")
			_ = sm.Set("key3", "123")

			t.So(sm.Len(), ShouldEqual, 3)

			sm.Delete("key1")
			t.So(sm.Len(), ShouldEqual, 2)
		})

		Convey("items key", func(t C) {
			_ = sm.Set("key1", "121")
			_ = sm.Set("key2", "122")
			_ = sm.Set("key3", "123")

			//var keys []interface{}
			//keys = append(keys, "key1",  "key2", "key3")
			t.So(len(sm.Keys()), ShouldEqual, 3)
			t.So(len(sm.Values()), ShouldEqual, 3)
			t.So(len(sm.Items()), ShouldEqual, 3)
		})
	})
}
