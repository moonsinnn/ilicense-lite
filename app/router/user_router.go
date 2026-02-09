package router

import (
	"ilicense-lite/controller"
	"ilicense-lite/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(e *gin.Engine) {
	group := e.Group("/api/user")
	group.GET("/get", middleware.AuthMiddleware(), controller.UserGet)
	group.POST("/add", controller.UserAdd)
	group.POST("/query", controller.UserQuery)
	group.POST("/login", controller.UserLogin)
	group.GET("/sign/in", controller.UserSignIn)
	group.GET("/sign/back", controller.UserSignBack)
}
