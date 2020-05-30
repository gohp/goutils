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
