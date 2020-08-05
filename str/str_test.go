package str

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
* @Author: Jam Wong
* @Date: 2020/8/4
 */

func TestLongestPalindromic(t *testing.T) {
	Convey("Test ", t, func() {
		So(LongestPalindromic("babad"), ShouldEqual, "bab")
		So(LongestPalindromic("cbbd"), ShouldEqual, "bb")
		So(LongestPalindromic("cbbc"), ShouldEqual, "cbbc")
	})

}
