package verify

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	helper "github.com/wzyonggege/goutils/gin-helper"
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

// Verify is is the verify model.
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
			params := req.PostForm

			//// check timestamp is not empty
			//if params.Get("ts") == "" {
			//	gin_helper.WriteError(ctx, 10001, "timestamp is required.")
			//	ctx.Abort()
			//}

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
