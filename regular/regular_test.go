package regular

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsEmail(t *testing.T) {
	Convey("Test is email", t, func() {
		So(IsEmail("1234@gmail.com"), ShouldBeTrue)
		So(IsEmail("1234@qq.com"), ShouldBeTrue)
		So(IsEmail("1234.gmail.com"), ShouldBeFalse)
		So(IsEmail("1234.gmail.qq@com"), ShouldBeTrue)
	})
}

func TestIsMobile(t *testing.T) {
	Convey("Test is mobile", t, func() {
		So(IsMobile("0755110"), ShouldBeFalse)
		So(IsMobile("8613500001111"), ShouldBeTrue)
		So(IsMobile("+8613500001111"), ShouldBeTrue)
		So(IsMobile("13500001111"), ShouldBeTrue)
		So(IsMobile("1350000111"), ShouldBeFalse)
	})
}

func TestIsBankNo(t *testing.T) {
	Convey("test bank no", t, func() {
		t.Log(IsBankNo("6228481101100634315"))
		So(IsBankNo("6222"), ShouldBeFalse)
	})
}