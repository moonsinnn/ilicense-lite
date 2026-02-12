package router

import (
	"ilicense-lite/controller"

	"github.com/gin-gonic/gin"
)

func IssuerRouterInit(e *gin.Engine) {
	group := e.Group("/api/issuer")
	group.GET("/get", controller.IssuerGet)
	group.POST("/add", controller.IssuerAdd)
	group.POST("/query", controller.IssuerQuery)
	group.POST("/delete/:id", controller.IssuerDeleteOne)
	group.POST("/delete", controller.IssuerDelete)
}
