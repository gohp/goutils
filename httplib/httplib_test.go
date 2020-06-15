package httplib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/15
 */

const baseHost = "http://127.0.0.1:7777"

type RespData struct {
	Params string `json:"params"`
	Success bool `json:"success"`
	UA string `json:"ua"`
	Charset string `json:"charset"`
	TestHeader string `json:"test_header"`
	ContentType string `json:"content_type"`
}

func TestGet(t *testing.T) {
	Convey("Test Http Get", t, func() {
		response, err := Get(baseHost+"/test").
			Header("User-Agent", "Golang").
			Header("Test-Header", "asasasas").
			String()
		So(err, ShouldBeNil)
		t.Log(response)
	})
}

func TestGet_Param(t *testing.T) {
	Convey("Test Http Get", t, func() {
		req := Get(baseHost+"/test").
			Header("User-Agent", "Golang").
			Header("Test-Header", "asasasas")
		var resp RespData
		// ?key1=value1&key2=value2...
		req.Param("key1", "value1")
		req.Param("key2", "value2")
		err := req.ToJson(&resp)

		So(err, ShouldBeNil)
		So(resp.Success, ShouldEqual, true)
		So(resp.UA, ShouldEqual, "Golang")
		So(resp.TestHeader, ShouldEqual, "asasasas")
		So(resp.Params, ShouldContainSubstring, "key1=value1")
		So(resp.Params, ShouldContainSubstring, "key2=value2")
	})
}

func TestPost(t *testing.T) {
	Convey("Test Http Post", t, func() {
		//req := Post(url).Debug(true)

		var resp RespData

		err := Post(baseHost+"/test").
			SetTimeout(20*time.Second, 30*time.Second).
			Header("Connection", "Keep-Alive").
			Header("Charset", "UTF-8").
			Header("User-Agent", "Golang").
			Header("Content-Type", "application/json").
			Header("Test-Header", "asasasas").
			ToJson(&resp)

		So(err, ShouldBeNil)
		So(resp.Success, ShouldEqual, true)
		So(resp.UA, ShouldEqual, "Golang")
		So(resp.TestHeader, ShouldEqual, "asasasas")
		So(resp.Charset, ShouldEqual, "UTF-8")
		So(resp.ContentType, ShouldEqual, "application/json")
	})
}
