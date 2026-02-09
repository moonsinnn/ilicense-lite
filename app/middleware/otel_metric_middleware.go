package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	meter              = otel.Meter("gin-app")
	httpRequestCount   metric.Int64Counter
	httpRequestLatency metric.Float64Histogram
)

func init() {
	httpRequestCount, _ = meter.Int64Counter(
		"otel_http_server_requests_total",
		metric.WithDescription("The total number of http requests from otel."),
	)
	httpRequestLatency, _ = meter.Float64Histogram(
		"otel_http_server_requests_latency",
		metric.WithDescription("The latency of http requests from otel."),
	)
}

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()

		labels := metric.WithAttributes(
			attribute.String("http.method", c.Request.Method),
			attribute.String("http.path", c.Request.URL.Path),
			attribute.String("http.host", c.Request.Host),
			attribute.String("http.scheme", c.Request.URL.Scheme),
			attribute.Int("http.status", c.Writer.Status()),
		)

		httpRequestCount.Add(c.Request.Context(), 1, labels)
		httpRequestLatency.Record(c.Request.Context(), duration, labels)
	}
}
