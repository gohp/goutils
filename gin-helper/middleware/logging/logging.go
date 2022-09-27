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
		responseBody := blw.body.Bytes()
		// get code and message
		var response gin_helper.BaseResponse
		if err := json.Unmarshal(responseBody, &response); err != nil {
			logger.Info(
				"response body can not unmarshal to model.Response struct",
				logger.With("trace", xRequests),
				logger.With("latency", latencys),
				logger.With("response", responseBody),
			)

		} else {
			logger.Info(
				"response",
				logger.With("success", response.Success),
				logger.With("trace", xRequests),
				logger.With("latency", latencys),
			)
		}
	}
}

func TraceLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 每个请求生成的请求traceId具有全局唯一性
		logger.GAddContext(ctx, logger.With("trace", ctx.MustGet("X-Request-Id")))
		ctx.Next()
	}
}
