package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"ilicense-lite/bootstrap/logger"
	"ilicense-lite/library/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CustomResponseWriter 自定义响应写入器
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 重写 Write 方法
func (crw *CustomResponseWriter) Write(b []byte) (int, error) {
	crw.body.Write(b)                  // 将响应体写入 buffer
	return crw.ResponseWriter.Write(b) // 调用原始的 Write 方法
}

// LoggerMiddleware 访问日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		start := time.Now()

		// 读取请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 创建自定义响应写入器
		body := bytes.NewBufferString("")
		crw := &CustomResponseWriter{
			ResponseWriter: c.Writer,
			body:           body,
		}
		c.Writer = crw

		// 处理请求
		c.Next()

		// 记录结束时间
		duration := time.Since(start)

		// 获取响应内容和 Content-Type
		responseBody := body.String()
		contentType := c.Writer.Header().Get("Content-Type")

		traceID, spanID := util.GetTraceInfo(c.Request.Context())
		// 记录请求信息
		logger.AppLogger.WithFields(logrus.Fields{
			"trace_id":      traceID,
			"span_id":       spanID,
			"method":        c.Request.Method,
			"path":          c.Request.URL.Path,
			"status":        c.Writer.Status(),
			"duration":      duration,
			"client_ip":     c.ClientIP(),
			"user_agent":    c.Request.UserAgent(),
			"request_body":  string(requestBody),
			"response_body": responseBody,
			"content_type":  contentType,
		}).Info("input processed")
	}
}
