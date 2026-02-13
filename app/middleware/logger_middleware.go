package middleware

import (
	"bytes"
	"io"
	"strings"
	"time"

	"ilicense-lite/bootstrap/logger"
	"ilicense-lite/library/util"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

// CustomResponseWriter 自定义响应写入器
type CustomResponseWriter struct {
	gin.ResponseWriter
	body      *bytes.Buffer
	bodyLimit int
	isLimited bool
}

// Write 重写 Write 方法
func (crw *CustomResponseWriter) Write(b []byte) (int, error) {
	if !crw.isLimited {
		remaining := crw.bodyLimit - crw.body.Len()
		if remaining <= 0 {
			crw.isLimited = true
		} else if len(b) > remaining {
			crw.body.Write(b[:remaining])
			crw.isLimited = true
		} else {
			crw.body.Write(b)
		}
	}
	return crw.ResponseWriter.Write(b) // 调用原始的 Write 方法
}

// LoggerMiddleware 访问日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		start := time.Now()

		// 读取请求体
		var requestBody []byte
		requestContentType := c.Request.Header.Get("Content-Type")
		if c.Request.Body != nil && shouldLogBody(requestContentType) {
			requestBody, _ = io.ReadAll(io.LimitReader(c.Request.Body, maxBodyLength+1))
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 创建自定义响应写入器
		body := bytes.NewBufferString("")
		crw := &CustomResponseWriter{
			ResponseWriter: c.Writer,
			body:           body,
			bodyLimit:      maxBodyLength + 1,
		}
		c.Writer = crw

		// 处理请求
		c.Next()

		// 记录结束时间
		duration := time.Since(start)

		// 获取响应内容和 Content-Type
		responseBody := body.Bytes()
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
			"request_body":  sanitizeBody(requestBody, requestContentType),
			"response_body": sanitizeBody(responseBody, contentType),
			"content_type":  contentType,
		}).Info("input processed")
	}
}

var sensitiveFields = map[string]struct{}{
	"password":        {},
	"old_password":    {},
	"new_password":    {},
	"token":           {},
	"authorization":   {},
	"activation_code": {},
	"private_key":     {},
}

const (
	redactedValue = "***REDACTED***"
	maxBodyLength = 2048
)

func shouldLogBody(contentType string) bool {
	ct := strings.ToLower(contentType)
	return strings.Contains(ct, "application/json") ||
		strings.Contains(ct, "text/plain") ||
		strings.Contains(ct, "application/x-www-form-urlencoded")
}

func sanitizeBody(raw []byte, contentType string) string {
	if len(raw) == 0 {
		return ""
	}

	if !strings.Contains(strings.ToLower(contentType), "application/json") {
		return truncateString(string(raw), maxBodyLength)
	}

	var payload interface{}
	if err := json.Unmarshal(raw, &payload); err != nil {
		return truncateString(string(raw), maxBodyLength)
	}

	redact(payload)
	sanitized, err := json.Marshal(payload)
	if err != nil {
		return "[unmarshalable response]"
	}

	return truncateString(string(sanitized), maxBodyLength)
}

func redact(v interface{}) {
	switch data := v.(type) {
	case map[string]interface{}:
		for k, value := range data {
			if _, ok := sensitiveFields[strings.ToLower(k)]; ok {
				data[k] = redactedValue
				continue
			}
			redact(value)
		}
	case []interface{}:
		for _, item := range data {
			redact(item)
		}
	}
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "...(truncated)"
}
