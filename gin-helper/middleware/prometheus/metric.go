package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

/**
* @Author: Jam Wong
* @Date: 2020/8/10
 */

var (
	ReqDuration *prometheus.HistogramVec
	ReqTotal    *prometheus.CounterVec
)

func init() {
	// HistogramVec
	ReqDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "The HTTP request latencies in seconds.",
		Buckets: nil,
	}, []string{"method", "path"})

	ReqTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests made.",
	}, []string{"method", "path", "status"})

	prometheus.MustRegister(
		ReqDuration,
		ReqTotal,
	)
}

// Metric metric middleware
func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		tBegin := time.Now()
		c.Next()

		duration := float64(time.Since(tBegin)) / float64(time.Second)

		ReqTotal.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"status": strconv.Itoa(c.Writer.Status()),
		}).Inc()

		ReqDuration.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}).Observe(duration)
	}
}
