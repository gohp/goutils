package hash

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMd5Byte(t *testing.T) {
	Convey("md5 byte", t, func() {
		So(Md5Byte([]byte("123")), ShouldEqual, "202cb962ac59075b964b07152d234b70")
	})
}

func TestMd5String(t *testing.T) {
	Convey("md5 string", t, func() {
		So(Md5String("123"), ShouldEqual, "202cb962ac59075b964b07152d234b70")
	})
}

func TestMd5File(t *testing.T) {
	Convey("md5 file", t, func() {
		v, err := Md5File("./test.txt")
		So(err, ShouldBeNil)
		So(v, ShouldEqual, "202cb962ac59075b964b07152d234b70")
	})
}

func TestSha1Byte(t *testing.T) {
	Convey("sha1 byte", t, func() {
		So(Sha1Byte([]byte("123")), ShouldEqual, "40bd001563085fc35165329ea1ff5c5ecbdbbeef")
	})
}

func TestSha1String(t *testing.T) {
	Convey("sha1 string", t, func() {
		So(Sha1String("123"), ShouldEqual, "40bd001563085fc35165329ea1ff5c5ecbdbbeef")
	})
}

func TestSha1File(t *testing.T) {
	Convey("sha1 file", t, func() {
		v, err := Sha1File("./test.txt")
		So(err, ShouldBeNil)
		So(v, ShouldEqual, "40bd001563085fc35165329ea1ff5c5ecbdbbeef")
	})
}

func TestSha256Byte(t *testing.T) {
	Convey("sha256 byte", t, func() {
		So(Sha256Byte([]byte("123")), ShouldEqual, "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3")
	})
}

func TestSha256String(t *testing.T) {
	Convey("sha256 string", t, func() {
		So(Sha256String("123"), ShouldEqual, "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3")
	})
}

func TestSha256File(t *testing.T) {
	Convey("sha256 file", t, func() {
		v, err := Sha256File("./test.txt")
		So(err, ShouldBeNil)
		So(v, ShouldEqual, "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3")
	})
}

func TestSha512Byte(t *testing.T) {
	Convey("sha512 byte", t, func() {
		So(Sha512Byte([]byte("123")), ShouldEqual, "3c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2")
	})
}

func TestSha512String(t *testing.T) {
	Convey("sha512 string", t, func() {
		So(Sha512String("123"), ShouldEqual, "3c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2")
	})
}

func TestSha512File(t *testing.T) {
	Convey("sha512 file", t, func() {
		v, err := Sha512File("./test.txt")
		So(err, ShouldBeNil)
		So(v, ShouldEqual, "3c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2")
	})
}
