package slice

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsContains(t *testing.T) {
	Convey("Test String contain", t, func() {
		a := "111"
		b := []string{"111", "222"}
		c := []string{"222", "333"}
		So(IsContains(a, b), ShouldBeTrue)
		So(IsContains(a, c), ShouldBeFalse)
	})
}

func TestIsContainsInterface(t *testing.T) {
	Convey("Test interface contain", t, func() {
		a := 1
		b := []interface{}{"111", "222", 1, 34}
		c := []interface{}{"222", "333", 90}
		So(IsContainsInterface(a, b), ShouldBeTrue)
		So(IsContainsInterface(a, c), ShouldBeFalse)
	})
}

func TestDeleteValueByIndex(t *testing.T) {
	Convey("Test DeleteValueByIndex", t, func() {
		a := []interface{}{"111", "222", 1, 34}
		t.Log(DeleteValueByIndex(a, 3))
		//So(DeleteValueByIndex(a, 1), ShouldEqual, []interface{}{"111", 1, 34})
		//So(DeleteValueByIndex(a, 0), ShouldEqual, []interface{}{"222", 1, 34})
	})
}

func BenchmarkDeleteValueByIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []interface{}{"1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4"}
		DeleteValueByIndex(a, 2)
	}
}

func BenchmarkDeleteValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []interface{}{"1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4"}
		ret := make([]interface{}, 0)
		for idx, val := range a {
			if idx != 2 {
				ret = append(ret, val)
			}
		}
	}
}
