package middleware

import (
	"fmt"
	"net/http"

	http2 "ilicense-lite/library/http"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware 自定义异常处理中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, http2.BaseResponse[any]{
					Code:    http.StatusInternalServerError,
					Message: fmt.Sprintf("Internal Server Error: %v", r),
				})
				c.Abort() // 终止处理
			}
		}()
		c.Next() // 继续处理请求
	}
}
