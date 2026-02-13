package router

import (
	"ilicense-lite/controller"
	"ilicense-lite/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRouterInit(e *gin.Engine) {
	group := e.Group("/api/product")
	group.Use(middleware.RequireAuth())
	group.GET("/get", controller.ProductGet)
	group.POST("/add", controller.ProductAdd)
	group.POST("/query", controller.ProductQuery)
	group.POST("/delete", controller.ProductDelete)
	group.POST("/delete/:id", controller.ProductDeleteOne)
}
