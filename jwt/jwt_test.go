package jwt

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestTokenBucket_TakeOne(t *testing.T) {
	Convey("test jwt", t, func() {
		j := New(&Config{
			Key:        "1234",
			ExpireTime: time.Hour * 1,
		})

		token, err := j.EncodeToken("uid_12345", time.Now())
		So(err, ShouldBeNil)
		obj, err := j.DecodeToken(token)
		So(err, ShouldBeNil)
		So(obj.ID, ShouldEqual, "uid_12345")
	})

	Convey("test jwt time out", t, func() {
		j := New(&Config{
			Key:        "1234",
			ExpireTime: time.Second * 2,
		})
		token, err := j.EncodeToken("uid_12345", time.Now())
		time.Sleep(time.Second * 2)
		So(err, ShouldBeNil)
		obj, err := j.DecodeToken(token)
		So(obj, ShouldBeNil)
	})
}
