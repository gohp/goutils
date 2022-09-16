package verify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	helper "github.com/gohp/goutils/gin-helper"
	"github.com/gohp/goutils/hash"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

/**
* @Author: Jam Wong
* @Date: 2020/7/20
 */

var (
	router          *gin.Engine
	verifyTestRoute = "/verify"
	defaultSecret   = "secret"
)

func TestMain(m *testing.M) {
	verifySrv := New(&Config{
		Secret:  defaultSecret,
		ErrCode: 30000,
		ErrMsg:  "not allowed",
	})

	router = gin.Default()
	gin.SetMode(gin.TestMode)
	router.POST(verifyTestRoute, verifySrv.Verify(), func(ctx *gin.Context) {

		helper.WriteResponse(ctx, nil)
	})

	os.Exit(m.Run())
}

func postJson(uri string, params interface{}, router *gin.Engine) (helper.BaseResponse, error) {
	jsonByte, _ := json.Marshal(params)
	fmt.Println(string(jsonByte))
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()
	var response helper.BaseResponse
	body, err := ioutil.ReadAll(result.Body)
	err = json.Unmarshal(body, &response)

	return response, err
}

func TestVerify_Verify(t *testing.T) {
	Convey("Test verify", t, func() {

		Convey("test right sign", func() {
			params := map[string]interface{}{"a": 1, "b": 1}
			data, _ := json.Marshal(params)
			sign := hash.Md5String(strings.ToLower(string(data)) + defaultSecret)
			params["sign"] = sign

			resp, err := postJson(verifyTestRoute, params, router)
			So(err, ShouldBeNil)
			So(resp.Success, ShouldBeTrue)
		})

		Convey("test not sign", func() {
			params := map[string]interface{}{"a": 1, "b": 1}
			resp, err := postJson(verifyTestRoute, params, router)
			So(err, ShouldBeNil)
			So(resp.Success, ShouldBeFalse)
			So(resp.Code, ShouldEqual, 30000)
		})

		Convey("test wrong secret", func() {
			params := map[string]interface{}{"a": 1, "b": 1}
			data, _ := json.Marshal(params)
			sign := hash.Md5String(strings.ToLower(string(data)) + "wrong secret")
			params["sign"] = sign
			resp, err := postJson(verifyTestRoute, params, router)
			So(err, ShouldBeNil)
			So(resp.Success, ShouldBeFalse)
			So(resp.Code, ShouldEqual, 30000)
		})

		Convey("test wrong sign key", func() {
			params := map[string]interface{}{"a": 1, "b": 1}
			data, _ := json.Marshal(params)
			sign := hash.Md5String(strings.ToLower(string(data)) + defaultSecret)
			params["sign2"] = sign
			resp, err := postJson(verifyTestRoute, params, router)
			So(err, ShouldBeNil)
			So(resp.Success, ShouldBeFalse)
			So(resp.Code, ShouldEqual, 30000)
		})
	})
}
