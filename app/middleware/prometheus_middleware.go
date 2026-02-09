package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// 定义 Prometheus 指标
var (
	// 统计请求总数
	httpRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "total number of http requests",
		},
		[]string{"method", "path", "status"},
	)
	// 统计请求耗时
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "histogram of response time for handler in seconds",
		},
		[]string{"method", "path"},
	)
)

func init() {
	// 注册 Prometheus 指标
	prometheus.MustRegister(httpRequestTotal, httpRequestDuration)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 处理请求
		c.Next()
		// 计算请求耗时
		duration := time.Since(start).Seconds()
		// 获取请求信息
		method := c.Request.Method
		path := c.FullPath() // 直接获取路由 path
		status := c.Writer.Status()
		// 记录请求总数
		httpRequestTotal.WithLabelValues(method, path, strconv.Itoa(status)).Inc()
		// 记录请求耗时
		httpRequestDuration.WithLabelValues(method, path).Observe(duration)
	}
}
