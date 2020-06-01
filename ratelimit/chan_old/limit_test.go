package chan_old

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestTokenBucket_TakeOne(t *testing.T) {
	Convey("test chan old bucket", t, func() {
		tb := New(100, time.Millisecond*10)
		time.Sleep(time.Second * 1)
		ok := tb.TakeOne()
		So(ok, ShouldBeTrue)
	})
}
