package middleware

import (
	"net/http"
	"runtime/debug"

	"ilicense-lite/bootstrap/logger"
	http2 "ilicense-lite/library/http"
	"ilicense-lite/library/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// RecoveryMiddleware 自定义异常处理中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				traceID, spanID := util.GetTraceInfo(c.Request.Context())
				logger.AppLogger.WithFields(logrus.Fields{
					"trace_id": traceID,
					"span_id":  spanID,
					"panic":    r,
					"stack":    string(debug.Stack()),
				}).Error("panic recovered")
				c.JSON(http.StatusInternalServerError, http2.BaseResponse[any]{
					Code:    http.StatusInternalServerError,
					Message: "internal server error",
				})
				c.Abort() // 终止处理
			}
		}()
		c.Next() // 继续处理请求
	}
}
