package middleware

import (
	"net/http"
	"strings"

	http2 "ilicense-lite/library/http"
	token2 "ilicense-lite/library/token"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		tokenString = strings.TrimSpace(tokenString)
		if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
			tokenString = strings.TrimSpace(tokenString[7:])
		}

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
			c.Set("userID", token.Claims.(jwt.MapClaims)["sub"])
		}

		c.Next()
	}
}
