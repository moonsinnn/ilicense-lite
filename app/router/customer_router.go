package router

import (
	"ilicense-lite/controller"

	"github.com/gin-gonic/gin"
)

func CustomerRouterInit(e *gin.Engine) {
	group := e.Group("/api/customer")
	group.GET("/get", controller.CustomerGet)
	group.POST("/add", controller.CustomerAdd)
	group.POST("/query", controller.CustomerQuery)
}
