package rand

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCreateRandomBytes(t *testing.T) {
	convey.Convey("test CreateRandomBytes", t, func() {
		genLength := 10
		s1 := CreateRandomBytes(genLength)
		s2 := CreateRandomBytes(genLength)
		//fmt.Println(string(s1))
		//fmt.Println(string(s2))
		convey.So(s1, convey.ShouldHaveLength, genLength)
		convey.So(s2, convey.ShouldHaveLength, genLength)
		convey.So(s1, convey.ShouldNotEqual, s2)
	})
}
