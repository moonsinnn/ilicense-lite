package router

import (
	"ilicense-lite/controller"
	"ilicense-lite/middleware"

	"github.com/gin-gonic/gin"
)

func CustomerRouterInit(e *gin.Engine) {
	group := e.Group("/api/customer")
	group.Use(middleware.RequireAuth())
	group.GET("/get", controller.CustomerGet)
	group.POST("/add", controller.CustomerAdd)
	group.POST("/query", controller.CustomerQuery)
	group.POST("/delete", controller.CustomerDelete)
	group.POST("/delete/:id", controller.CustomerDeleteOne)
}
