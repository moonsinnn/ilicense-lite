package router

import (
	"ilicense-lite/controller"

	"github.com/gin-gonic/gin"
)

func ProductRouterInit(e *gin.Engine) {
	group := e.Group("/api/product")
	group.GET("/get", controller.ProductGet)
	group.POST("/add", controller.ProductAdd)
	group.POST("/query", controller.ProductQuery)
}
