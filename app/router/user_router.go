package router

import (
	"ilicense-lite/controller"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(e *gin.Engine) {
	group := e.Group("/api/user")
	group.POST("/register", controller.UserRegister)
	group.POST("/login", controller.UserLogin)
	group.GET("/profile", controller.UserProfileGet)
	group.POST("/profile/update", controller.UserProfileUpdate)
	group.POST("/password/update", controller.UserPasswordUpdate)
}
