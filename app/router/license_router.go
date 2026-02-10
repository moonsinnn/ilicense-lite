package router

import (
	"ilicense-lite/controller"

	"github.com/gin-gonic/gin"
)

func LicenseRouterInit(e *gin.Engine) {
	group := e.Group("/api/license")
	group.GET("/get", controller.LicenseGet)
	group.POST("/add", controller.LicenseAdd)
	group.POST("/query", controller.LicenseQuery)
	group.POST("/activate", controller.LicenseActivate)
	group.POST("/renew", controller.LicenseRenew)
}
