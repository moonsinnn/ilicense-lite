package middleware

import (
	"net/http"

	http2 "ilicense-lite/library/http"
	token2 "ilicense-lite/library/token"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		//if tokenString == "" {
		//	c.JSON(http.StatusUnauthorized, http2.BaseResponse[any]{
		//		Code:    http.StatusUnauthorized,
		//		Message: "Authorization token is required",
		//	})
		//	c.Abort()
		//	return
		//}

		if tokenString != "" {
			token, err := token2.ValidateJWT(tokenString)
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, http2.BaseResponse[any]{
					Code:    http.StatusUnauthorized,
					Message: "Invalid or expired token",
				})
				c.Abort()
				return
			}
			c.Set("userID", token.Claims.(jwt.MapClaims)["sub"]) // 存储用户ID到上下文
		}

		c.Next()
	}
}
