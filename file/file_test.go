package file

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSelfPath(t *testing.T) {
	convey.Convey("test self path", t, func() {
		convey.So(SelfPath(), convey.ShouldNotEqual, "")
	})
}

func TestSelfDir(t *testing.T) {
	convey.Convey("test self dir", t, func() {
		convey.So(SelfDir(), convey.ShouldNotEqual, "")
	})
}

func TestFileExists(t *testing.T) {
	convey.Convey("test file exist", t, func() {
		convey.So(FileExists("./file.go"), convey.ShouldBeTrue)
		convey.So(FileExists("./not_exist_file.go"), convey.ShouldBeFalse)
	})
}

func TestReadStringsFromFile(t *testing.T) {

}

func TestWriteStringsToFile(t *testing.T) {

}
