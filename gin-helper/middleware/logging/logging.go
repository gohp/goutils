package logging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	gin_helper "github.com/gohp/goutils/gin-helper"
	"github.com/gohp/logger"
	"io/ioutil"
	"regexp"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logging is a middleware function that logs the request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		reg := regexp.MustCompile("(/register|/login|/ws|/metrics|/upload)")
		if reg.MatchString(path) {
			return
		}

		// Skip for the health check requests.
		if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
			return
		}

		// Read the Body content
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// The basic information.
		method := c.Request.Method
		ip := c.ClientIP()
		xRequests := c.MustGet("X-Request-Id")
		headers := c.Request.Header
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw
		logger.Info(
			"request in",
			logger.With("ip", ip),
			logger.With("method", method),
			logger.With("path", path),
			logger.With("requests_body", string(bodyBytes)),
			logger.With("content_type", headers.Get("Content-Type")),
			logger.With("trace", xRequests),
			logger.With("user_agent", c.Request.UserAgent()),
		)
		// Continue.
		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)
		latencys := fmt.Sprintf("%s", latency)
		uid := c.GetString("uid")
		//responseBody := blw.body.Bytes()
		// get code and message
		var response gin_helper.BaseResponse
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			logger.Info(
				"response body can not unmarshal to model.Response struct",
				logger.With("uid", uid),
				logger.With("trace", xRequests),
				logger.With("ip", ip),
				logger.With("method", method),
				logger.With("path", path),
				logger.With("requests_body", string(bodyBytes)),
				logger.With("content_type", headers.Get("Content-Type")),
				logger.With("user_agent", c.Request.UserAgent()),
			)

		} else {
			logger.Info(
				fmt.Sprintf("response success: %v", response.Success),
				logger.With("uid", uid),
				logger.With("trace", xRequests),
				logger.With("code", response.Code),
				logger.With("message", response.Message),
				logger.With("ip", ip),
				logger.With("latency", latencys),
				logger.With("method", method),
				logger.With("path", path),
				logger.With("requests_body", string(bodyBytes)),
				logger.With("content_type", headers.Get("Content-Type")),
				logger.With("user_agent", c.Request.UserAgent()),
			)
		}
	}
}

func TraceLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 每个请求生成的请求traceId具有全局唯一性
		logger.AddContext(ctx, logger.With("trace", ctx.MustGet("X-Request-Id")))
		// 为日志添加请求的地址以及请求参数等信息
		logger.AddContext(ctx, logger.With("method", ctx.Request.Method))
		logger.AddContext(ctx, logger.With("user_agent", ctx.Request.UserAgent()))
		logger.AddContext(ctx, logger.With("path", ctx.Request.URL.String()))

		// 将请求参数json序列化后添加进日志上下文
		if ctx.Request.Form == nil {
			_ = ctx.Request.ParseMultipartForm(32 << 20)
		}
		form, _ := json.Marshal(ctx.Request.Form)
		logger.AddContext(ctx, logger.With("params", string(form)))
		ctx.Next()
	}
}
