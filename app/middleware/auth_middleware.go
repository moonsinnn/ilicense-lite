package middleware

import (
	"net/http"
	"strings"

	http2 "ilicense-lite/library/http"
	token2 "ilicense-lite/library/token"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const contextUserIDKey = "userID"

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
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				c.JSON(http.StatusUnauthorized, http2.BaseResponse[any]{
					Code:    http.StatusUnauthorized,
					Message: "Invalid token claims",
				})
				c.Abort()
				return
			}
			sub, ok := claims["sub"].(string)
			if !ok || strings.TrimSpace(sub) == "" {
				c.JSON(http.StatusUnauthorized, http2.BaseResponse[any]{
					Code:    http.StatusUnauthorized,
					Message: "Invalid token subject",
				})
				c.Abort()
				return
			}
			c.Set(contextUserIDKey, sub)
		}

		c.Next()
	}
}

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get(contextUserIDKey)
		userIDStr, typeOK := userID.(string)
		if !ok || !typeOK || strings.TrimSpace(userIDStr) == "" {
			c.JSON(http.StatusUnauthorized, http2.BaseResponse[any]{
				Code:    http.StatusUnauthorized,
				Message: "Authentication required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
