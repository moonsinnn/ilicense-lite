package router

import (
	_ "ilicense-lite/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DocsRouterInit(e *gin.Engine) {
	group := e.Group("/docs")
	group.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
