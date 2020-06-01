package ratelimit

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestTokenBucket_TakeOne(t *testing.T) {
	Convey("test new bucket", t, func() {
		tb := New(10, time.Second*1)
		time.Sleep(time.Second * 1)

		So(tb.Take(1), ShouldBeTrue)
		So(tb.Take(100), ShouldBeFalse)
	})
}
