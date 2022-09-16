package verify

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	helper "github.com/gohp/goutils/gin-helper"
	"io/ioutil"
	"net/url"
	"strings"
	"sync"
)

/**
* @Author: Jam Wong
* @Date: 2020/7/14
* 请求参数加密
 */

type Config struct {
	Secret  string
	ErrCode int
	ErrMsg  string
}

// Verify is to verify model.
type Verify struct {
	lock    sync.RWMutex
	keys    map[string]string
	secret  string
	errCode int
	errMsg  string
}

var _defaultConfig = &Config{
	Secret: "default-secret",
}

/*
加密规则函数: 请求参数加盐之后MD5
*/
func Sign(params url.Values, secret string, lower bool) string {
	data := params.Encode()
	if strings.IndexByte(data, '+') > -1 {
		data = strings.Replace(data, "+", "%20", -1)
	}
	if lower {
		data = strings.ToLower(data)
	}
	digest := md5.Sum([]byte(data + secret))
	return hex.EncodeToString(digest[:])
}

func SignJson(params map[string]interface{}, secret string, lower bool) string {
	dataBytes, _ := json.Marshal(params)
	data := string(dataBytes)
	if lower {
		data = strings.ToLower(data)
	}
	digest := md5.Sum([]byte(data + secret))
	return hex.EncodeToString(digest[:])
}

func New(conf *Config) *Verify {
	if conf == nil {
		conf = _defaultConfig
	}
	v := &Verify{
		keys:    make(map[string]string),
		secret:  conf.Secret,
		errCode: conf.ErrCode,
		errMsg:  conf.ErrMsg,
	}
	return v
}

func (v *Verify) Verify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request

		// necessary
		_ = req.ParseForm()

		if ctx.Request.Method == "POST" {
			if ctx.Request.Header.Get("Content-Type") == "application/json" {
				data, _ := ioutil.ReadAll(ctx.Request.Body)
				var params map[string]interface{}
				if err := json.Unmarshal(data, &params); err != nil {
					helper.WriteError(ctx, v.errCode, v.errMsg)
					ctx.Abort()
				}

				sign, _ := params["sign"]
				delete(params, "sign")

				if hSign := SignJson(params, v.secret, true); hSign != sign {
					helper.WriteError(ctx, v.errCode, v.errMsg)
					ctx.Abort()
				}

			} else {
				params := req.PostForm
				sign := params.Get("sign")
				params.Del("sign")
				defer params.Set("sign", sign)

				if hSign := Sign(params, v.secret, true); hSign != sign {
					helper.WriteError(ctx, v.errCode, v.errMsg)
					ctx.Abort()
				}
			}
		}
	}
}
