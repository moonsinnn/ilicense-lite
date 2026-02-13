package router

import (
	"ilicense-lite/controller"
	"ilicense-lite/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(e *gin.Engine) {
	publicGroup := e.Group("/api/user")
	publicGroup.POST("/register", controller.UserRegister)
	publicGroup.POST("/login", controller.UserLogin)

	protectedGroup := e.Group("/api/user")
	protectedGroup.Use(middleware.RequireAuth())
	protectedGroup.GET("/profile", controller.UserProfileGet)
	protectedGroup.POST("/profile/update", controller.UserProfileUpdate)
	protectedGroup.POST("/password/update", controller.UserPasswordUpdate)
}
